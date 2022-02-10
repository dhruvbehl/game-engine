package grpc

import (
	"math"

	pbengine "github.com/dhruvbehl/game-apis/game-engine/v1"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/rs/zerolog/log"
)

type Grpc struct {
	address string
	server *grpc.Server
}

var size = math.MaxFloat64
var score = 0

func NewServer(address string) *Grpc {
	return &Grpc{address: address}
}

func (g *Grpc) GetSize(ctx context.Context, input *pbengine.GetSizeRequest) (*pbengine.GetSizeResponse, error) {
	log.Info().Msg("getSize in game-engine service")
	return &pbengine.GetSizeResponse{
		Size: size,
	}, nil
}

func (g *Grpc) SetScore(ctx context.Context, input *pbengine.SetScoreRequest) (*pbengine.SetScoreResponse, error) {
	return nil, nil
}