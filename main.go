package main

import (
	"fmt"
	"log"

	"github.com/Rickykn/article-api/database"
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

	err := r.Run()
	if err != nil {
		panic(err)
	}

}
