package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/jwolski/logstore/api"

	"google.golang.org/grpc"
)

type DbBackend string

// List of supported storage backends for persistence of log records
const (
	Mongo DbBackend = "mongo"
)

var (
	port = flag.Uint("port", 5000, "Listen port")
	db   = flag.String("db", string(Mongo), "Database backend")
)

func setupStore() (store, error) {
	// Only MongoDB is supported thus far.
	switch *db {
	default:
		mongoConf := defaultMongoConf
		// Use the env variable (provided by Wercker) to set the Mongo URL. This
		// is not great as it assumes operating within a Wercker environmemnt.
		if addr := os.Getenv("MONGO_PORT_27017_TCP_ADDR"); addr != "" {
			mongoConf.url = addr
		}
		return dialMongo(mongoConf)
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
		log.Fatalf("setup store failed: %s", err.Error())
	}
	log.Printf("connected to db @ %s", store.addr())

	// Setup TCP server
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("server listen failed")
	}
	log.Printf("server listening on %d", *port)

	// Setup GRPC server and start serving traffic
	grpcServer := grpc.NewServer()
	api.RegisterLogServer(grpcServer, newRoutes(store))
	grpcServer.Serve(l)
}
