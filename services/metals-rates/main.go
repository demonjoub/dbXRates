package main

import (
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	// get all
	e.GET("/latest", getAllHandler)
	// get
	e.GET("/latest/:date", getHandler)
	// update
	e.PUT("/latest/:date", updateHandler)
	// delate
	e.DELETE("/delete/:date", deteteHandle)

	e.Logger.Fatal(e.Start(":1323"))
}
