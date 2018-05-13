package main

import (
	"log"
	"net"

	"github.com/UmaruCMS/auth-system/config"
	pb "github.com/UmaruCMS/auth-system/rpc/base/auth"
	rpc "github.com/UmaruCMS/auth-system/rpc/service"
	"google.golang.org/grpc"
)

func release() {
	config.Release()
}

func main() {
	defer release()
	// user.RegisterUser("Lawrence", "lawrence.lee@foxmail.com", "123456")
	// _, tokenString := user.Login("lawrence.lee@foxmail.com", "123456")
	// fmt.Println(tokenString)
	// tokenInfo := &model.TokenInfo{}
	// tokenInfo.GetFromRedis(tokenString)
	// tokenInfo.Delete()
	// fmt.Println(user.VerifyTokenString(token))

	// r := router.DefaultRouter()
	// router.RegisterHandlers(r)
	// r.Run(":2333")

	lis, err := net.Listen("tcp", ":2334")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterAuthServer(grpcServer, &rpc.AuthServer{})
	grpcServer.Serve(lis)
}
