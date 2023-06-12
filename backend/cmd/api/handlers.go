package main

import (
	"fmt"
	"net/http"
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

func (app *application) homeGet(w http.ResponseWriter, r *http.Request) {
	app.render(w, http.StatusOK, "home.html")
}

func (app *application) homePost(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	shortUrl := r.FormValue("shortUrl")
	fmt.Println(url, shortUrl)
}
