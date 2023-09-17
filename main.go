// main.go
package main

import (
	"fmt"
	"log"
	"mini-online-store/cmd"
	config "mini-online-store/config/database"
	"mini-online-store/db/migrations"

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

	cmd.SetupRouters(router, db)
	migrations.Run(db)

	port := os.Getenv("app_port")
	router.Run(fmt.Sprintf(":%s", port))
}
