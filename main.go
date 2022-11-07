package main

import (
	"fmt"
	"log"

	"github.com/Rickykn/article-api/database"
	"github.com/Rickykn/article-api/handlers"
	"github.com/Rickykn/article-api/repositories"
	"github.com/Rickykn/article-api/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("E-Wallet backend")

	errEnv := godotenv.Load()

	errConnect := database.Connect()

	if errConnect != nil {
		panic(errConnect)
	}

	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	ar := repositories.NewArticleRepository(&repositories.ARConfig{
		DB: database.Get(),
	})

	as := services.NewArticleService(&services.ASConfig{
		ArticleRepository: ar,
	})

	h := handlers.New(&handlers.HandlerConfig{
		ArticleService: as,
	})

	article := r.Group("/article")
	{
		article.POST("/", h.CreateArticle)
	}

	err := r.Run()
	if err != nil {
		panic(err)
	}

}
