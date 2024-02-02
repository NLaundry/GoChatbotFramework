package main

import (
    // "context"
    "fmt"

    "net/http"
    "github.com/a-h/templ"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", ComponentHandler(home("Bob")))
    http.ListenAndServe(":8080", r)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

// // creates a function that matches the HandlerFunc signature and calls the component's Render method
func ComponentHandler(component templ.Component) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		component.Render(r.Context(), w)
	}
}

