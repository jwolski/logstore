package main

import (
	"bufio"
	"flag"
	"log"
	"os"

	"github.com/jwolski/logstore/api"
)

const (
	// Expected format of log file contents. See `file` flag.
	logPattern = `(\\S+) (\\S+) (\\S+ \\+\\S+\\]) (\\S+) (\\S+) (\\S+) (\\S+)
	(\\S+) \"(\\S+) (\\S+) (\\S+)\" (\\S+) (\\S+) (\\S+) (\\S+) (\\S+) (\\S+)
	\"(.+)\" \"(.+)\" (\\S+)`
)

var (
	filename = flag.String("file", "", "Log file")
	verbose  = flag.Bool("verbose", false, "Debug output")
)

// Prepares requests by parsing log files and transforming log lines into
// `PutRequest` objects. Returns an empty slice of requests (and the error)
// if scanning the log file fails.
func parseLogFile(logFile *os.File) ([]api.PutRequest, error) {
	putReqs := make([]api.PutRequest, 0)
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
func line2Req(line string) (api.PutRequest, error) {
	// TODO: Implement this.
	return api.PutRequest{Owner: line}, nil
}

// Runs logstore client
func main() {
	flag.Parse()

	// file flag is required
	if *filename == "" {
		log.Fatalf("file is required")
	}

	file, err := os.Open(*filename)
	if err != nil {
		log.Fatalf("failed to open file")
	}

	putReqs, err := parseLogFile(file)
	if err != nil {
		log.Fatalf("failed to parse log file")
	}

	// Send all requests individually to logstore server.
	for _, r := range putReqs {
		log.Println(r.Owner)
	}

	// Cleanup
	file.Close()
}
