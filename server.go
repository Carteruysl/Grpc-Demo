package main

import (
	"context"
	"flag"
	"fmt"
	poker "grpc-demo/poker"
	pb "grpc-demo/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedPokerServer
}

func (s *server) GetNuts(ctx context.Context, req *pb.GetNutsRequest) (*pb.GetNutsResponse, error) {
	res, err := poker.PokerEvaluator(req.Hand, req.River)
	return &pb.GetNutsResponse{
		Card: res,
	}, err
}

func main() {
	flag.Parse()
	// 設定要監聽的 port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// 使用 gRPC 的 NewServer meethod 來建立 gRPC Server
	s := grpc.NewServer()
	pb.RegisterPokerServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	// 在 gRPC 伺服器上註冊反射服務。
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
