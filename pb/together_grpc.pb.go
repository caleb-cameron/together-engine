// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// GameServiceClient is the client API for GameService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GameServiceClient interface {
	Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (GameService_ConnectClient, error)
	SendPlayerUpdates(ctx context.Context, opts ...grpc.CallOption) (GameService_SendPlayerUpdatesClient, error)
	LoadChunk(ctx context.Context, in *Vector, opts ...grpc.CallOption) (*Chunk, error)
	UpdateTile(ctx context.Context, in *TileUpdate, opts ...grpc.CallOption) (*Ack, error)
	Register(ctx context.Context, in *UserRegistration, opts ...grpc.CallOption) (*LoginResponse, error)
	Login(ctx context.Context, in *UserLogin, opts ...grpc.CallOption) (*LoginResponse, error)
}

type gameServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGameServiceClient(cc grpc.ClientConnInterface) GameServiceClient {
	return &gameServiceClient{cc}
}

func (c *gameServiceClient) Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (GameService_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &GameService_ServiceDesc.Streams[0], "/pb.GameService/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &gameServiceConnectClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GameService_ConnectClient interface {
	Recv() (*GameState, error)
	grpc.ClientStream
}

type gameServiceConnectClient struct {
	grpc.ClientStream
}

func (x *gameServiceConnectClient) Recv() (*GameState, error) {
	m := new(GameState)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gameServiceClient) SendPlayerUpdates(ctx context.Context, opts ...grpc.CallOption) (GameService_SendPlayerUpdatesClient, error) {
	stream, err := c.cc.NewStream(ctx, &GameService_ServiceDesc.Streams[1], "/pb.GameService/SendPlayerUpdates", opts...)
	if err != nil {
		return nil, err
	}
	x := &gameServiceSendPlayerUpdatesClient{stream}
	return x, nil
}

type GameService_SendPlayerUpdatesClient interface {
	Send(*PlayerEvent) error
	CloseAndRecv() (*Ack, error)
	grpc.ClientStream
}

type gameServiceSendPlayerUpdatesClient struct {
	grpc.ClientStream
}

func (x *gameServiceSendPlayerUpdatesClient) Send(m *PlayerEvent) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gameServiceSendPlayerUpdatesClient) CloseAndRecv() (*Ack, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Ack)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gameServiceClient) LoadChunk(ctx context.Context, in *Vector, opts ...grpc.CallOption) (*Chunk, error) {
	out := new(Chunk)
	err := c.cc.Invoke(ctx, "/pb.GameService/LoadChunk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) UpdateTile(ctx context.Context, in *TileUpdate, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, "/pb.GameService/UpdateTile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) Register(ctx context.Context, in *UserRegistration, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/pb.GameService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gameServiceClient) Login(ctx context.Context, in *UserLogin, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/pb.GameService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameServiceServer is the server API for GameService service.
// All implementations must embed UnimplementedGameServiceServer
// for forward compatibility
type GameServiceServer interface {
	Connect(*ConnectRequest, GameService_ConnectServer) error
	SendPlayerUpdates(GameService_SendPlayerUpdatesServer) error
	LoadChunk(context.Context, *Vector) (*Chunk, error)
	UpdateTile(context.Context, *TileUpdate) (*Ack, error)
	Register(context.Context, *UserRegistration) (*LoginResponse, error)
	Login(context.Context, *UserLogin) (*LoginResponse, error)
	mustEmbedUnimplementedGameServiceServer()
}

// UnimplementedGameServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGameServiceServer struct {
}

func (UnimplementedGameServiceServer) Connect(*ConnectRequest, GameService_ConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedGameServiceServer) SendPlayerUpdates(GameService_SendPlayerUpdatesServer) error {
	return status.Errorf(codes.Unimplemented, "method SendPlayerUpdates not implemented")
}
func (UnimplementedGameServiceServer) LoadChunk(context.Context, *Vector) (*Chunk, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadChunk not implemented")
}
func (UnimplementedGameServiceServer) UpdateTile(context.Context, *TileUpdate) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTile not implemented")
}
func (UnimplementedGameServiceServer) Register(context.Context, *UserRegistration) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedGameServiceServer) Login(context.Context, *UserLogin) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedGameServiceServer) mustEmbedUnimplementedGameServiceServer() {}

// UnsafeGameServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GameServiceServer will
// result in compilation errors.
type UnsafeGameServiceServer interface {
	mustEmbedUnimplementedGameServiceServer()
}

func RegisterGameServiceServer(s grpc.ServiceRegistrar, srv GameServiceServer) {
	s.RegisterService(&GameService_ServiceDesc, srv)
}

func _GameService_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConnectRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GameServiceServer).Connect(m, &gameServiceConnectServer{stream})
}

type GameService_ConnectServer interface {
	Send(*GameState) error
	grpc.ServerStream
}

type gameServiceConnectServer struct {
	grpc.ServerStream
}

func (x *gameServiceConnectServer) Send(m *GameState) error {
	return x.ServerStream.SendMsg(m)
}

func _GameService_SendPlayerUpdates_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GameServiceServer).SendPlayerUpdates(&gameServiceSendPlayerUpdatesServer{stream})
}

type GameService_SendPlayerUpdatesServer interface {
	SendAndClose(*Ack) error
	Recv() (*PlayerEvent, error)
	grpc.ServerStream
}

type gameServiceSendPlayerUpdatesServer struct {
	grpc.ServerStream
}

func (x *gameServiceSendPlayerUpdatesServer) SendAndClose(m *Ack) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gameServiceSendPlayerUpdatesServer) Recv() (*PlayerEvent, error) {
	m := new(PlayerEvent)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GameService_LoadChunk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Vector)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).LoadChunk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GameService/LoadChunk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).LoadChunk(ctx, req.(*Vector))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_UpdateTile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TileUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).UpdateTile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GameService/UpdateTile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).UpdateTile(ctx, req.(*TileUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRegistration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GameService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).Register(ctx, req.(*UserRegistration))
	}
	return interceptor(ctx, in, info, handler)
}

func _GameService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLogin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.GameService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameServiceServer).Login(ctx, req.(*UserLogin))
	}
	return interceptor(ctx, in, info, handler)
}

// GameService_ServiceDesc is the grpc.ServiceDesc for GameService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GameService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.GameService",
	HandlerType: (*GameServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoadChunk",
			Handler:    _GameService_LoadChunk_Handler,
		},
		{
			MethodName: "UpdateTile",
			Handler:    _GameService_UpdateTile_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _GameService_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _GameService_Login_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _GameService_Connect_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SendPlayerUpdates",
			Handler:       _GameService_SendPlayerUpdates_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "together.proto",
}
