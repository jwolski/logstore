package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"

	"golang.org/x/net/context"

	"google.golang.org/grpc"

	"github.com/jwolski/logstore/api"
)

const (
	// Expected format of log file contents. See `file` flag.
	logPattern = `(?P<Owner>\S+) (?P<Bucket>\S+) (?P<RawTimestamp>\S+ \+\S+\]) (?P<ClientIp>\S+) (?P<Requester>\S+) (?P<RequestId>\S+) (?P<Operation>\S+) (?P<Key>\S+) "(?P<Verb>\S+) (?P<URI>\S+) (?P<Protocol>\S+)" (?P<StatusCode>\S+) (?P<ErrorCode>\S+) (?P<BytesSent>\S+) (?P<ObjectSize>\S+) (?P<TimeTotal>\S+) (?P<TimeTurnAround>\S+) "(?P<Referrer>.+)" "(?P<UserAgent>.+)" (?P<VersionId>\S+)`
)

var (
	filename = flag.String("file", "", "Log file")
	server   = flag.String("server", "", "Server address")
	verbose  = flag.Bool("verbose", false, "Debug output")
)

// Prepares requests by parsing log files and transforming log lines into
// `PutRequest` objects. Returns an empty slice of requests (and the error)
// if scanning the log file fails.
func parseLogFile(logFile *os.File) ([]*api.PutRequest, error) {
	putReqs := make([]*api.PutRequest, 0)
	// Scan the log file, parsing each log line (provided expected format), and
	// build a `PutRequest` out of each matching line.
	scanner := bufio.NewScanner(logFile)
	for scanner.Scan() {
		// Transform log line into request object. If parsing the log line
		// fails, continue parsing the rest of the log lines. If the verbose
		// flag has been specified, print an error as to why parsing the log
		// line had failed.
		req, err := line2Req(scanner.Text())
		if err != nil {
			if *verbose {
				// TODO: Print something.
			}
			continue
		}
		putReqs = append(putReqs, req)
	}

	if err := scanner.Err(); err != nil {
		return putReqs, err
	}

	return putReqs, nil
}

// Transforms a log line into a `PutRequest`.
func line2Req(line string) (*api.PutRequest, error) {
	// TODO: Implement this.
	regex, err := regexp.Compile(logPattern)
	if err != nil {
		return nil, err
	}

	// Build map of named groups to matched values. Construct a `PutRequest`
	// from matched values.
	names := regex.SubexpNames()
	matches := regex.FindStringSubmatch(line)

	name2Match := map[string]string{}
	for i, m := range matches {
		name2Match[names[i]] = m
	}

	log.Printf("%s", name2Match)
	return &api.PutRequest{
		Owner:          name2Match["Owner"],
		Bucket:         name2Match["Bucket"],
		Timestamp:      "",
		RawTimestamp:   name2Match["RawTimestamp"],
		ClientIp:       name2Match["ClientIp"],
		Requester:      name2Match["Requester"],
		Operation:      name2Match["Operation"],
		Key:            name2Match["Key"],
		Verb:           name2Match["Verb"],
		Uri:            name2Match["URI"],
		Protocol:       name2Match["Protocol"],
		StatusCode:     intOrDefault(name2Match, "StatusCode", -1),
		ErrorCode:      intOrDefault(name2Match, "ErrorCode", -1),
		BytesSent:      intOrDefault(name2Match, "BytesSent", -1),
		ObjectSize:     intOrDefault(name2Match, "ObjectSize", -1),
		TimeTotal:      intOrDefault(name2Match, "TimeTotal", -1),
		TimeTurnAround: intOrDefault(name2Match, "TimeTurnAround", -1),
		Referrer:       name2Match["Referrer"],
		UserAgent:      name2Match["UserAgent"],
		VersionId:      name2Match["VersionId"],
	}, nil
}

// Runs logstore client
func main() {
	flag.Parse()

	// file flag is required
	if *filename == "" {
		log.Fatalf("file is required")
	}

	if *server == "" {
		log.Fatalf("server is required")
	}

	file, err := os.Open(*filename)
	if err != nil {
		log.Fatalf("failed to open file")
	}
	defer file.Close()

	putReqs, err := parseLogFile(file)
	if err != nil {
		log.Fatalf("failed to parse log file")
	}
	os.Exit(1)

	// Connect to the gRPC server
	conn, err := grpc.Dial(*server, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to server: %s", err.Error())
	}
	defer conn.Close()

	// Create a new gRPC client and send requests to server. Catalog
	// all errors that have occurred. Report them after we're all done.
	client := api.NewLogClient(conn)
	putErrs := make([]error, 0)
	for _, r := range putReqs {
		resp, err := client.Put(context.Background(), r)
		// Treat this as a (gRPC) client error.
		if err != nil {
			putErrs = append(putErrs, err)
			continue
		}

		// Treat this as a server error.
		if resp.ErrCode != 0 {
			putErrs = append(putErrs,
				errors.New(fmt.Sprintf("server error: %d", resp.ErrCode)))
			continue
		}
	}

	if len(putErrs) > 0 {
		log.Printf("%d of %d requests failed.", len(putErrs), len(putReqs))
	}
	os.Exit(1)
}
