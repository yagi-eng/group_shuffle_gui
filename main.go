package main

import (
	"github.com/labstack/echo"
	"github.com/yagi_eng/group_shuffle_gui/handler"
)

func main() {
	e := echo.New()
	e.Static("/static", "web")
	initRouting(e)
	e.Logger.Fatal(e.Start(":8080"))
}

func initRouting(e *echo.Echo) {
	e.File("/", "web/index.html")
	// http://localhost:1323/api/v1/simulateCombinations?allParticipants=18&participantsInEachGroup=6&repeatCnt=3&trials=10000
	e.GET("/api/v1/simulateCombinations", handler.SimulateCombinations)
}
