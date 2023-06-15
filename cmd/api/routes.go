package main

import "net/http"

func (app *application) getRoutes() http.Handler {
	router := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./static/"))

	router.Handle("/static/", http.StripPrefix("/static", fileServer))

	router.HandleFunc("/", app.home)
	router.HandleFunc("/URL/", app.redirectUser)

	return router
}
