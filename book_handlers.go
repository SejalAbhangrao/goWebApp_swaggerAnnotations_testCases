package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Booklist struct {
	BookName string `json:"bookname"`
	Summary  string `json:"summary"`
}

var books []Booklist

func getBookHandler(w http.ResponseWriter, r *http.Request) {

	// swagger:operation GET /book getBookHandler GetBookHandler
	//
	// Returns the name of the book and its summaries
	// ---
	// consumes:
	// - text/plain
	// produces:
	// - text/plain
	// parameters:
	// - name: name
	//   in: path
	//   description: List of books along with summaries to be returned.
	//   required: true
	//   type: string
	// responses:
	//   '200':
	//     description: The hello message
	//     type: string

	//Convert the "books" variable to json
	birdListBytes, err := json.Marshal(books)

	// If there is an error, print it to the console, and return a server
	// error response to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// If all goes well, write the JSON list of books to the response
	w.Write(birdListBytes)
}

func createBookHandler(w http.ResponseWriter, r *http.Request) {

	// swagger:operation POST /book createBookHandler CreateBookHandler
	//
	// Returns the name of the book and its summaries
	// ---
	// consumes:
	// - text/plain
	// produces:
	// - text/plain
	// parameters:
	// - name: name
	//   in: path
	//   description: List of books along with summaries to be returned.
	//   required: true
	//   type: string
	// responses:
	//   '200':
	//     description: The hello message
	//     type: string

	// Create a new instance of Book
	newBook := Booklist{}

	// We send all our data as HTML form data
	// the `ParseForm` method of the request, parses the
	// form values
	err := r.ParseForm()

	// In case of any error, we respond with an error to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get the information about the book from the form info
	newBook.BookName = r.Form.Get("bookname")
	newBook.Summary = r.Form.Get("summary")

	// Append our existing list of books with a new entry
	books = append(books, newBook)

	//Finally, we redirect the user to the original HTMl page (located at `/assets/`)
	http.Redirect(w, r, "/assets/", http.StatusFound)
}
