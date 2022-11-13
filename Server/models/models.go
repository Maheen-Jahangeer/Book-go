package models

import "time"

type Book struct {
	ID          int         `json:"id"`
	BookName    string      `json:"book_name"`
	Discription string      `json:"discription"`
	ImagUrl     string      `json:"image_url"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	BookGenre   []BookGenre `json:"book_genre"`
}

type Genre struct {
	ID        int       `json:"id"`
	GenreName string    `json:"genre_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BookGenre struct {
	ID        int       `json:"id"`
	GenreId   int       `json:"genre_id"`
	Genre     Genre     `json:"genre"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
