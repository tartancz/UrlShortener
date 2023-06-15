package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/tartancz/UrlShortener/internal/models"
	"github.com/tartancz/UrlShortener/internal/validator"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	switch r.Method {
	case "GET":
		app.homeGet(w, r)
		return
	case "POST":
		app.homePost(w, r)
		return
	}
	app.clientError(w, http.StatusMethodNotAllowed)
}

type homeForm struct {
	URL        string
	ShortenUrl string
	validator.Validator
}

func (app *application) homeGet(w http.ResponseWriter, r *http.Request) {
	form := homeForm{
		ShortenUrl: uuid.NewString(),
	}
	data := newTemplateData()
	data.Form = form
	data.FullURL = fmt.Sprintf("%s/URL/%s", r.Host, form.ShortenUrl)
	app.render(w, http.StatusOK, "home.html", data)
}

func (app *application) homePost(w http.ResponseWriter, r *http.Request) {
	form := homeForm{
		URL:        r.FormValue("URL"),
		ShortenUrl: r.FormValue("ShortenUrl"),
	}

	//URL
	form.CheckField(validator.ValidUrl(form.URL), "url", "This is not valid URL, please insert valid url")
	form.CheckField(validator.MaxLenght(form.URL, 2048), "url", "Url is too long, maximum of characters is 2048")
	//ShortenUrl
	formattedShortUrl := fmt.Sprintf("%s/URL/%s", r.Host, form.ShortenUrl)
	form.CheckField(validator.ValidUrl(formattedShortUrl), "ShortenUrl", "This is not valid URL, please insert valid url")
	form.CheckField(validator.MinLenght(form.ShortenUrl, 12), "ShortenUrl", "Url is too short, minimum of characters is 12")

	data := newTemplateData()

	if !form.Valid() {
		data.Form = form
		data.FullURL = fmt.Sprintf("%s/URL/%s", r.Host, form.ShortenUrl)
		app.render(w, http.StatusBadRequest, "home.html", data)
		return
	}
	_, err := app.redirects.Insert(form.URL, form.ShortenUrl)
	if err != nil {
		if errors.Is(err, models.ErrDuplicateShortenUrl) {
			form.AddFieldError("ShortenUrl", "this URL already exist, please enter new one")
			data.Form = form
			data.FullURL = fmt.Sprintf("%s/URL/%s", r.Host, form.ShortenUrl)
			app.render(w, http.StatusBadRequest, "home.html", data)
		} else {
			app.serverError(w, err)
		}
		return
	}

	data.FullURL = fmt.Sprintf("%s/URL/%s", r.Host, form.ShortenUrl)
	app.render(w, http.StatusCreated, "createdURL.html", data)
}

func (app *application) redirectUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	ShortenUrl := strings.TrimPrefix(r.URL.Path, "/URL/")
	url, err := app.redirects.GetUrl(ShortenUrl)
	if err != nil {
		if errors.As(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	http.Redirect(w, r, url, http.StatusMovedPermanently)
}
