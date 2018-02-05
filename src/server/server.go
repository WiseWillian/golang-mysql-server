package main

import (
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

func main() {

}