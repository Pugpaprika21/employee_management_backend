package main

import (
	"log"
	"os"

	"github.com/Pugpaprika21/db"
	"github.com/Pugpaprika21/router"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	dbc, err := db.New().UseDB()
	if err != nil {
		panic(err)
	}

	migrate := db.NewMigration(dbc)
	migrate.Run()

	router := router.NewServerRouter()
	router.Setup(e, dbc)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	e.Logger.Fatal(e.Start(":" + port))
}
