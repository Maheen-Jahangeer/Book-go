package main

import (
	"backend/models"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getOneBook(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))

	if err != nil {
		app.logger.Println("Error :", err)
		app.jsonError(w, http.StatusBadRequest, err)
	}

	newBook := models.Book{
		ID:          id,
		BookName:    "The book of names",
		ImagUrl:     "https://i.gr-assets.com/images/S/compressed.photo.goodreads.com/books/1311999970l/172763.jpg",
		Discription: "The new book for demo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	app.jsonParser(w, http.StatusOK, newBook, "book")
}

func (app *application) getBooks(w http.ResponseWriter, r *http.Request) {

}
