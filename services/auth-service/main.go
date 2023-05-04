package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc/reflection"

	pb "github.com/prongbang/auth-service/proto/auth"
	"google.golang.org/grpc"
)

const port = ":50052"

// server is used to implement auth.AuthServer
type authServer struct {
	pb.UnimplementedAuthServer
}

func (a *authServer) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	log.Printf("Received: username=%v, password=%v", request.GetUsername(), request.GetPassword())

	return &pb.LoginResponse{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJlbSIsIm5hbWUiOiJkZXYgZGF5IiwiaWF0IjoxNTE2MjM5MDIyfQ.yNC-7RUVZCveMOANZcT7KWMczVkb_T7KnHv3fmMLiCI",
	}, nil
}

func NewAuthServer() pb.AuthServer {
	return &authServer{}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAuthServer(s, NewAuthServer())

	// Register reflection service on gRPC server.
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
