package students

import (
	student "sample/internal/models/student"
	"strconv"

	"sample/internal/models/pagination"
	"github.com/labstack/echo/v4"
)

type handler struct {
	service Service
}

func NewHandler() *handler {
	service := NewService()
	return &handler{service}
}

func (h *handler) CreateStudent(c echo.Context) (err error) {
	var data student.StudentDetails
	err = h.service.CreateStudent(&data)
	if err != nil {
		return err
	}
	return nil
}

func (h *handler) UpdateStudent(c echo.Context) (err error) {

	var id = c.Param("id")
	var query = make(map[string]interface{}, 0)
	ID, _ := strconv.Atoi(id)
	query["id"] = ID
	var data student.StudentDetails
	err = h.service.UpdateStudent(query, &data)
	if err != nil {
		return err
	}
	return nil
}

func (h *handler) ListStudent(c echo.Context) (err error) {

	p := new(pagination.Paginatevalue)
	c.Bind(p)
	_,err = h.service.ListStudent(p)
	if err != nil {
		return err
	}
	return nil
}

func (h *handler) FindStudentById(c echo.Context) (err error) {
	ID := c.Param("id")
	id, _ := strconv.Atoi(ID)
	_, err = h.service.GetOneStudent(uint(id))
	if err != nil {
		return err
	}
	return nil
}

func (h *handler) DeleteStudent(c echo.Context) (err error) {
	ID := c.Param("id")
	id, _ := strconv.Atoi(ID)
	err = h.service.DeleteStudent(uint(id))
	if err != nil {
		return nil
	}
	return err
}
