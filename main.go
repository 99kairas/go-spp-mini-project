package main

import (
	"fmt"
	"go-spp/configs"
	"go-spp/middlewares"
	"go-spp/routes"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("failed to load .env")
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")
	time.Local = loc

	db := configs.InitDB()
	e := echo.New()

	server := os.Getenv("SERVER_PORT")

	routes.Routes(e, db)

	middlewares.LogMiddleware(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", server)))
}
