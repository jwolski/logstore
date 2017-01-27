package main

import "github.com/jwolski/logstore/api"

func req2LogRecord(req *api.PutRequest) *logRecord {
	return &logRecord{
		Owner:          req.Owner,
		Bucket:         req.Bucket,
		Timestamp:      req.Timestamp,
		RawTimestamp:   req.RawTimestamp,
		ClientIp:       req.ClientIp,
		Requester:      req.Requester,
		RequestId:      req.RequestId,
		Operation:      req.Operation,
		Key:            req.Key,
		Verb:           req.Verb,
		Uri:            req.Uri,
		Protocol:       req.Protocol,
		StatusCode:     req.StatusCode,
		ErrorCode:      req.ErrorCode,
		BytesSent:      req.BytesSent,
		ObjectSize:     req.ObjectSize,
		TimeTotal:      req.TimeTotal,
		TimeTurnAround: req.TimeTurnAround,
		Referrer:       req.Referrer,
		UserAgent:      req.UserAgent,
		VersionId:      req.VersionId,
	}
}
