package main

import (
	database "sample/db"
	"github.com/labstack/echo/v4"
)

func main() {
	database.DBConnection()

	e := echo.New()



	e.Start(":8000")

}
