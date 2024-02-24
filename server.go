package user

import (
	"context"
	"google.golang.org/grpc"
	"net"
)

type UserRpc struct {
	UnimplementedStreamGreeterServer
}

func (r *UserRpc) Greet(context.Context, *StreamReq) (*StreamResp, error) {
	return nil, nil
}

func NewRpcServer() {
	grpcServer := grpc.NewServer()
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	RegisterStreamGreeterServer(grpcServer, new(UserRpc))
	err = grpcServer.Serve(listen)
	if err != nil {
		panic(err)
	}
}
