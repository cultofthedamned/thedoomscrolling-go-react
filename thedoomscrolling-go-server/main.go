package main

import (
    "fmt"
    "net/http"
    "github.com/cultofthedamned/thedoomscrolling-go-react/handlers"
    "github.com/cultofthedamned/thedoomscrolling-go-react/db"
	"github.com/gorilla/mux"
)

func main() {
    db.InitDB()

	myRouter := mux.NewRouter().StrictSlash(true)

    myRouter.HandleFunc("/", handlers.ArticlesHandler)

    fmt.Println("Server listening on :8080")
    http.ListenAndServe(":8080", myRouter)
}