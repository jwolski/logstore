package main

import (
	"flag"
	"log"
	"os"
)

var (
	filename = flag.String("file", "", "Log file")
	server   = flag.String("server", "", "Server address")
)

// Checks flags for validity. If there is a missing field, the program will be
// aborted.
func checkFlags() {
	// file flag is required
	if *filename == "" {
		log.Fatalf("file is required")
	}

	if *server == "" {
		log.Fatalf("server is required")
	}
}

// Runs logstore client
func main() {
	// Parse and check flags
	flag.Parse()
	checkFlags()

	file, err := os.Open(*filename)
	if err != nil {
		log.Fatalf("failed to open file")
	}
	defer file.Close()

	// Parse log file and convert into `PutRequests` to be sent to server.
	parser := &Parser{}
	putReqs, err := parser.Parse(file)
	if err != nil {
		log.Fatalf("failed to parse log file: %s", err.Error())
	}

	// Create a request sender and send requests.
	sender, err := NewSender(*server)
	if err != nil {
		log.Fatalf("failed to create sender: %s", err.Error())
	}
	defer sender.Close()

	// Send requests. If any result in errors, report them.
	errs := sender.Send(putReqs)
	if len(errs) > 0 {
		log.Fatalf("%d of %d requests failed.", len(errs), len(putReqs))
	}
	log.Printf("%d requests were sent successfully.", len(putReqs))
}
