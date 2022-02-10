package main

import (
	"flag"
	"time"

	pbengine "github.com/dhruvbehl/game-apis/game-engine/v1"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	address := flag.String("address", "localhost:9003", "address to connect to game-engine service")
	connection, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatal().Err(err).Msg("failed to dial game-engine gRPC service")
	}

	defer func(){
		if err:= connection.Close(); err != nil {
			log.Fatal().Err(err).Str("address", *address).Msg("failed to close the connection")
		}
	}()

	timeOutContext, cancel := context.WithTimeout(context.Background(), time.Second * 10)
	defer cancel()

	client := pbengine.NewGameEngineClient(connection)
	if client == nil {
		log.Info().Msg("Client nil")
	}

	request, err := client.GetSize(timeOutContext, &pbengine.GetSizeRequest{})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get a response")
	}
	if request != nil {
		log.Info().Interface("size", request.GetSize()).Msg("Size received from game-engine service")
	} else {
		log.Error().Msg("Couldn't get highscore")
	}
	


}