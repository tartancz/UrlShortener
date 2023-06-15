package models

import "errors"

var (
	ErrDuplicateShortenUrl = errors.New("models: duplicate shorten_url")

	ErrNoRecord = errors.New("models: no matching record found")
)
