package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("failed to load .env")
	}

	// db := configs.InitDB()
	e := echo.New()

	e.Logger.Fatal(e.Start(":8080"))
}
