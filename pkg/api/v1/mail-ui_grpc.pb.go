// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// MailUIClient is the client API for MailUI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MailUIClient interface {
	// ListEngineSpecs returns a list of Email Engine(s) that can be started through the UI.
	ListEngineSpecs(ctx context.Context, in *ListEngineSpecsRequest, opts ...grpc.CallOption) (MailUI_ListEngineSpecsClient, error)
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error)
}

type mailUIClient struct {
	cc grpc.ClientConnInterface
}

func NewMailUIClient(cc grpc.ClientConnInterface) MailUIClient {
	return &mailUIClient{cc}
}

func (c *mailUIClient) ListEngineSpecs(ctx context.Context, in *ListEngineSpecsRequest, opts ...grpc.CallOption) (MailUI_ListEngineSpecsClient, error) {
	stream, err := c.cc.NewStream(ctx, &MailUI_ServiceDesc.Streams[0], "/v1.MailUI/ListEngineSpecs", opts...)
	if err != nil {
		return nil, err
	}
	x := &mailUIListEngineSpecsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MailUI_ListEngineSpecsClient interface {
	Recv() (*ListEngineSpecsResponse, error)
	grpc.ClientStream
}

type mailUIListEngineSpecsClient struct {
	grpc.ClientStream
}

func (x *mailUIListEngineSpecsClient) Recv() (*ListEngineSpecsResponse, error) {
	m := new(ListEngineSpecsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *mailUIClient) IsReadOnly(ctx context.Context, in *IsReadOnlyRequest, opts ...grpc.CallOption) (*IsReadOnlyResponse, error) {
	out := new(IsReadOnlyResponse)
	err := c.cc.Invoke(ctx, "/v1.MailUI/IsReadOnly", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MailUIServer is the server API for MailUI service.
// All implementations must embed UnimplementedMailUIServer
// for forward compatibility
type MailUIServer interface {
	// ListEngineSpecs returns a list of Email Engine(s) that can be started through the UI.
	ListEngineSpecs(*ListEngineSpecsRequest, MailUI_ListEngineSpecsServer) error
	// IsReadOnly returns true if the UI is readonly.
	IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error)
	mustEmbedUnimplementedMailUIServer()
}

// UnimplementedMailUIServer must be embedded to have forward compatible implementations.
type UnimplementedMailUIServer struct {
}

func (UnimplementedMailUIServer) ListEngineSpecs(*ListEngineSpecsRequest, MailUI_ListEngineSpecsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListEngineSpecs not implemented")
}
func (UnimplementedMailUIServer) IsReadOnly(context.Context, *IsReadOnlyRequest) (*IsReadOnlyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsReadOnly not implemented")
}
func (UnimplementedMailUIServer) mustEmbedUnimplementedMailUIServer() {}

// UnsafeMailUIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MailUIServer will
// result in compilation errors.
type UnsafeMailUIServer interface {
	mustEmbedUnimplementedMailUIServer()
}

func RegisterMailUIServer(s grpc.ServiceRegistrar, srv MailUIServer) {
	s.RegisterService(&MailUI_ServiceDesc, srv)
}

func _MailUI_ListEngineSpecs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListEngineSpecsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MailUIServer).ListEngineSpecs(m, &mailUIListEngineSpecsServer{stream})
}

type MailUI_ListEngineSpecsServer interface {
	Send(*ListEngineSpecsResponse) error
	grpc.ServerStream
}

type mailUIListEngineSpecsServer struct {
	grpc.ServerStream
}

func (x *mailUIListEngineSpecsServer) Send(m *ListEngineSpecsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _MailUI_IsReadOnly_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsReadOnlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailUIServer).IsReadOnly(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.MailUI/IsReadOnly",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailUIServer).IsReadOnly(ctx, req.(*IsReadOnlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MailUI_ServiceDesc is the grpc.ServiceDesc for MailUI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MailUI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.MailUI",
	HandlerType: (*MailUIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsReadOnly",
			Handler:    _MailUI_IsReadOnly_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListEngineSpecs",
			Handler:       _MailUI_ListEngineSpecs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "mail-ui.proto",
}
