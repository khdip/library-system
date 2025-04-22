package handler

import (
	"net/http"
)

type SearchedBookList struct {
	Searched_Book_list []Books
}

func (h *Handler) SearchBook(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	books := []Books{}
	const searchQuery = "SELECT * FROM books WHERE book_name ILIKE '%%' || $1 || '%%' OR author ILIKE '%%' || $1 || '%%'"
	h.db.Select(&books, searchQuery, r.FormValue("Search"))
	slt := SearchedBookList{Searched_Book_list: books}
	err = h.templates.ExecuteTemplate(w, "search-result.html", slt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
