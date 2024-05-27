// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.0--rc3
// source: phase.proto

package phase

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	PhaseEqualibrium_Vle_FullMethodName  = "/phase_proto.PhaseEqualibrium/vle"
	PhaseEqualibrium_Init_FullMethodName = "/phase_proto.PhaseEqualibrium/Init"
)

// PhaseEqualibriumClient is the client API for PhaseEqualibrium service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PhaseEqualibriumClient interface {
	// rpc Init(InitMessageRequest) returns (initMessageResponse);
	Vle(ctx context.Context, in *VleMessageRequest, opts ...grpc.CallOption) (*VleMessageResponse, error)
	Init(ctx context.Context, in *InitMessageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type phaseEqualibriumClient struct {
	cc grpc.ClientConnInterface
}

func NewPhaseEqualibriumClient(cc grpc.ClientConnInterface) PhaseEqualibriumClient {
	return &phaseEqualibriumClient{cc}
}

func (c *phaseEqualibriumClient) Vle(ctx context.Context, in *VleMessageRequest, opts ...grpc.CallOption) (*VleMessageResponse, error) {
	out := new(VleMessageResponse)
	err := c.cc.Invoke(ctx, PhaseEqualibrium_Vle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *phaseEqualibriumClient) Init(ctx context.Context, in *InitMessageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, PhaseEqualibrium_Init_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PhaseEqualibriumServer is the server API for PhaseEqualibrium service.
// All implementations must embed UnimplementedPhaseEqualibriumServer
// for forward compatibility
type PhaseEqualibriumServer interface {
	// rpc Init(InitMessageRequest) returns (initMessageResponse);
	Vle(context.Context, *VleMessageRequest) (*VleMessageResponse, error)
	Init(context.Context, *InitMessageRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedPhaseEqualibriumServer()
}

// UnimplementedPhaseEqualibriumServer must be embedded to have forward compatible implementations.
type UnimplementedPhaseEqualibriumServer struct {
}

func (UnimplementedPhaseEqualibriumServer) Vle(context.Context, *VleMessageRequest) (*VleMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Vle not implemented")
}
func (UnimplementedPhaseEqualibriumServer) Init(context.Context, *InitMessageRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (UnimplementedPhaseEqualibriumServer) mustEmbedUnimplementedPhaseEqualibriumServer() {}

// UnsafePhaseEqualibriumServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PhaseEqualibriumServer will
// result in compilation errors.
type UnsafePhaseEqualibriumServer interface {
	mustEmbedUnimplementedPhaseEqualibriumServer()
}

func RegisterPhaseEqualibriumServer(s grpc.ServiceRegistrar, srv PhaseEqualibriumServer) {
	s.RegisterService(&PhaseEqualibrium_ServiceDesc, srv)
}

func _PhaseEqualibrium_Vle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VleMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhaseEqualibriumServer).Vle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PhaseEqualibrium_Vle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhaseEqualibriumServer).Vle(ctx, req.(*VleMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhaseEqualibrium_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhaseEqualibriumServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PhaseEqualibrium_Init_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhaseEqualibriumServer).Init(ctx, req.(*InitMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PhaseEqualibrium_ServiceDesc is the grpc.ServiceDesc for PhaseEqualibrium service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PhaseEqualibrium_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "phase_proto.PhaseEqualibrium",
	HandlerType: (*PhaseEqualibriumServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "vle",
			Handler:    _PhaseEqualibrium_Vle_Handler,
		},
		{
			MethodName: "Init",
			Handler:    _PhaseEqualibrium_Init_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "phase.proto",
}
