package main

import "net/http"

func (app *aplication) getRoutes() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("/", app.home)
	
	return router
}