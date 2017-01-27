package main

import (
	"bufio"
	"os"
	"regexp"

	"github.com/jwolski/logstore/api"
)

const (
	// Expected format of log file contents. See `file` flag.
	logPattern = `(?P<Owner>\S+) (?P<Bucket>\S+) (?P<RawTimestamp>\S+ \+\S+\]) (?P<ClientIp>\S+) (?P<Requester>\S+) (?P<RequestId>\S+) (?P<Operation>\S+) (?P<Key>\S+) "(?P<Verb>\S+) (?P<URI>\S+) (?P<Protocol>\S+)" (?P<StatusCode>\S+) (?P<ErrorCode>\S+) (?P<BytesSent>\S+) (?P<ObjectSize>\S+) (?P<TimeTotal>\S+) (?P<TimeTurnAround>\S+) "(?P<Referrer>.+)" "(?P<UserAgent>.+)" (?P<VersionId>\S+)`
)

type Parser struct{}

// Prepares requests by parsing log files and transforming log lines into
// `PutRequest` objects. Returns an empty slice of requests (and the error)
// if scanning the log file fails.
func (p *Parser) Parse(logFile *os.File) ([]*api.PutRequest, error) {
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
			// TODO: Don't ignore this error outright. Do something better.
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
	// Match the log line against the expected regex and convert matched groups
	// to a prepared `PutRequest`.
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

	// Translate matched groups into a `PutRequest`.
	return &api.PutRequest{
		Owner:          name2Match["Owner"],
		Bucket:         name2Match["Bucket"],
		Timestamp:      "", // TODO: Fill this field in.
		RawTimestamp:   name2Match["RawTimestamp"],
		ClientIp:       name2Match["ClientIp"],
		Requester:      name2Match["Requester"],
		RequestId:      name2Match["RequestId"],
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
