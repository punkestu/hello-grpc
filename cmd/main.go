package main

import (
	userService "github.com/punkestu/hello-grpc/internal/user/service"
	userProto "github.com/punkestu/hello-grpc/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":8080") // create listener network
	if err != nil {
		log.Fatalln(err.Error())
	}
	grpcServer := grpc.NewServer()
	// register the user service to the grpc server
	userProto.RegisterUserServiceServer(grpcServer, userService.NewUserService())

	if err := grpcServer.Serve(listen); err != nil { // run grpc server
		log.Fatalln(err.Error())
	}
}
