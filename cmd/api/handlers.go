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
	URL      string
	ShortURL string
	validator.Validator
}

func (app *application) homeGet(w http.ResponseWriter, r *http.Request) {
	form := homeForm{
		ShortURL: uuid.NewString(),
	}
	data := newTemplateData()
	data.Form = form
	data.FullURL = fmt.Sprintf("%s/URL/%s", r.Host, form.ShortURL)
	app.render(w, http.StatusOK, "home.html", data)
}

func (app *application) homePost(w http.ResponseWriter, r *http.Request) {
	form := homeForm{
		URL:      r.FormValue("URL"),
		ShortURL: r.FormValue("ShortURL"),
	}

	//URL
	form.CheckField(validator.ValidUrl(form.URL), "url", "This is not valid URL, please insert valid url")
	form.CheckField(validator.MaxLenght(form.URL, 2048), "url", "Url is too long, maximum of characters is 2048")
	//shortUrl
	formattedShortUrl := fmt.Sprintf("%s/URL/%s", r.Host, form.ShortURL)
	form.CheckField(validator.ValidUrl(formattedShortUrl), "ShortURL", "This is not valid URL, please insert valid url")
	form.CheckField(validator.MinLenght(form.ShortURL, 12), "ShortURL", "Url is too short, minimum of characters is 12")

	data := newTemplateData()

	if !form.Valid() {
		data.Form = form
		data.FullURL = fmt.Sprintf("%s/URL/%s", r.Host, form.ShortURL)
		app.render(w, http.StatusBadRequest, "home.html", data)
		return
	}
	_, err := app.redirects.Insert(form.URL, form.ShortURL)
	if err != nil {
		if errors.Is(err, models.ErrDuplicateShortenUrl) {
			form.AddFieldError("ShortURL", "this URL already exist, please enter new one")
			data.Form = form
			data.FullURL = fmt.Sprintf("%s/URL/%s", r.Host, form.ShortURL)
			app.render(w, http.StatusBadRequest, "home.html", data)
		} else {
			app.serverError(w, err)
		}
		return
	}

	data.FullURL = fmt.Sprintf("%s/URL/%s", r.Host, form.ShortURL)
	app.render(w, http.StatusCreated, "createdURL.html", data)
}

func (app *application) redirectUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	shortUrl := strings.TrimPrefix(r.URL.Path, "/URL/")
	fmt.Println(shortUrl)
}
