package main

import (
	"log"
	"net/http"
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
)

var database *sql.DB //Variável que guarda a conexão com o banco de dados

//Modelo da tabela "book" no banco de dados
type Book struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Author string `json:"author"`
	Publisher string `json:"publisher"`
}

//Função que faz a requisição de todos os livros na DB
func getAllBooks(w http.ResponseWriter, r *http.Request) {
	var books []*Book //Cria um array de livros, a serem transformados em JSON

	rows, err := database.Query("SELECT * FROM book") //Faz a requisição ao banco de dados

	//Caso haja um erro, as operações não podem continuar
	if err != nil {
		log.Println(err)
		return
	}

	for rows.Next() { //Percorre cada linha da resposta e insere no array de livros
		book := new (Book)
		
		err = rows.Scan(&book.Id, &book.Name, &book.Description, &book.Author, &book.Publisher) //Lê os campos da resposta

		if err != nil {
			log.Println(err)
			return
		}

		books = append(books, book) //Insere o livro no array
	}

	booksJson, err := json.Marshal(books) //Serializa o array em JSON

	//Caso haja um erro, como type mismatch, a operação deve ser cancelada
	if err != nil {
		log.Println(err)
		return
	}

	w.Write(booksJson) //Devolve a resposta
}

func getSingleBook(w http.ResponseWriter, r *http.Request) {

}

func postSingleBook(w http.ResponseWriter, r *http.Request) {

}

func main() {
	router := mux.NewRouter() //Cria um roteador de requisições

	database, _ = sql.Open("mysql", "root:Mercenary.1@/BookSchema") //Cria uma conexão com o banco

	//Configura-se dois endpoints serem acessados por outras aplicações
	router.HandleFunc("/books", getAllBooks).Methods("GET") 
	router.HandleFunc("/book/{id}/", getSingleBook).Methods("GET")
	router.HandleFunc("/book", postSingleBook).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}