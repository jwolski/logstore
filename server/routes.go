package main

import (
	"log"

	"github.com/jwolski/logstore/api"

	"golang.org/x/net/context"
)

// Route handler for the logstore server
type routes struct {
	store store
}

type errorCode uint

// Possible errors from route handlers
const (
	errNone = 0
	errSave = 1
)

// Creates a route handler for the logstore server
func newRoutes(store store) *routes {
	return &routes{store}
}

// Controls the route that puts log records into the storage backend
func (r *routes) Put(ctx context.Context, req *api.PutRequest) (*api.PutResponse, error) {
	resp := &api.PutResponse{errNone}

	// Save log record to the storage backend. Respond with a `errSave` if
	// anything goes wrong.
	rec := req2LogRecord(req)
	if err := r.store.save(rec); err != nil {
		resp.ErrCode = errSave
		log.Printf("Failed to store log record with request ID: %s", rec.RequestId)
		return resp, nil
	}
	// TODO: We probably don't need to log requests to this degree. But it is
	// nice debug information...
	log.Printf("Stored log record with request ID: %s", rec.RequestId)
	return resp, nil
}
