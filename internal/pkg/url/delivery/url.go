package delivery

import (
	"context"
	"github.com/asaskevich/govalidator"
	_ "github.com/jackc/pgx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
	"tinyUrl/internal/pkg/models"
	"tinyUrl/internal/pkg/url/delivery/server"
	"tinyUrl/internal/pkg/url/usecase"
	"tinyUrl/internal/tinyUrl/utils"
)

type DecreaseUrlServer struct {
	server.UnimplementedDecreaseUrlServer
	Usecase usecase.UrlUsecaseInterface
}

func (d DecreaseUrlServer) Create(ctx context.Context, url *server.Url) (*server.TinyUrl, error) {
	ok := govalidator.IsURL(url.GetValue())
	if !ok {
		return &server.TinyUrl{}, status.Error(codes.InvalidArgument, "invalid URL")
	}

	tinyURL, err := d.Usecase.CreateURL(models.Url{Value: url.GetValue()})
	if err != nil {
		return &server.TinyUrl{}, status.Error(codes.Internal, "Server Error")
	}

	return &server.TinyUrl{Value: tinyURL.Value}, nil
}

func (d DecreaseUrlServer) Get(ctx context.Context, url *server.TinyUrl) (*server.Url, error) {
	url.Value = strings.TrimLeft(url.Value, "http://")

	getUrl, err := d.Usecase.GetUrl(models.Url{Value: url.GetValue()})

	if err != nil {
		utils.MainLogger.LogInfo(err)
		return nil, status.Error(codes.NotFound, "URL is missing")
	}

	utils.MainLogger.LogInfo("Success Get URL")
	return &server.Url{Value: getUrl.Value}, nil
}
