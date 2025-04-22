package handler

import (
	"net/http"
)

type FormData struct {
	Book   Books
	Errors map[string]string
}

func (h *Handler) CreateBook(w http.ResponseWriter, r *http.Request) {
	ErrorValue := map[string]string{}
	book := Books{}
	h.LoadCreateForm(w, book, ErrorValue)
}

func (h *Handler) StoreBook(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bookName := r.FormValue("Book")
	author := r.FormValue("Author")
	bookDesc := r.FormValue("Description")
	book := Books{
		BookName: bookName,
		Author:   author,
		BookDesc: bookDesc,
	}
	if bookName == "" {
		ErrorValue := map[string]string{
			"Error": "Book Name field can not be empty.",
		}
		h.LoadCreateForm(w, book, ErrorValue)
		return
	} else if len(bookName) < 3 {
		ErrorValue := map[string]string{
			"Error": "Book Name field should have atleast 3 characters",
		}
		h.LoadCreateForm(w, book, ErrorValue)
		return
	}
	const insertBook = `INSERT INTO books(book_name, author, book_description) VALUES($1, $2, $3);`
	res := h.db.MustExec(insertBook, bookName, author, bookDesc)
	ok, err := res.RowsAffected()
	if err != nil || ok == 0 {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func (h *Handler) EditBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/edit/"):]
	if id == "" {
		http.Error(w, "Invalid URL", http.StatusInternalServerError)
		return
	}

	const getBook = `SELECT * FROM books WHERE id=$1`
	var book Books
	h.db.Get(&book, getBook, id)

	if book.ID == 0 {
		http.Error(w, "Invalid URL", http.StatusInternalServerError)
		return
	}
	h.LoadEditForm(w, book, map[string]string{})
}

func (h *Handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/update/"):]
	if id == "" {
		http.Error(w, "Invalid URL", http.StatusInternalServerError)
		return
	}

	const getBook = `SELECT * FROM books WHERE id=$1`
	var book Books
	h.db.Get(&book, getBook, id)

	if book.ID == 0 {
		http.Error(w, "Invalid URL", http.StatusInternalServerError)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newBook := r.FormValue("Book")
	newAuthor := r.FormValue("Author")
	newBookDesc := r.FormValue("Description")
	book.BookName = newBook
	book.Author = newAuthor
	book.BookDesc = newBookDesc
	if newBook == "" {
		ErrorValue := map[string]string{
			"Error": "Book Name field can not be empty.",
		}
		h.LoadEditForm(w, book, ErrorValue)
		return
	} else if len(newBook) < 3 {
		ErrorValue := map[string]string{
			"Error": "Book Name field should have atleast 3 characters",
		}
		h.LoadEditForm(w, book, ErrorValue)
		return
	}

	const updateTodo = `UPDATE books SET book_name = $2, author = $3, book_description = $4 WHERE id=$1`
	res := h.db.MustExec(updateTodo, id, newBook, newAuthor, newBookDesc)
	ok, err := res.RowsAffected()
	if err != nil || ok == 0 {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func (h *Handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/delete/"):]
	if id == "" {
		http.Error(w, "Invalid URL", http.StatusInternalServerError)
		return
	}

	const getBook = `SELECT * FROM books WHERE id=$1`
	var book Books
	h.db.Get(&book, getBook, id)

	if book.ID == 0 {
		http.Error(w, "Invalid URL", http.StatusInternalServerError)
		return
	}
	const deleteBook = `DELETE FROM books WHERE id=$1`
	res := h.db.MustExec(deleteBook, id)
	ok, err := res.RowsAffected()
	if err != nil || ok == 0 {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

// Form Validation
func (h *Handler) LoadCreateForm(w http.ResponseWriter, book Books, myErrors map[string]string) {
	form := FormData{
		Book:   book,
		Errors: myErrors,
	}

	err := h.templates.ExecuteTemplate(w, "create-book.html", form)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) LoadEditForm(w http.ResponseWriter, book Books, myErrors map[string]string) {
	form := FormData{
		Book:   book,
		Errors: myErrors,
	}

	err := h.templates.ExecuteTemplate(w, "edit-book.html", form)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
