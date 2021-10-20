// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	proto "github.com/wosai/ultron/v2/pkg/statistics/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UltronServiceClient is the client API for UltronService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UltronServiceClient interface {
	Subscribe(ctx context.Context, in *Session, opts ...grpc.CallOption) (UltronService_SubscribeClient, error)
	SendUp(ctx context.Context, in *proto.AttackStatisticsDTO, opts ...grpc.CallOption) (*Feedback, error)
}

type ultronServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUltronServiceClient(cc grpc.ClientConnInterface) UltronServiceClient {
	return &ultronServiceClient{cc}
}

func (c *ultronServiceClient) Subscribe(ctx context.Context, in *Session, opts ...grpc.CallOption) (UltronService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &UltronService_ServiceDesc.Streams[0], "/wosai.ultron.UltronService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &ultronServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UltronService_SubscribeClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type ultronServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *ultronServiceSubscribeClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *ultronServiceClient) SendUp(ctx context.Context, in *proto.AttackStatisticsDTO, opts ...grpc.CallOption) (*Feedback, error) {
	out := new(Feedback)
	err := c.cc.Invoke(ctx, "/wosai.ultron.UltronService/SendUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UltronServiceServer is the server API for UltronService service.
// All implementations should embed UnimplementedUltronServiceServer
// for forward compatibility
type UltronServiceServer interface {
	Subscribe(*Session, UltronService_SubscribeServer) error
	SendUp(context.Context, *proto.AttackStatisticsDTO) (*Feedback, error)
}

// UnimplementedUltronServiceServer should be embedded to have forward compatible implementations.
type UnimplementedUltronServiceServer struct {
}

func (UnimplementedUltronServiceServer) Subscribe(*Session, UltronService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedUltronServiceServer) SendUp(context.Context, *proto.AttackStatisticsDTO) (*Feedback, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendUp not implemented")
}

// UnsafeUltronServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UltronServiceServer will
// result in compilation errors.
type UnsafeUltronServiceServer interface {
	mustEmbedUnimplementedUltronServiceServer()
}

func RegisterUltronServiceServer(s grpc.ServiceRegistrar, srv UltronServiceServer) {
	s.RegisterService(&UltronService_ServiceDesc, srv)
}

func _UltronService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Session)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UltronServiceServer).Subscribe(m, &ultronServiceSubscribeServer{stream})
}

type UltronService_SubscribeServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type ultronServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *ultronServiceSubscribeServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func _UltronService_SendUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto.AttackStatisticsDTO)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UltronServiceServer).SendUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wosai.ultron.UltronService/SendUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UltronServiceServer).SendUp(ctx, req.(*proto.AttackStatisticsDTO))
	}
	return interceptor(ctx, in, info, handler)
}

// UltronService_ServiceDesc is the grpc.ServiceDesc for UltronService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UltronService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wosai.ultron.UltronService",
	HandlerType: (*UltronServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendUp",
			Handler:    _UltronService_SendUp_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _UltronService_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ultron.proto",
}
