package main

import (
	"net/http"
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/go-sql-driver/mysql"
	"encodin/json"
)

type Book struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Author string `json:"author"`
	Publisher string `json:"publisher"`
	PublishDate string `json:"publishDate"`
	AddDate string `json:"addDate"`
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {

}

func getSingleBook(w http.ResponseWriter, r *http.Request) {

}

func postSingleBook(w http.ResponseWriter, r *http.Request) {

}

func main() {
	router := mux.NewRouter() //Cria um roteador de requisições

	//Configura-se dois endpoints serem acessados por outras aplicações
	router.HandleFunc("/books", getAllBooks).Methods("GET") 
	router.HandleFunc("/book/{id}/", getSingleBook).Methods("GET")
	router.HandleFunc("/book", postSingleBook).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}