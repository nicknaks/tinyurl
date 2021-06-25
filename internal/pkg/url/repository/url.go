package repository

import (
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"tinyUrl/internal/pkg/models"
)

type UrlRepositoryInterface interface {
	AddTinyURLBYURL(url models.Url, tinyUrl models.Url) error
	GetTinyUrlByUrl(url models.Url) (tinyUrl models.Url, err error)
	GetURLByTinyURL(tinyUrl models.Url) (url models.Url, err error)
}

type UrlRepository struct {
	DB *sqlx.DB
}

func (u UrlRepository) AddTinyURLBYURL(url models.Url, tinyUrl models.Url) error {
	_, err := u.DB.Exec(`INSERT INTO urls (url, tinyurl) VALUES ($1, $2)`, url.Value, tinyUrl.Value)
	return err
}

func (u UrlRepository) GetTinyUrlByUrl(url models.Url) (tinyUrl models.Url, err error) {
	err = u.DB.QueryRowx(`SELECT tinyurl FROM urls WHERE url = $1`, url.Value).Scan(&tinyUrl.Value)
	return tinyUrl, err
}

func (u UrlRepository) GetURLByTinyURL(tinyUrl models.Url) (url models.Url, err error) {
	err = u.DB.QueryRowx(`SELECT url FROM urls WHERE tinyurl = $1`, tinyUrl.Value).Scan(&url.Value)
	return url, err
}
