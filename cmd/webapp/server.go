package main

import (
	"os"

	"local.packages/api"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.Static("/static", "website")
	initRouting(e)
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

func initRouting(e *echo.Echo) {
	e.File("/", "website/index.html")
	// http://localhost:1323/api/v1/simulateCombinations?allParticipants=18&participantsInEachGroup=6&repeatCnt=3&trials=10000
	e.GET("/api/v1/simulateCombinations", api.SimulateCombinations)
}
