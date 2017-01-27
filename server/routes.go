package main

import (
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
	if err := r.store.save(req2LogRecord(req)); err != nil {
		resp.ErrCode = errSave
	}

	return resp, nil
}
