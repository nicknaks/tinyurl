package usecase

import (
	_ "github.com/jackc/pgx"
	"math/rand"
	"strings"
	"tinyUrl/internal/pkg/models"
	"tinyUrl/internal/pkg/url/repository"
	"tinyUrl/internal/tinyUrl/utils"
)

type UrlUsecaseInterface interface {
	CreateTinyUrl() models.Url
	CreateURL(url models.Url) (tinyUrl models.Url, err error)
	GetUrl(tinyUrl models.Url) (url models.Url, err error)
}

type UrlUsecase struct {
	DB repository.UrlRepositoryInterface
}

func (u UrlUsecase) GetUrl(tinyUrl models.Url) (url models.Url, err error) {
	return u.DB.GetURLByTinyURL(tinyUrl)
}

func (u UrlUsecase) CreateURL(url models.Url) (tinyUrl models.Url, err error) {
	tinyUrl = u.CreateTinyUrl()

	err = u.DB.AddTinyURLBYURL(url, tinyUrl)

	if utils.PgxErrorCode(err) != "23505" {
		utils.MainLogger.LogError(err)
		return models.Url{}, err
	}

	if err != nil {
		tinyUrl, err = u.DB.GetTinyUrlByUrl(url)
		if err != nil {
			utils.MainLogger.LogError(err)
			return models.Url{}, err
		}
	}

	tinyUrl.Value = "http://" + tinyUrl.Value
	return tinyUrl, nil
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func (u UrlUsecase) CreateTinyUrl() models.Url {
	n := 10
	sb := strings.Builder{}
	sb.Grow(n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	utils.MainLogger.LogInfo("Generate random short url")
	return models.Url{Value: sb.String()}
}
