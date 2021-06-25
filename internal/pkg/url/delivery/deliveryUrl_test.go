package delivery

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/go-playground/assert.v1"
	"strings"
	"testing"
	"tinyUrl/internal/pkg/models"
	"tinyUrl/internal/pkg/url/delivery/server"
	mock_repository "tinyUrl/internal/pkg/url/repository/mocks"
	"tinyUrl/internal/pkg/url/usecase"
	"tinyUrl/internal/tinyUrl/utils"
)

func TestDecreaseUrlServer_Create_OK(t *testing.T) {
	utils.MainLogger = &utils.Logger{Logger: logrus.NewEntry(logrus.StandardLogger())}

	mockCtrl := gomock.NewController(t)
	dbMock := mock_repository.NewMockUrlRepositoryInterface(mockCtrl)

	urlUsecase := usecase.UrlUsecase{DB: dbMock}
	shortServer := DecreaseUrlServer{Usecase: urlUsecase}

	testCase := struct {
		in  string
		err error
	}{
		in:  "http://url.com",
		err: nil,
	}

	dbMock.EXPECT().AddTinyURLBYURL(models.Url{Value: testCase.in}, gomock.Any())
	_, err := shortServer.Create(context.Background(), &server.Url{Value: testCase.in})

	if err != testCase.err {
		t.Error(err)
	}
}

func TestDecreaseUrlServer_Create_AlreadyAdd(t *testing.T) {
	utils.MainLogger = &utils.Logger{Logger: logrus.NewEntry(logrus.StandardLogger())}

	mockCtrl := gomock.NewController(t)
	dbMock := mock_repository.NewMockUrlRepositoryInterface(mockCtrl)

	urlUsecase := usecase.UrlUsecase{DB: dbMock}
	shortServer := DecreaseUrlServer{Usecase: urlUsecase}

	testCase := struct {
		in  string
		out string
		err error
	}{
		in:  "http://url.com",
		out: "http://18jsowm132",
		err: nil,
	}

	dbMock.EXPECT().AddTinyURLBYURL(models.Url{Value: testCase.in}, gomock.Any()).Return(pgx.PgError{Code: "23505"})
	dbMock.EXPECT().GetTinyUrlByUrl(models.Url{Value: testCase.in}).Return(models.Url{Value: strings.TrimLeft(testCase.out, "http://")}, nil)
	url, err := shortServer.Create(context.Background(), &server.Url{Value: testCase.in})

	if err != testCase.err {
		t.Error(err)
	}

	assert.Equal(t, testCase.out, url.GetValue())
}

func TestDecreaseUrlServer_Create_InvalidURL(t *testing.T) {
	utils.MainLogger = &utils.Logger{Logger: logrus.NewEntry(logrus.StandardLogger())}

	mockCtrl := gomock.NewController(t)
	dbMock := mock_repository.NewMockUrlRepositoryInterface(mockCtrl)

	urlUsecase := usecase.UrlUsecase{DB: dbMock}
	shortServer := DecreaseUrlServer{Usecase: urlUsecase}

	testCase := struct {
		in   string
		code codes.Code
		err  string
	}{
		in:   "url",
		code: codes.InvalidArgument,
		err:  "invalid URL",
	}

	_, err := shortServer.Create(context.Background(), &server.Url{Value: testCase.in})
	st, ok := status.FromError(err)
	if !ok {
		t.Fatal(err)
	}

	if st.Code() != testCase.code {
		t.Fatal("Got" + st.Code().String())
	}

	assert.Equal(t, st.Message(), testCase.err)
}

func TestDecreaseUrlServer_Get(t *testing.T) {
	utils.MainLogger = &utils.Logger{Logger: logrus.NewEntry(logrus.StandardLogger())}

	mockCtrl := gomock.NewController(t)
	dbMock := mock_repository.NewMockUrlRepositoryInterface(mockCtrl)

	urlUsecase := usecase.UrlUsecase{DB: dbMock}
	shortServer := DecreaseUrlServer{Usecase: urlUsecase}

	testCase := struct {
		in  string
		out string
		err error
	}{
		in:  "http://18jsowm132",
		out: "http://url.com",
		err: nil,
	}

	dbMock.EXPECT().GetURLByTinyURL(models.Url{Value: strings.TrimLeft(testCase.in, "http://")}).Return(models.Url{Value: testCase.out}, nil)
	url, err := shortServer.Get(context.Background(), &server.TinyUrl{Value: testCase.in})
	if err != nil {
		t.Fatal(err.Error())
	}

	assert.Equal(t, testCase.out, url.GetValue())
}
