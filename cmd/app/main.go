package main

import (
	"InfotecsTestCase/database"
	getenv "InfotecsTestCase/internal/get_env"
	"InfotecsTestCase/server"
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load("../../configs/.env"); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	err := connectToDatabase()
	if err != nil {
		return err
	}
	defer database.DB.Close(context.Background())

	err = database.CreateBaseTables()
	if err != nil {
		return err
	}

	router := startRouter()

	server.AllRequests(router)

	router.Run(":8080")
	return nil
}

func connectToDatabase() error {
	dbUrl := getenv.GetEnv("DATABASE_URL")

	err := database.DBConnectionInit(dbUrl)
	if err != nil {
		return err
	}

	err = database.DB.Ping(context.Background())
	if err != nil {
		return fmt.Errorf("Error while Ping db connection: %v", err)
	}

	return nil
}

func startRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	return router
}
