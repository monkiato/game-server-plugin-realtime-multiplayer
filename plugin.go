package main

import (
	"context"
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/monkiato/game-server-plugin-realtime-multiplayer/internal/game"
	"github.com/monkiato/game-server-plugin-realtime-multiplayer/multiplayer"
	"github.com/monkiato/game-server-core/pkg/framework"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type AbstractCoreServer struct {

}

type MultiplayerConnectionServer struct {
	multiplayer.UnimplementedMultiplayerConnectionServer
	AbstractCoreServer
}

type ClientAuthoritativeRealtimeMultiplayerPlugin struct {
	matches []*game.Match
}

func Create() framework.Plugin {
	return &ClientAuthoritativeRealtimeMultiplayerPlugin{}
}

func (p *ClientAuthoritativeRealtimeMultiplayerPlugin) GetApiName() string {
	return "multiplayer/realtime"
}

func (p *ClientAuthoritativeRealtimeMultiplayerPlugin) OnLoad() error {
	logrus.Debug("on load")
	//TODO: create a loop to control player sessions

	p.matches = []*game.Match{}

	return nil
}

func (p *ClientAuthoritativeRealtimeMultiplayerPlugin) Initialize(serverManager *framework.ServerManager) error {
	multiplayerConnectionServer := MultiplayerConnectionServer {
	}
	multiplayer.RegisterMultiplayerConnectionServer(serverManager.GrpcModule.Server, multiplayerConnectionServer)
	return nil
}

func (mp MultiplayerConnectionServer) CreateMatch(context.Context, *multiplayer.CreateMatchInfo) (*flatbuffers.Builder, error) {
	log.Println("CreateMatch called...")

	b := flatbuffers.NewBuilder(0)
	multiplayer.MatchInfoStart(b)
	multiplayer.MatchInfoAddMatchId(b, 23)
	matchInfoOffset := multiplayer.MatchInfoEnd(b)
	multiplayer.MatchResponseStart(b)
	multiplayer.MatchResponseAddMatchInfo(b, matchInfoOffset)
	b.Finish(multiplayer.MatchResponseEnd(b))
	return b, nil
}

func (mp MultiplayerConnectionServer) JoinMatch(context.Context, *multiplayer.MatchInfo) (*flatbuffers.Builder, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinMatch not implemented")
}

func (mp MultiplayerConnectionServer) LeaveMatch(context.Context, *multiplayer.MatchInfo) (*flatbuffers.Builder, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveMatch not implemented")
}