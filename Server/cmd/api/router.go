package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) router() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/status", app.statusHelper)

	router.HandlerFunc(http.MethodGet, "/v1/book/:id", app.getOneBook)
	router.HandlerFunc(http.MethodGet, "/v1/books", app.getBooks)

	return router
}
