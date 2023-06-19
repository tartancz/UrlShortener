package main

import (
	"net/http"
)

func (app *application) getRoutes() http.Handler {
	router := http.NewServeMux()

	basicChain := chainer(app.logRequest)

	fileServer := http.FileServer(http.Dir("./static/"))

	router.Handle("/static/", basicChain(http.StripPrefix("/static", fileServer)))

	standard := chainer(app.recoverPanic, basicChain, secureHeaders)
	_ = standard
	router.Handle("/", standard(http.HandlerFunc(app.home)))
	router.Handle("/URL/", standard(http.HandlerFunc(app.redirectUser)))

	return router
}

func chainer(chain ...func(http.Handler) http.Handler) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		for i := 0; i < len(chain); i++ {
			next = chain[i](next)
		}
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
	}
}