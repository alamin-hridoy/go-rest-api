package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	_"database/sql"
	_"github.com/lib/pq"
	_"github.com/subosito/gotenv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     int    `json:id`
	Title  string `json:title`
	Author string `json:author`
	Year   string `json:year`
}

var books []Book

func main() {
	router := mux.NewRouter()

	books=append(books,Book{ID:1,Title: "Golang Pointers",Author:"Mr.golang",Year:"2021"},
	Book{ID:2,Title: "Golang 2",Author:"Mr.golang2",Year:"2021"},
	Book{ID:3,Title: "Golang 3",Author:"Mr.golang3",Year:"2021"},
	Book{ID:4,Title: "Golang 4",Author:"Mr.golang4",Year:"2021"},
	Book{ID:5,Title: "Golang 5",Author:"Mr.golang5",Year:"2021"},

)
	router.HandleFunc("/books",getBooks).Methods("GET")
	router.HandleFunc("/books/{id}",getBook).Methods("GET")
	router.HandleFunc("/books",addBook).Methods("POST")
	router.HandleFunc("/books",updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}",removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8870",router))
}


func getBooks(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request){
	params:=mux.Vars(r)
	i,_:=strconv.Atoi(params["id"])
	
	for _,book:=range books{
		if book.ID==i{
			json.NewEncoder(w).Encode((&book))
		}
	}
}

func addBook(w http.ResponseWriter, r *http.Request){
	var book Book
	_=json.NewDecoder(r.Body).Decode(&book)
	books=append(books, book)
	json.NewEncoder(w).Encode(books)
}

func updateBook(w http.ResponseWriter, r *http.Request){
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	for i,item:=range books{
		if item.ID==book.ID{
			books[i]=book
		}
	}
	json.NewEncoder(w).Encode(books)
}

func removeBook(w http.ResponseWriter, r *http.Request){
	params:=mux.Vars(r)
	id,_:= strconv.Atoi(params["id"])

	for i,item := range books{
		if item.ID == id{
			books = append(books[:i], books[i+1:]... )
		}
	}
	json.NewEncoder(w).Encode(books)
}