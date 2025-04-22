package main

import (
	"fmt"
	"log"
	"net/http"

	"practice/library-system/handler"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	var schema = `
			CREATE TABLE IF NOT EXISTS books (
				id serial,
				book_name text,
				author text,
				book_description text,

				primary key(id)
			);`

	db, err := sqlx.Connect("postgres", "user=postgres password=Anubis0912 dbname=library sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	db.MustExec(schema)

	h := handler.GetHandler(db)
	http.HandleFunc("/", h.GetBooks)
	http.HandleFunc("/create", h.CreateBook)
	http.HandleFunc("/store", h.StoreBook)
	http.HandleFunc("/search", h.SearchBook)
	http.HandleFunc("/edit/", h.EditBook)
	http.HandleFunc("/Update/", h.UpdateBook)
	http.HandleFunc("/delete/", h.DeleteBook)
	fmt.Println("Server Starting...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server Not Found", err)
	}
}
