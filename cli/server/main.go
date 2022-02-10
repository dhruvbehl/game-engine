package main

import (
	"flag"

	grpcbackend "github.com/dhruvbehl/game-engine/internal/server/grpc"
	"github.com/rs/zerolog/log"
)

func main() {
	address := flag.String("address", ":9003", "address to connect to game-engine service")
	flag.Parse()

	server := grpcbackend.NewServer(*address)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize gRPC server for service game-engine ")
	}
}