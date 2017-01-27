package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/jwolski/logstore/api"

	"google.golang.org/grpc"
)

type DbBackend string

// List of supported storage backends for persistence of log records
const (
	Mongo DbBackend = "mongo"
)

var (
	port = flag.Uint("port", 3333, "Listen port")
	db   = flag.String("db", string(Mongo), "Database backend")
)

func setupStore() (store, error) {
	// Only MongoDB is supported thus far.
	switch *db {
	default:
		// TODO: Expose Mongo config as optional command-line flag
		return dialMongo(defaultMongoConf)
	}
}

// Runs logstore server
// TODO: Implement handler for clean shutdown
func main() {
	flag.Parse()

	// Setup storage backend for log record persistence.
	// TODO: Close() the `store` on shutdown
	store, err := setupStore()
	if err != nil {
		log.Fatalf("setup store failed")
	}

	// Setup TCP server
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("server listen failed")
	}
	log.Printf("tcp listening on %d", *port)

	// Setup GRPC server and start serving traffic
	grpcServer := grpc.NewServer()
	api.RegisterLogServer(grpcServer, newRoutes(store))
	grpcServer.Serve(l)
}
