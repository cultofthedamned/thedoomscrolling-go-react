package handlers

import (
    "fmt"
    "encoding/json"
    "net/http"
    "github.com/cultofthedamned/thedoomscrolling-go-react/model"
)

func ArticlesHandler(w http.ResponseWriter, r *http.Request) {
    articles, err := model.GetAllArticles()

    if err != nil {
        fmt.Println("Error fetching data:", err)
        http.Error(w, "Error fetching data", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(articles)
}