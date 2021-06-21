package delivery

import (
	"context"
	"tinyUrl/internal/pkg/url/delivery/server"
)

type DecreaseUrl struct {
	server.UnimplementedDecreaseUrlServer
}

func (DecreaseUrl) Create(ctx context.Context, url *server.Url) (*server.TinyUrl, error) {
	panic("implement me")
}

func (DecreaseUrl) Get(ctx context.Context, url *server.TinyUrl) (*server.Url, error) {
	panic("implement me")
}
