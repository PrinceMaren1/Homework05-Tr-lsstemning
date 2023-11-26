// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: Proto/proto.proto

package Homework05

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

const (
	ServerConnection_UpdateAuctionState_FullMethodName = "/proto.ServerConnection/UpdateAuctionState"
	ServerConnection_GetAuctionState_FullMethodName    = "/proto.ServerConnection/GetAuctionState"
	ServerConnection_Bid_FullMethodName                = "/proto.ServerConnection/Bid"
	ServerConnection_ConnectToReplica_FullMethodName   = "/proto.ServerConnection/ConnectToReplica"
	ServerConnection_GetAlternateServer_FullMethodName = "/proto.ServerConnection/GetAlternateServer"
)

// ServerConnectionClient is the client API for ServerConnection service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerConnectionClient interface {
	UpdateAuctionState(ctx context.Context, in *AuctionState, opts ...grpc.CallOption) (*AuctionState, error)
	GetAuctionState(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AuctionState, error)
	Bid(ctx context.Context, in *ClientBid, opts ...grpc.CallOption) (*Acknowledgement, error)
	ConnectToReplica(ctx context.Context, in *Connection, opts ...grpc.CallOption) (*Empty, error)
	GetAlternateServer(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Connection, error)
}

type serverConnectionClient struct {
	cc grpc.ClientConnInterface
}

func NewServerConnectionClient(cc grpc.ClientConnInterface) ServerConnectionClient {
	return &serverConnectionClient{cc}
}

func (c *serverConnectionClient) UpdateAuctionState(ctx context.Context, in *AuctionState, opts ...grpc.CallOption) (*AuctionState, error) {
	out := new(AuctionState)
	err := c.cc.Invoke(ctx, ServerConnection_UpdateAuctionState_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverConnectionClient) GetAuctionState(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AuctionState, error) {
	out := new(AuctionState)
	err := c.cc.Invoke(ctx, ServerConnection_GetAuctionState_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverConnectionClient) Bid(ctx context.Context, in *ClientBid, opts ...grpc.CallOption) (*Acknowledgement, error) {
	out := new(Acknowledgement)
	err := c.cc.Invoke(ctx, ServerConnection_Bid_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverConnectionClient) ConnectToReplica(ctx context.Context, in *Connection, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, ServerConnection_ConnectToReplica_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serverConnectionClient) GetAlternateServer(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Connection, error) {
	out := new(Connection)
	err := c.cc.Invoke(ctx, ServerConnection_GetAlternateServer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServerConnectionServer is the server API for ServerConnection service.
// All implementations must embed UnimplementedServerConnectionServer
// for forward compatibility
type ServerConnectionServer interface {
	UpdateAuctionState(context.Context, *AuctionState) (*AuctionState, error)
	GetAuctionState(context.Context, *Empty) (*AuctionState, error)
	Bid(context.Context, *ClientBid) (*Acknowledgement, error)
	ConnectToReplica(context.Context, *Connection) (*Empty, error)
	GetAlternateServer(context.Context, *Empty) (*Connection, error)
	mustEmbedUnimplementedServerConnectionServer()
}

// UnimplementedServerConnectionServer must be embedded to have forward compatible implementations.
type UnimplementedServerConnectionServer struct {
}

func (UnimplementedServerConnectionServer) UpdateAuctionState(context.Context, *AuctionState) (*AuctionState, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAuctionState not implemented")
}
func (UnimplementedServerConnectionServer) GetAuctionState(context.Context, *Empty) (*AuctionState, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuctionState not implemented")
}
func (UnimplementedServerConnectionServer) Bid(context.Context, *ClientBid) (*Acknowledgement, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bid not implemented")
}
func (UnimplementedServerConnectionServer) ConnectToReplica(context.Context, *Connection) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConnectToReplica not implemented")
}
func (UnimplementedServerConnectionServer) GetAlternateServer(context.Context, *Empty) (*Connection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAlternateServer not implemented")
}
func (UnimplementedServerConnectionServer) mustEmbedUnimplementedServerConnectionServer() {}

// UnsafeServerConnectionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerConnectionServer will
// result in compilation errors.
type UnsafeServerConnectionServer interface {
	mustEmbedUnimplementedServerConnectionServer()
}

func RegisterServerConnectionServer(s grpc.ServiceRegistrar, srv ServerConnectionServer) {
	s.RegisterService(&ServerConnection_ServiceDesc, srv)
}

func _ServerConnection_UpdateAuctionState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuctionState)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerConnectionServer).UpdateAuctionState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServerConnection_UpdateAuctionState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerConnectionServer).UpdateAuctionState(ctx, req.(*AuctionState))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerConnection_GetAuctionState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerConnectionServer).GetAuctionState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServerConnection_GetAuctionState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerConnectionServer).GetAuctionState(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerConnection_Bid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientBid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerConnectionServer).Bid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServerConnection_Bid_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerConnectionServer).Bid(ctx, req.(*ClientBid))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerConnection_ConnectToReplica_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Connection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerConnectionServer).ConnectToReplica(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServerConnection_ConnectToReplica_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerConnectionServer).ConnectToReplica(ctx, req.(*Connection))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServerConnection_GetAlternateServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServerConnectionServer).GetAlternateServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServerConnection_GetAlternateServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServerConnectionServer).GetAlternateServer(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ServerConnection_ServiceDesc is the grpc.ServiceDesc for ServerConnection service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServerConnection_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ServerConnection",
	HandlerType: (*ServerConnectionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateAuctionState",
			Handler:    _ServerConnection_UpdateAuctionState_Handler,
		},
		{
			MethodName: "GetAuctionState",
			Handler:    _ServerConnection_GetAuctionState_Handler,
		},
		{
			MethodName: "Bid",
			Handler:    _ServerConnection_Bid_Handler,
		},
		{
			MethodName: "ConnectToReplica",
			Handler:    _ServerConnection_ConnectToReplica_Handler,
		},
		{
			MethodName: "GetAlternateServer",
			Handler:    _ServerConnection_GetAlternateServer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Proto/proto.proto",
}
