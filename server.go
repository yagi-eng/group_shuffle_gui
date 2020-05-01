package main

import (
	"local.packages/handlers"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.Static("/static", "views")
	initRouting(e)
	e.Logger.Fatal(e.Start(":1323"))
}

func initRouting(e *echo.Echo) {
	e.File("/", "views/index.html")
	// http://localhost:1323/api/v1/simulateCombinations?allParticipants=18&participantsInEachGroup=6&repeatCnt=3&trials=10000
	e.GET("/api/v1/simulateCombinations", handlers.SimulateCombinations)
}
