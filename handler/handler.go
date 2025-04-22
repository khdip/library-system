package handler

import (
	"html/template"

	"github.com/jmoiron/sqlx"
)

type Books struct {
	ID       int    `db:"id" json:"id"`
	BookName string `db:"book_name" json:"book_name"`
	Author   string `db:"author" json:"author"`
	BookDesc string `db:"book_description" json:"book_description"`
}

type Handler struct {
	templates *template.Template
	db        *sqlx.DB
}

func GetHandler(db *sqlx.DB) *Handler {
	hand := &Handler{
		db: db,
	}
	hand.GetTemplate()
	return hand
}

func (h *Handler) GetTemplate() {
	h.templates = template.Must(template.ParseFiles("templates/create-book.html", "templates/list-book.html", "templates/edit-book.html", "templates/search-result.html"))
}
