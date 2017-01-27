package main

import (
	"errors"
	"fmt"

	"github.com/jwolski/logstore/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Sender struct {
	conn *grpc.ClientConn
}

func NewSender(server string) (*Sender, error) {
	sender := &Sender{}
	// Connect to the gRPC server
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		return sender, err
	}
	sender.conn = conn
	return sender, nil
}

func (s *Sender) Close() {
	s.conn.Close()
}

func (s *Sender) Send(putReqs []*api.PutRequest) []error {
	// Create a new gRPC client and send requests to server. Catalog
	// all errors that have occurred. Report them after we're all done.
	putErrs := make([]error, 0)

	client := api.NewLogClient(s.conn)
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

	return putErrs
}
