package main

import (
	"fmt"
	"log"
	"os"

	db "github.com/SssHhhAaaDddOoWww/Goiler/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	db.Connect()
	router := gin.Default()
	Routes(router)
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}
	log.Fatal(router.Run(":" + port))
}

func Routes(router *gin.Engine) {
	router.GET("/health")
}
