package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"tinyUrl/internal/pkg/models"
	"tinyUrl/internal/pkg/url/delivery/server"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	authConn, err := grpc.Dial("localhost:5400", opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
		panic(err)
	}

	defer authConn.Close()

	authClient := server.NewDecreaseUrlClient(authConn)

	url := models.Url{Value: "http://vk.com"}

	tinyUrl, err := authClient.Create(context.Background(), &server.Url{Value: url.Value})
	if err != nil {
		return
	}

	fmt.Println(tinyUrl.GetValue())
}
