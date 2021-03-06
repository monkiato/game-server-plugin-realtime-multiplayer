//Generated by gRPC Go plugin
//If you make any local changes, they will be lost
//source: multiplayer

package multiplayer

import (
	context "context"
	flatbuffers "github.com/google/flatbuffers/go"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Client API for MultiplayerConnection service
type MultiplayerConnectionClient interface {
	CreateMatch(ctx context.Context, in *flatbuffers.Builder,
		opts ...grpc.CallOption) (*MatchResponse, error)
	JoinMatch(ctx context.Context, in *flatbuffers.Builder,
		opts ...grpc.CallOption) (*MatchResponse, error)
	LeaveMatch(ctx context.Context, in *flatbuffers.Builder,
		opts ...grpc.CallOption) (*LeaveMatchResponse, error)
}

type multiplayerConnectionClient struct {
	cc grpc.ClientConnInterface
}

func NewMultiplayerConnectionClient(cc grpc.ClientConnInterface) MultiplayerConnectionClient {
	return &multiplayerConnectionClient{cc}
}

func (c *multiplayerConnectionClient) CreateMatch(ctx context.Context, in *flatbuffers.Builder,
	opts ...grpc.CallOption) (*MatchResponse, error) {
	out := new(MatchResponse)
	err := c.cc.Invoke(ctx, "/multiplayer.MultiplayerConnection/CreateMatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *multiplayerConnectionClient) JoinMatch(ctx context.Context, in *flatbuffers.Builder,
	opts ...grpc.CallOption) (*MatchResponse, error) {
	out := new(MatchResponse)
	err := c.cc.Invoke(ctx, "/multiplayer.MultiplayerConnection/JoinMatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *multiplayerConnectionClient) LeaveMatch(ctx context.Context, in *flatbuffers.Builder,
	opts ...grpc.CallOption) (*LeaveMatchResponse, error) {
	out := new(LeaveMatchResponse)
	err := c.cc.Invoke(ctx, "/multiplayer.MultiplayerConnection/LeaveMatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MultiplayerConnection service
type MultiplayerConnectionServer interface {
	CreateMatch(context.Context, *MatchRequest) (*flatbuffers.Builder, error)
	JoinMatch(context.Context, *MatchRequest) (*flatbuffers.Builder, error)
	LeaveMatch(context.Context, *MatchRequest) (*flatbuffers.Builder, error)
	mustEmbedUnimplementedMultiplayerConnectionServer()
}

type UnimplementedMultiplayerConnectionServer struct {
}

func (UnimplementedMultiplayerConnectionServer) CreateMatch(context.Context, *MatchRequest) (*flatbuffers.Builder, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMatch not implemented")
}

func (UnimplementedMultiplayerConnectionServer) JoinMatch(context.Context, *MatchRequest) (*flatbuffers.Builder, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinMatch not implemented")
}

func (UnimplementedMultiplayerConnectionServer) LeaveMatch(context.Context, *MatchRequest) (*flatbuffers.Builder, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveMatch not implemented")
}

func (UnimplementedMultiplayerConnectionServer) mustEmbedUnimplementedMultiplayerConnectionServer() {}

type UnsafeMultiplayerConnectionServer interface {
	mustEmbedUnimplementedMultiplayerConnectionServer()
}

func RegisterMultiplayerConnectionServer(s grpc.ServiceRegistrar, srv MultiplayerConnectionServer) {
	s.RegisterService(&_MultiplayerConnection_serviceDesc, srv)
}

func _MultiplayerConnection_CreateMatch_Handler(srv interface{}, ctx context.Context,
	dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MultiplayerConnectionServer).CreateMatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/multiplayer.MultiplayerConnection/CreateMatch",
	}

	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MultiplayerConnectionServer).CreateMatch(ctx, req.(*MatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _MultiplayerConnection_JoinMatch_Handler(srv interface{}, ctx context.Context,
	dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MultiplayerConnectionServer).JoinMatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/multiplayer.MultiplayerConnection/JoinMatch",
	}

	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MultiplayerConnectionServer).JoinMatch(ctx, req.(*MatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _MultiplayerConnection_LeaveMatch_Handler(srv interface{}, ctx context.Context,
	dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MultiplayerConnectionServer).LeaveMatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/multiplayer.MultiplayerConnection/LeaveMatch",
	}

	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MultiplayerConnectionServer).LeaveMatch(ctx, req.(*MatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}
var _MultiplayerConnection_serviceDesc = grpc.ServiceDesc{
	ServiceName: "multiplayer.MultiplayerConnection",
	HandlerType: (*MultiplayerConnectionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMatch",
			Handler:    _MultiplayerConnection_CreateMatch_Handler,
		},
		{
			MethodName: "JoinMatch",
			Handler:    _MultiplayerConnection_JoinMatch_Handler,
		},
		{
			MethodName: "LeaveMatch",
			Handler:    _MultiplayerConnection_LeaveMatch_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
	},
}
