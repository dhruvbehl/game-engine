package grpc

import (
	"net"

	pbengine "github.com/dhruvbehl/game-apis/game-engine/v1"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	gamelogic "github.com/dhruvbehl/game-engine/internal/server/logic"
)

type Grpc struct {
	address string
	server *grpc.Server
}

func NewServer(address string) *Grpc {
	return &Grpc{address: address}
}

func (g *Grpc) GetSize(ctx context.Context, input *pbengine.GetSizeRequest) (*pbengine.GetSizeResponse, error) {
	log.Info().Msg("getSize in game-engine service")
	return &pbengine.GetSizeResponse{
		Size: gamelogic.GetSize(),
	}, nil
}

func (g *Grpc) SetScore(ctx context.Context, input *pbengine.SetScoreRequest) (*pbengine.SetScoreResponse, error) {
	log.Info().Msg("setScore in game-engine service")
	return &pbengine.SetScoreResponse{
		Set: gamelogic.SetScore(input.Score),
	}, nil
}

func (g *Grpc) ListenAndServe() error {
	listener, err := net.Listen("tcp", g.address)
	if err != nil {
		return errors.Wrap(err, "Failed to initialize tcp port")
	}

	serverOpt := []grpc.ServerOption{}
	g.server = grpc.NewServer(serverOpt...)

	pbengine.RegisterGameEngineServer(g.server, g)

	log.Info().Str("address", g.address).Msg("Initializing gRPC server for game-engine service")

	if err := g.server.Serve(listener); err != nil {
		return errors.Wrap(err, "Failed to initialize gRPC server for game-engine service")
	}
	return nil
}