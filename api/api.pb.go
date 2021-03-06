// Code generated by protoc-gen-go.
// source: api.proto
// DO NOT EDIT!

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	PutRequest
	PutResponse
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PutRequest struct {
	Owner          string `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	Bucket         string `protobuf:"bytes,2,opt,name=bucket" json:"bucket,omitempty"`
	Timestamp      string `protobuf:"bytes,3,opt,name=timestamp" json:"timestamp,omitempty"`
	RawTimestamp   string `protobuf:"bytes,4,opt,name=rawTimestamp" json:"rawTimestamp,omitempty"`
	ClientIp       string `protobuf:"bytes,5,opt,name=clientIp" json:"clientIp,omitempty"`
	Requester      string `protobuf:"bytes,6,opt,name=requester" json:"requester,omitempty"`
	RequestId      string `protobuf:"bytes,7,opt,name=requestId" json:"requestId,omitempty"`
	Operation      string `protobuf:"bytes,8,opt,name=operation" json:"operation,omitempty"`
	Key            string `protobuf:"bytes,9,opt,name=key" json:"key,omitempty"`
	Verb           string `protobuf:"bytes,10,opt,name=verb" json:"verb,omitempty"`
	Uri            string `protobuf:"bytes,11,opt,name=uri" json:"uri,omitempty"`
	Protocol       string `protobuf:"bytes,12,opt,name=protocol" json:"protocol,omitempty"`
	StatusCode     int32  `protobuf:"varint,13,opt,name=statusCode" json:"statusCode,omitempty"`
	ErrorCode      int32  `protobuf:"varint,14,opt,name=errorCode" json:"errorCode,omitempty"`
	BytesSent      int32  `protobuf:"varint,15,opt,name=bytesSent" json:"bytesSent,omitempty"`
	ObjectSize     int32  `protobuf:"varint,16,opt,name=objectSize" json:"objectSize,omitempty"`
	TimeTotal      int32  `protobuf:"varint,17,opt,name=timeTotal" json:"timeTotal,omitempty"`
	TimeTurnAround int32  `protobuf:"varint,18,opt,name=timeTurnAround" json:"timeTurnAround,omitempty"`
	Referrer       string `protobuf:"bytes,19,opt,name=referrer" json:"referrer,omitempty"`
	UserAgent      string `protobuf:"bytes,20,opt,name=userAgent" json:"userAgent,omitempty"`
	VersionId      string `protobuf:"bytes,21,opt,name=versionId" json:"versionId,omitempty"`
}

func (m *PutRequest) Reset()                    { *m = PutRequest{} }
func (m *PutRequest) String() string            { return proto.CompactTextString(m) }
func (*PutRequest) ProtoMessage()               {}
func (*PutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PutRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *PutRequest) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *PutRequest) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *PutRequest) GetRawTimestamp() string {
	if m != nil {
		return m.RawTimestamp
	}
	return ""
}

func (m *PutRequest) GetClientIp() string {
	if m != nil {
		return m.ClientIp
	}
	return ""
}

func (m *PutRequest) GetRequester() string {
	if m != nil {
		return m.Requester
	}
	return ""
}

func (m *PutRequest) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *PutRequest) GetOperation() string {
	if m != nil {
		return m.Operation
	}
	return ""
}

func (m *PutRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *PutRequest) GetVerb() string {
	if m != nil {
		return m.Verb
	}
	return ""
}

func (m *PutRequest) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *PutRequest) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *PutRequest) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *PutRequest) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *PutRequest) GetBytesSent() int32 {
	if m != nil {
		return m.BytesSent
	}
	return 0
}

func (m *PutRequest) GetObjectSize() int32 {
	if m != nil {
		return m.ObjectSize
	}
	return 0
}

func (m *PutRequest) GetTimeTotal() int32 {
	if m != nil {
		return m.TimeTotal
	}
	return 0
}

func (m *PutRequest) GetTimeTurnAround() int32 {
	if m != nil {
		return m.TimeTurnAround
	}
	return 0
}

func (m *PutRequest) GetReferrer() string {
	if m != nil {
		return m.Referrer
	}
	return ""
}

func (m *PutRequest) GetUserAgent() string {
	if m != nil {
		return m.UserAgent
	}
	return ""
}

func (m *PutRequest) GetVersionId() string {
	if m != nil {
		return m.VersionId
	}
	return ""
}

type PutResponse struct {
	ErrCode int32 `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
}

func (m *PutResponse) Reset()                    { *m = PutResponse{} }
func (m *PutResponse) String() string            { return proto.CompactTextString(m) }
func (*PutResponse) ProtoMessage()               {}
func (*PutResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PutResponse) GetErrCode() int32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func init() {
	proto.RegisterType((*PutRequest)(nil), "api.PutRequest")
	proto.RegisterType((*PutResponse)(nil), "api.PutResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Log service

type LogClient interface {
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error)
}

type logClient struct {
	cc *grpc.ClientConn
}

func NewLogClient(cc *grpc.ClientConn) LogClient {
	return &logClient{cc}
}

func (c *logClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := grpc.Invoke(ctx, "/api.Log/Put", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Log service

type LogServer interface {
	Put(context.Context, *PutRequest) (*PutResponse, error)
}

func RegisterLogServer(s *grpc.Server, srv LogServer) {
	s.RegisterService(&_Log_serviceDesc, srv)
}

func _Log_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Log/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Log_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Log",
	HandlerType: (*LogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Put",
			Handler:    _Log_Put_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x92, 0xcf, 0x6e, 0x13, 0x31,
	0x10, 0x87, 0x59, 0x36, 0x49, 0xbb, 0xd3, 0xd2, 0x06, 0x53, 0xd0, 0xa8, 0x42, 0xa8, 0xda, 0x03,
	0x54, 0x1c, 0x2a, 0x01, 0x4f, 0x50, 0x71, 0x5a, 0x89, 0x43, 0x95, 0xf6, 0x05, 0xf6, 0xcf, 0x10,
	0x99, 0x24, 0xf6, 0x32, 0xb6, 0x13, 0x85, 0xa7, 0xe5, 0x51, 0x90, 0x6d, 0xd6, 0x9b, 0xf4, 0xe6,
	0xdf, 0xf7, 0x4d, 0x26, 0xe3, 0x59, 0x43, 0x51, 0xf7, 0xf2, 0xae, 0x67, 0x6d, 0xb5, 0xc8, 0xeb,
	0x5e, 0x96, 0x7f, 0x27, 0x00, 0x0f, 0xce, 0x2e, 0xe8, 0xb7, 0x23, 0x63, 0xc5, 0x15, 0x4c, 0xf5,
	0x4e, 0x11, 0x63, 0x76, 0x93, 0xdd, 0x16, 0x8b, 0x18, 0xc4, 0x3b, 0x98, 0x35, 0xae, 0x5d, 0x91,
	0xc5, 0x97, 0x01, 0xff, 0x4f, 0xe2, 0x3d, 0x14, 0x56, 0x6e, 0xc8, 0xd8, 0x7a, 0xd3, 0x63, 0x1e,
	0xd4, 0x08, 0x44, 0x09, 0xe7, 0x5c, 0xef, 0x9e, 0x52, 0xc1, 0x24, 0x14, 0x1c, 0x31, 0x71, 0x0d,
	0xa7, 0xed, 0x5a, 0x92, 0xb2, 0x55, 0x8f, 0xd3, 0xe0, 0x53, 0xf6, 0xdd, 0x39, 0x8e, 0x45, 0x8c,
	0xb3, 0xd8, 0x3d, 0x81, 0x03, 0x5b, 0x75, 0x78, 0x72, 0x64, 0xab, 0xce, 0x5b, 0xdd, 0x13, 0xd7,
	0x56, 0x6a, 0x85, 0xa7, 0xd1, 0x26, 0x20, 0xe6, 0x90, 0xaf, 0x68, 0x8f, 0x45, 0xe0, 0xfe, 0x28,
	0x04, 0x4c, 0xb6, 0xc4, 0x0d, 0x42, 0x40, 0xe1, 0xec, 0xab, 0x1c, 0x4b, 0x3c, 0x8b, 0x55, 0x8e,
	0xa5, 0x9f, 0x36, 0xac, 0xae, 0xd5, 0x6b, 0x3c, 0x8f, 0xd3, 0x0e, 0x59, 0x7c, 0x00, 0x30, 0xb6,
	0xb6, 0xce, 0x7c, 0xd7, 0x1d, 0xe1, 0xab, 0x9b, 0xec, 0x76, 0xba, 0x38, 0x20, 0x7e, 0x22, 0x62,
	0xd6, 0x1c, 0xf4, 0x45, 0xd0, 0x23, 0xf0, 0xb6, 0xd9, 0x5b, 0x32, 0x8f, 0xa4, 0x2c, 0x5e, 0x46,
	0x9b, 0x80, 0xef, 0xad, 0x9b, 0x5f, 0xd4, 0xda, 0x47, 0xf9, 0x87, 0x70, 0x1e, 0x7b, 0x8f, 0x64,
	0xf8, 0x0e, 0x4f, 0xda, 0xd6, 0x6b, 0x7c, 0x1d, 0x7f, 0x9d, 0x80, 0xf8, 0x08, 0x17, 0x21, 0x38,
	0x56, 0xf7, 0xac, 0x9d, 0xea, 0x50, 0x84, 0x92, 0x67, 0xd4, 0xdf, 0x8e, 0xe9, 0x27, 0x31, 0x13,
	0xe3, 0x9b, 0x78, 0xbb, 0x21, 0xfb, 0x7f, 0x70, 0x86, 0xf8, 0x7e, 0xe9, 0xe7, 0xbb, 0x8a, 0xfb,
	0x4c, 0xc0, 0xdb, 0x2d, 0xb1, 0x91, 0x5a, 0x55, 0x1d, 0xbe, 0x8d, 0x36, 0x81, 0xf2, 0x13, 0x9c,
	0x85, 0x17, 0x66, 0x7a, 0xad, 0x0c, 0x09, 0x84, 0x13, 0xe2, 0xb8, 0x86, 0x2c, 0xcc, 0x31, 0xc4,
	0xaf, 0x5f, 0x20, 0xff, 0xa1, 0x97, 0xe2, 0x33, 0xe4, 0x0f, 0xce, 0x8a, 0xcb, 0x3b, 0xff, 0x54,
	0xc7, 0xb7, 0x79, 0x3d, 0x1f, 0x41, 0x6c, 0x55, 0xbe, 0x68, 0x66, 0x61, 0xff, 0xdf, 0xfe, 0x05,
	0x00, 0x00, 0xff, 0xff, 0x94, 0xfa, 0x74, 0x9e, 0xd7, 0x02, 0x00, 0x00,
}
