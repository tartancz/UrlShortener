package models

import (
	"database/sql"
	"time"
)

type Redirect struct {
	ID      int
	Url     string
	ShortenUrl string
	Created time.Time
}

type RedirectModel struct {
	DB *sql.DB
}

