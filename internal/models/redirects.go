package models

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/lib/pq"
)

type Redirect struct {
	ID         int
	Url        string
	ShortenUrl string
	Created    time.Time
}

type RedirectModel struct {
	DB *sql.DB
}

func (m *RedirectModel) Insert(url, ShortenUrl string) (int, error) {
	stmt := `INSERT INTO redirects (url, shorten_url) VALUES ($1, $2) RETURNING id;`
	var id int
	row := m.DB.QueryRow(stmt, url, ShortenUrl)
	err := row.Scan(&id)
	if err != nil {
		var postSQLerror *pq.Error
		if errors.As(err, &postSQLerror) {
			if postSQLerror.Code == "23505" && strings.Contains(postSQLerror.Message, "redirects_shorten_url") {
				return 0, ErrDuplicateShortenUrl
			}
		}
		return 0, err
	}
	return id, nil
}

func (m *RedirectModel) GetUrl(ShortenUrl string) (string, error) {
	stmt := `SELECT url from redirects where shorten_url=$1`
	var url string
	row := m.DB.QueryRow(stmt, ShortenUrl)
	if err := row.Scan(&url); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", ErrNoRecord
		}
		return "", err
	}
	return url, nil
}
