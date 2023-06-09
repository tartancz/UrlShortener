package main

import (
	"html/template"
	"path/filepath"
)

type templateData struct {
	Form any
	FullURL string
}

func newTemplateData() *templateData {
	return &templateData{}
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.html")
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).ParseFiles("./ui/html/base.html", page)
		if err != nil {
			return nil, err
		}
		cache[name] = ts

	}
	return cache, nil
}
