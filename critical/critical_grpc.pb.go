// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package criticalpackage

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

// CommunicationClient is the client API for Communication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommunicationClient interface {
	ServerReply(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
	SendRequest(ctx context.Context, opts ...grpc.CallOption) (Communication_SendRequestClient, error)
}

type communicationClient struct {
	cc grpc.ClientConnInterface
}

func NewCommunicationClient(cc grpc.ClientConnInterface) CommunicationClient {
	return &communicationClient{cc}
}

func (c *communicationClient) ServerReply(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/criticalpackage.Communication/serverReply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *communicationClient) SendRequest(ctx context.Context, opts ...grpc.CallOption) (Communication_SendRequestClient, error) {
	stream, err := c.cc.NewStream(ctx, &Communication_ServiceDesc.Streams[0], "/criticalpackage.Communication/sendRequest", opts...)
	if err != nil {
		return nil, err
	}
	x := &communicationSendRequestClient{stream}
	return x, nil
}

type Communication_SendRequestClient interface {
	Send(*Request) error
	CloseAndRecv() (*Ack, error)
	grpc.ClientStream
}

type communicationSendRequestClient struct {
	grpc.ClientStream
}

func (x *communicationSendRequestClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *communicationSendRequestClient) CloseAndRecv() (*Ack, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Ack)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CommunicationServer is the server API for Communication service.
// All implementations must embed UnimplementedCommunicationServer
// for forward compatibility
type CommunicationServer interface {
	ServerReply(context.Context, *Request) (*Reply, error)
	SendRequest(Communication_SendRequestServer) error
	mustEmbedUnimplementedCommunicationServer()
}

// UnimplementedCommunicationServer must be embedded to have forward compatible implementations.
type UnimplementedCommunicationServer struct {
}

func (UnimplementedCommunicationServer) ServerReply(context.Context, *Request) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServerReply not implemented")
}
func (UnimplementedCommunicationServer) SendRequest(Communication_SendRequestServer) error {
	return status.Errorf(codes.Unimplemented, "method SendRequest not implemented")
}
func (UnimplementedCommunicationServer) mustEmbedUnimplementedCommunicationServer() {}

// UnsafeCommunicationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommunicationServer will
// result in compilation errors.
type UnsafeCommunicationServer interface {
	mustEmbedUnimplementedCommunicationServer()
}

func RegisterCommunicationServer(s grpc.ServiceRegistrar, srv CommunicationServer) {
	s.RegisterService(&Communication_ServiceDesc, srv)
}

func _Communication_ServerReply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommunicationServer).ServerReply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/criticalpackage.Communication/serverReply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommunicationServer).ServerReply(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Communication_SendRequest_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CommunicationServer).SendRequest(&communicationSendRequestServer{stream})
}

type Communication_SendRequestServer interface {
	SendAndClose(*Ack) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type communicationSendRequestServer struct {
	grpc.ServerStream
}

func (x *communicationSendRequestServer) SendAndClose(m *Ack) error {
	return x.ServerStream.SendMsg(m)
}

func (x *communicationSendRequestServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Communication_ServiceDesc is the grpc.ServiceDesc for Communication service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Communication_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "criticalpackage.Communication",
	HandlerType: (*CommunicationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "serverReply",
			Handler:    _Communication_ServerReply_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "sendRequest",
			Handler:       _Communication_SendRequest_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "critical/critical.proto",
}
