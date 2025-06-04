package main

import (
	"handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/books", handlers.Get_books).Methods("GET")
	r.HandleFunc("/api/book/{id}", handlers.Get_book_by_id).Methods("GET")
	r.HandleFunc("/api/book", handlers.Create_book).Methods("POST")
	r.HandleFunc("/api/book/{id}", handlers.Update_book).Methods("PUT")
	r.HandleFunc("/api/book/{id}", handlers.Delete_book).Methods("DELETE")

	log.Fatal(http.ListenAndServe("my_host:8000", r))
}
