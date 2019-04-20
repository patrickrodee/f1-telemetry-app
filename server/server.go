package server

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/patrickrodee/f1-telemetry-app/game"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 10000, "The server port")
)

type F1 struct {
	store *game.Store
}

func NewF1() *F1 {
	return &F1{game.NewStore()}
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	grpcServer.Serve(lis)
}
