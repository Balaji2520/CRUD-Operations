package students

import (
	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {

	g.GET("/students", h.ListStudent)
	g.GET("/student/:id", h.FindStudentById)
	g.POST("/student/create", h.CreateStudent)
	g.POST("/student/:id/update", h.UpdateStudent)
	g.DELETE("/student/:id/delete", h.DeleteStudent)

}
