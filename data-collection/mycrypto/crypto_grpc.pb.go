// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.14.0
// source: protos/crypto.proto

package mycrypto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CryptoStreamClient is the client API for CryptoStream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CryptoStreamClient interface {
	StreamCrypto(ctx context.Context, opts ...grpc.CallOption) (CryptoStream_StreamCryptoClient, error)
}

type cryptoStreamClient struct {
	cc grpc.ClientConnInterface
}

func NewCryptoStreamClient(cc grpc.ClientConnInterface) CryptoStreamClient {
	return &cryptoStreamClient{cc}
}

func (c *cryptoStreamClient) StreamCrypto(ctx context.Context, opts ...grpc.CallOption) (CryptoStream_StreamCryptoClient, error) {
	stream, err := c.cc.NewStream(ctx, &CryptoStream_ServiceDesc.Streams[0], "/crypto.CryptoStream/StreamCrypto", opts...)
	if err != nil {
		return nil, err
	}
	x := &cryptoStreamStreamCryptoClient{stream}
	return x, nil
}

type CryptoStream_StreamCryptoClient interface {
	Send(*CryptoData) error
	CloseAndRecv() (*StreamResponse, error)
	grpc.ClientStream
}

type cryptoStreamStreamCryptoClient struct {
	grpc.ClientStream
}

func (x *cryptoStreamStreamCryptoClient) Send(m *CryptoData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *cryptoStreamStreamCryptoClient) CloseAndRecv() (*StreamResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CryptoStreamServer is the server API for CryptoStream service.
// All implementations must embed UnimplementedCryptoStreamServer
// for forward compatibility
type CryptoStreamServer interface {
	StreamCrypto(CryptoStream_StreamCryptoServer) error
	mustEmbedUnimplementedCryptoStreamServer()
}

// UnimplementedCryptoStreamServer must be embedded to have forward compatible implementations.
type UnimplementedCryptoStreamServer struct {
}

func (UnimplementedCryptoStreamServer) StreamCrypto(CryptoStream_StreamCryptoServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamCrypto not implemented")
}
func (UnimplementedCryptoStreamServer) mustEmbedUnimplementedCryptoStreamServer() {}

// UnsafeCryptoStreamServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CryptoStreamServer will
// result in compilation errors.
type UnsafeCryptoStreamServer interface {
	mustEmbedUnimplementedCryptoStreamServer()
}

func RegisterCryptoStreamServer(s grpc.ServiceRegistrar, srv CryptoStreamServer) {
	s.RegisterService(&CryptoStream_ServiceDesc, srv)
}

func _CryptoStream_StreamCrypto_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CryptoStreamServer).StreamCrypto(&cryptoStreamStreamCryptoServer{stream})
}

type CryptoStream_StreamCryptoServer interface {
	SendAndClose(*StreamResponse) error
	Recv() (*CryptoData, error)
	grpc.ServerStream
}

type cryptoStreamStreamCryptoServer struct {
	grpc.ServerStream
}

func (x *cryptoStreamStreamCryptoServer) SendAndClose(m *StreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *cryptoStreamStreamCryptoServer) Recv() (*CryptoData, error) {
	m := new(CryptoData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CryptoStream_ServiceDesc is the grpc.ServiceDesc for CryptoStream service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CryptoStream_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "crypto.CryptoStream",
	HandlerType: (*CryptoStreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamCrypto",
			Handler:       _CryptoStream_StreamCrypto_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "protos/crypto.proto",
}