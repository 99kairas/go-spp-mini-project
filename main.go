package main

import (
	"go-spp/configs"
	"go-spp/middlewares"
	"go-spp/routes"
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
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

	routes.Routes(e, db)

	middlewares.LogMiddleware(e)

	e.Logger.Fatal(e.Start(":8080"))
}
