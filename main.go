// Testing go-swagger generation
//
// The purpose of this application is to test go-swagger in a simple GET request.
//
//     Title: Read Book List
//     Version: 0.0.1
//     Description: A web application in Go to keep track of books read
//     Schemes: http
//     Host: localhost:8080
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Daniel<danielfs.ti@gmail.com>
//     Consumes:
//     - text/plain
//
//     Produces:
//     - text/plain
//
// swagger:meta

package main

import (
	"fmt"
	"net/http"

	_ "github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	_ "github.com/swaggo/gin-swagger"

	_ "github.com/swaggo/gin-swagger/swaggerFiles"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")

	// Declare the static file directory and point it to the directory we just made
	staticFileDirectory := http.Dir("./assets/")
	// Declare the handler, that routes requests to their respective filename.
	// The fileserver is wrapped in the `stripPrefix` method, because we want to
	// remove the "/assets/" prefix when looking for files.
	// For example, if we type "/assets/index.html" in our browser, the file server
	// will look for only "index.html" inside the directory declared above.
	// If we did not strip the prefix, the file server would look for "./assets/assets/index.html", and yield an error
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	// The "PathPrefix" method acts as a matcher, and matches all routes starting
	// with "/assets/", instead of the absolute route itself
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	r.HandleFunc("/book", getBookHandler).Methods("GET")
	r.HandleFunc("/book", createBookHandler).Methods("POST")
	return r
}

func main() {
	// The router is now formed by calling the `newRouter` constructor function
	// that we defined above. The rest of the code stays the same
	r := newRouter()
	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
