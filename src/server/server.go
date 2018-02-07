package main

import (
	"log"
	"net/http"
	"database/sql"
	"io/ioutil"
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

//Função que checa se houve um erro e faz display da mensagem
func checkError(err error, message string, w http.ResponseWriter) {
	if err != nil {
		log.Println(err)
		w.Write([]byte(message))
	}
}

//Função que faz a requisição de todos os livros no BD
func getAllBooks(w http.ResponseWriter, r *http.Request) {
	var books []*Book //Cria um array de livros, a serem transformados em JSON

	rows, err_query := database.Query("SELECT * FROM book") //Faz a requisição ao banco de dados

	//Caso haja um erro, as operações não podem continuar
	checkError(err_query, "Error making request to database", w)

	for rows.Next() { //Percorre cada linha da resposta e insere no array de livros
		book := new (Book)
		
		err_scan := rows.Scan(&book.Id, &book.Name, &book.Description, &book.Author, &book.Publisher) //Lê os campos da resposta

		checkError(err_scan, "Error scanning rows to type Book", w)

		books = append(books, book) //Insere o livro no array
	}

	booksJson, err_json := json.Marshal(books) //Serializa o array em JSON

	//Caso haja um erro, como type mismatch, a operação deve ser cancelada
	checkError(err_json, "Error parsing data type 'Book' to json", w)

	w.Write(booksJson) //Devolve a resposta
}

//Função que faz a requisição de um livro único no BD
func getSingleBook(w http.ResponseWriter, r *http.Request) {
	book := new (Book) //Cria um novo livro
	vars := mux.Vars(r) //Resgata as variáveis enviadas na request

	rows, err_query := database.Query("SELECT * FROM book WHERE id = " + vars["id"]) //Faz a requisição ao banco de dados

	checkError(err_query, "Error making request to database", w)

	for rows.Next() {
		err_scan := rows.Scan(&book.Id, &book.Name, &book.Description, &book.Author, &book.Publisher) //Lê os campos da resposta

		checkError(err_scan, "Error scanning rows to type Book", w)
	}

	if book.Id == 0 { //Se o id do livro é igual a zero, quer dizer que não houveram respostas
		w.Write([]byte("Não existem livros na base de dados"))
		return
	}

	bookJson, err_json := json.Marshal(book) //Serializa o livro em JSON

	//Caso haja um erro, como type mismatch, a operação deve ser cancelada
	checkError(err_json, "Error parsing data type 'Book' to json", w)


	w.Write(bookJson) //Retorna a resposta
}  

//Função que insere um livro no banco de dados
func postSingleBook(w http.ResponseWriter, r *http.Request) {
	book := new (Book) //Novo livro a ser adicionado
	body, err_reading := ioutil.ReadAll(r.Body) //Resgata as informações da request

	checkError(err_reading, "Error reading request body.", w) 

	err_json := json.Unmarshal(body, &book) //Transforma as informações da request em tipo Book

	checkError(err_json, "Error parsing body to Json", w)

	statement, err_prepare := database.Prepare("INSERT book SET `name` = ?, `description` = ?, `author` = ?, `publisher` = ?;") //Prepara uma request ao banco de dados

	checkError(err_prepare, "Error preparing insert statement", w) 

	_, err_insert := statement.Exec(book.Name, book.Description, book.Author, book.Publisher) //Executa a inserção

	checkError(err_insert, "Error executing insertion on Database", w)
}

func main() {
	router := mux.NewRouter() //Cria um roteador de requisições

	database, _ = sql.Open("mysql", "root:Rootpass.19@/BookSchema") //Cria uma conexão com o banco

	//Configura-se dois endpoints serem acessados por outras aplicações
	router.HandleFunc("/books", getAllBooks).Methods("GET") 
	router.HandleFunc("/book/{id}", getSingleBook).Methods("GET")
	router.HandleFunc("/book", postSingleBook).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}