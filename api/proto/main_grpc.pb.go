// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// SendMoveRequestClient is the client API for SendMoveRequest service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SendMoveRequestClient interface {
	Move(ctx context.Context, opts ...grpc.CallOption) (SendMoveRequest_MoveClient, error)
	Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error)
}

type sendMoveRequestClient struct {
	cc grpc.ClientConnInterface
}

func NewSendMoveRequestClient(cc grpc.ClientConnInterface) SendMoveRequestClient {
	return &sendMoveRequestClient{cc}
}

func (c *sendMoveRequestClient) Move(ctx context.Context, opts ...grpc.CallOption) (SendMoveRequest_MoveClient, error) {
	stream, err := c.cc.NewStream(ctx, &SendMoveRequest_ServiceDesc.Streams[0], "/proto.sendMoveRequest/Move", opts...)
	if err != nil {
		return nil, err
	}
	x := &sendMoveRequestMoveClient{stream}
	return x, nil
}

type SendMoveRequest_MoveClient interface {
	Send(*MoveRequest) error
	Recv() (*MoveAnswer, error)
	grpc.ClientStream
}

type sendMoveRequestMoveClient struct {
	grpc.ClientStream
}

func (x *sendMoveRequestMoveClient) Send(m *MoveRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sendMoveRequestMoveClient) Recv() (*MoveAnswer, error) {
	m := new(MoveAnswer)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sendMoveRequestClient) Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (*ConnectResponse, error) {
	out := new(ConnectResponse)
	err := c.cc.Invoke(ctx, "/proto.sendMoveRequest/Connect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SendMoveRequestServer is the server API for SendMoveRequest service.
// All implementations must embed UnimplementedSendMoveRequestServer
// for forward compatibility
type SendMoveRequestServer interface {
	Move(SendMoveRequest_MoveServer) error
	Connect(context.Context, *ConnectRequest) (*ConnectResponse, error)
	mustEmbedUnimplementedSendMoveRequestServer()
}

// UnimplementedSendMoveRequestServer must be embedded to have forward compatible implementations.
type UnimplementedSendMoveRequestServer struct {
}

func (UnimplementedSendMoveRequestServer) Move(SendMoveRequest_MoveServer) error {
	return status.Errorf(codes.Unimplemented, "method Move not implemented")
}
func (UnimplementedSendMoveRequestServer) Connect(context.Context, *ConnectRequest) (*ConnectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedSendMoveRequestServer) mustEmbedUnimplementedSendMoveRequestServer() {}

// UnsafeSendMoveRequestServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SendMoveRequestServer will
// result in compilation errors.
type UnsafeSendMoveRequestServer interface {
	mustEmbedUnimplementedSendMoveRequestServer()
}

func RegisterSendMoveRequestServer(s grpc.ServiceRegistrar, srv SendMoveRequestServer) {
	s.RegisterService(&SendMoveRequest_ServiceDesc, srv)
}

func _SendMoveRequest_Move_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SendMoveRequestServer).Move(&sendMoveRequestMoveServer{stream})
}

type SendMoveRequest_MoveServer interface {
	Send(*MoveAnswer) error
	Recv() (*MoveRequest, error)
	grpc.ServerStream
}

type sendMoveRequestMoveServer struct {
	grpc.ServerStream
}

func (x *sendMoveRequestMoveServer) Send(m *MoveAnswer) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sendMoveRequestMoveServer) Recv() (*MoveRequest, error) {
	m := new(MoveRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SendMoveRequest_Connect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SendMoveRequestServer).Connect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.sendMoveRequest/Connect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SendMoveRequestServer).Connect(ctx, req.(*ConnectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SendMoveRequest_ServiceDesc is the grpc.ServiceDesc for SendMoveRequest service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SendMoveRequest_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.sendMoveRequest",
	HandlerType: (*SendMoveRequestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Connect",
			Handler:    _SendMoveRequest_Connect_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Move",
			Handler:       _SendMoveRequest_Move_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/main.proto",
}
