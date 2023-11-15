package model

import (
    "github.com/cultofthedamned/thedoomscrolling-go-react/db"
	"github.com/go-sql-driver/mysql"
)

// Article represents a news article
type Article struct {
    ID        int       `json:"id"`
    Title     string    `json:"title"`
    Image     string    `json:"image"`
    Content   string    `json:"content"`
    Author    string    `json:"author"`
    CreatedAt mysql.NullTime `json:"created_at"`
    UpdatedAt mysql.NullTime `json:"updated_at"`
}

func GetAllArticles() ([]Article, error) {
    var articles []Article

    query := "SELECT * FROM news_posts;"

    posts, err := db.DB.Query(query)
    if err != nil {
        return nil, err
    }

    defer posts.Close()

    for posts.Next() {
        var article Article

        err := posts.Scan(
            &article.ID, &article.Title, &article.Image,
            &article.Content, &article.Author,
            &article.CreatedAt, &article.UpdatedAt)

        if err != nil {
            return nil, err
        }

        articles = append(articles, article)
    }

    return articles, nil
}