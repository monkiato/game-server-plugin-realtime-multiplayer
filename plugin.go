package main

import (
	"context"
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/monkiato/game-server-core/pkg/framework"
	"github.com/monkiato/game-server-plugin-realtime-multiplayer/internal/game"
	"github.com/monkiato/game-server-plugin-realtime-multiplayer/internal/model"
	"github.com/monkiato/game-server-plugin-realtime-multiplayer/multiplayer"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type BaseServer struct {
	ServerManager *framework.ServerManager
}

type MultiplayerConnectionServer struct {
	multiplayer.UnimplementedMultiplayerConnectionServer
	BaseServer
	plugin *ClientAuthoritativeRealtimeMultiplayerPlugin
}

type ClientAuthoritativeRealtimeMultiplayerPlugin struct {
	matches map[uint]*game.MatchRegistry
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

	p.matches = map[uint]*game.MatchRegistry{}

	return nil
}

func (p *ClientAuthoritativeRealtimeMultiplayerPlugin) Initialize(serverManager *framework.ServerManager) error {
	log.Println("automigrate DB tables...")
	serverManager.DB.AutoMigrate(&model.Match{})
	serverManager.DB.AutoMigrate(&model.MatchPlayer{})
	serverManager.DB.AutoMigrate(&model.Player{})

	multiplayerConnectionServer := MultiplayerConnectionServer {
		BaseServer: BaseServer{
			ServerManager: serverManager,
		},
		plugin: p,
	}

	log.Println("register grpc server...")
	multiplayer.RegisterMultiplayerConnectionServer(serverManager.GrpcModule.Server, multiplayerConnectionServer)
	return nil
}

func (mp MultiplayerConnectionServer) CreateMatch(_ context.Context, request *multiplayer.MatchRequest) (*flatbuffers.Builder, error) {
	log.Println("CreateMatch called...")

	playerToken := string(request.RequestPlayer(nil).Token())

	newMatch, err := game.CreateMatch(mp.ServerManager.DB, playerToken, &game.MatchInfo{Capacity:2, Visibility:0})
	if err != nil {
		logrus.Errorf(err.Error())
		return nil, err
	}

	//load match registry in memory, required for broadcast messages
	mp.plugin.matches[newMatch.ID] = game.NewMatchRegistry(newMatch.ID)

	b := CreateMatchInfoBuilder(newMatch.ID)
	return b, nil
}

func (mp MultiplayerConnectionServer) JoinMatch(context.Context, *multiplayer.MatchRequest) (*flatbuffers.Builder, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinMatch not implemented")
}

func (mp MultiplayerConnectionServer) LeaveMatch(context.Context, *multiplayer.MatchRequest) (*flatbuffers.Builder, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveMatch not implemented")
}

func CreateMatchInfoBuilder(matchId uint) *flatbuffers.Builder {
	b := flatbuffers.NewBuilder(0)
	multiplayer.MatchInfoStart(b)
	multiplayer.MatchInfoAddMatchId(b, uint32(matchId))
	matchInfoOffset := multiplayer.MatchInfoEnd(b)
	multiplayer.MatchResponseStart(b)
	multiplayer.MatchResponseAddMatchInfo(b, matchInfoOffset)
	b.Finish(multiplayer.MatchResponseEnd(b))
	return b
}