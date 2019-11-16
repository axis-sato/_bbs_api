package main

import (
	"github.com/c8112002/bbs_api/interfaces/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	questionController := controllers.NewQuestionController()

	// Routes
	e.GET("/questions", questionController.ShowList)

	// Start server
	e.Logger.Fatal(e.Start(":1234"))
}
