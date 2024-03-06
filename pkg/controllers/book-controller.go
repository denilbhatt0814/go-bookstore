package controllers

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	"github.com/denilbhatt0814/go-bookstore/pkg/utils"
	"github.com/denilbhatt0814/go-bookstore/pkg/models"
)

var( 
	NewBook models.Book
)

func GetBook(w http.ResponseWriter, r *http.Request){
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	Id, err := strconv.ParseInt(bookId,0,0)
	if err!=nil{
		fmt.Println("error while parsing")
	}

	bookDetails, _:=models.GetBookById(Id)
	res, _:= json.Marshal(bookDetails)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)	
}

func CreateBook(w http.ResponseWriter, r *http.Request){
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	ID, err:= strconv.ParseInt(bookId, 0,0)
	if err!=nil{
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	// creating empty book struct
	var UpdateBook = &models.Book{} 
	
	// Parsing the request and filling info in the empty book struct(UpdateBook)
	utils.ParseBody(r, UpdateBook)

	// Parsing bookId
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err!=nil{
		fmt.Println("error while parsing")
	}

	// Searching the book in DB 
	bookDetails, db := models.GetBookById(ID)

	// Verifying and creating update values
	if UpdateBook.Name != ""{
		bookDetails.Name = UpdateBook.Name
	}
	if UpdateBook.Author != ""{
		bookDetails.Author = UpdateBook.Author
	}
	if UpdateBook.Publication != ""{
		bookDetails.Publication = UpdateBook.Publication
	}

	// Now saving updation on DB
	db.Save(&bookDetails)

	// Sending the success response
	res, _:= json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}