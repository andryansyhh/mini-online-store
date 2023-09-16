// main.go
package main

import (
	"fmt"
	"log"
	config "mini-online-store/config/database"
	"mini-online-store/internal/routers"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	router := gin.Default()
	db := config.GetDatabase()

	routers.SetupRoutes(router, db)

	port := os.Getenv("app_port")
	router.Run(fmt.Sprintf(":%s", port))
}
