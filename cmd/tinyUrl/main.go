package main

import (
	"context"
	_ "github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"log"
	"math/rand"
	"net"
	"time"
	"tinyUrl/internal/pkg/url/delivery"
	"tinyUrl/internal/pkg/url/delivery/server"
	"tinyUrl/internal/pkg/url/repository"
	"tinyUrl/internal/pkg/url/usecase"
	"tinyUrl/internal/tinyUrl"
	"tinyUrl/internal/tinyUrl/utils"
)

type UserServerInterceptor struct {
	Logger *utils.Logger
}

func (s *UserServerInterceptor) logger(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	start := time.Now()
	md, _ := metadata.FromIncomingContext(ctx)

	reqId := rand.Uint64()

	s.Logger.Logger = s.Logger.Logger.WithFields(logrus.Fields{
		"requestId": reqId,
		"method":    info.FullMethod,
		"context":   md,
		"request":   req,
		"response":  resp,
		"error":     err,
		"work_time": time.Since(start),
	})

	s.Logger.LogInfo("Entry Point")

	reply, err := handler(ctx, req)

	s.Logger.LogInfo("USER Interceptor")
	return reply, err
}

func main() {
	rand.NewSource(time.Now().UnixNano())
	utils.MainLogger = &utils.Logger{Logger: logrus.NewEntry(logrus.StandardLogger())}
	listener, err := net.Listen("tcp", ":5400")

	ServerInterceptor := UserServerInterceptor{utils.MainLogger}
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(ServerInterceptor.logger))

	urlServer := delivery.DecreaseUrlServer{
		Usecase: usecase.UrlUsecase{DB: repository.UrlRepository{DB: tinyUrl.Init()}},
	}

	server.RegisterDecreaseUrlServer(grpcServer, &urlServer)
	log.Print("URL Server START at 5400")
	err = grpcServer.Serve(listener)

	if err != nil {
		grpclog.Fatalf("failed to serve: %v", err)
		return
	}
}
