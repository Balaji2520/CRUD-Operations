package route

import (
	"fmt"
	student "sample/internal/app/students"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func Init(g *echo.Group) {
	var (
		APP     = os.Getenv("APP")
		VERSION = os.Getenv("VERSION")
	)
	g.GET("/", func(c echo.Context) error {
		message := fmt.Sprintf("Welcome to %s version %s", APP, VERSION)
		return c.String(http.StatusOK, message)
	})

	student.NewHandler().Route(g.Group("api/v1/student"))

}
