package main

import (
	"log"

	"backend/internal/choice"
	"backend/internal/config"
	"backend/internal/db"
	"backend/internal/question"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.LoadEnv()
	pg := db.Connect(cfg.DatabaseURL)

	e := echo.New()

	e.Use(middleware.CORS())
	question.RegisterRoutes(e, pg)
	choice.RegisterRoutes(e, pg)

	log.Println("✅ Server started at :8080")
	e.Logger.Fatal(e.Start(":8080"))
}
