package main

import (
	"context"
	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/monkiato/game-server-plugin-realtime-multiplayer/multiplayer"
	"google.golang.org/grpc"
	"testing"
)

func TestCreateMatch(t *testing.T) {
	testCases := []struct {
		name string
		message string
		expectedError bool
	} {
		{
			name: "req ok",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			conn, err := grpc.Dial(":5000", grpc.WithInsecure(), grpc.WithDefaultCallOptions(grpc.ForceCodec(flatbuffers.FlatbuffersCodec{})))
			if err != nil {
				t.Fatalf("failed to dial: %v", err)
			}
			defer conn.Close()
			client := multiplayer.NewMultiplayerConnectionClient(conn)

			b := flatbuffers.NewBuilder(0)
			tokenOffset := b.CreateString("player123")
			multiplayer.RequestPlayerInfoStart(b)
			multiplayer.RequestPlayerInfoAddToken(b, tokenOffset)
			playerInfoOffset := multiplayer.RequestPlayerInfoEnd(b)
			multiplayer.MatchRequestStart(b)
			multiplayer.MatchRequestAddRequestPlayer(b, playerInfoOffset)
			b.Finish(multiplayer.MatchRequestEnd(b))

			resp, err := client.CreateMatch(context.Background(), b)

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if resp.PlayersLength() != 0 {
				t.Fatalf("expected count=0, got %d", resp.PlayersLength())
			}

			matchInfo := resp.MatchInfo(nil)
			t.Logf("match id: %d", matchInfo.MatchId())
		})
	}
}