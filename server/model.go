package main

// Acts as a shim between the GRPC server contract and that which is stored in
// the database.
type logRecord struct {
	Owner          string
	Bucket         string
	Timestamp      string
	RawTimestamp   string
	ClientIp       string
	Requester      string
	RequestId      string
	Operation      string
	Key            string
	Verb           string
	Uri            string
	Protocol       string
	StatusCode     int32
	ErrorCode      int32
	BytesSent      int32
	ObjectSize     int32
	TimeTotal      int32
	TimeTurnAround int32
	Referrer       string
	UserAgent      string
	VersionId      string
}
