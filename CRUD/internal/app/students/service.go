package students

import (
	"sample/internal/models/pagination"
	"sample/internal/models/student"
	repo "sample/internal/repository/studentsrepo"
)

// type Service interface {
// 	CreateContacts(data contact_models.ContactDetails) (interface{}, error)
// }
type Service interface {
	ListStudent(page *pagination.Paginatevalue) ([]student.StudentDetails, error)
	GetOneStudent( id uint) (interface{}, error)
	CreateStudent(data *student.StudentDetails) error
	UpdateStudent(query map[string]interface{}, data *student.StudentDetails) error
	DeleteStudent(id uint) error
}

type service struct {
	studentrepo repo.Contactstudent
}

func NewService() *service {

	studentrepoo := repo.NewRepository()
	return &service{studentrepoo}
}
func (s *service) ListStudent(page *pagination.Paginatevalue) ([]student.StudentDetails, error) {

	var result []student.StudentDetails
	result, err := s.studentrepo.Listcontactstudent(page)
	if err != nil {
		return nil, err

	}
	return result, nil

}
func (s *service) GetOneStudent(id uint) (interface{}, error) {
	result, err := s.studentrepo.GetOnecontactstudent(id)
	if err != nil {
		return result, err
	}
	return result, nil
}
func (s *service) CreateStudent(data *student.StudentDetails) error {
	err := s.studentrepo.Createcontactstudent(data)
	if err != nil {
		return err
	}
	return nil
}
func (s *service) UpdateStudent(query map[string]interface{}, data *student.StudentDetails) error {
	err := s.studentrepo.Updatecontactstudent(query, data)
	if err != nil {
		return err
	}
	return nil
}
func (s *service) DeleteStudent(id uint) error {
	err := s.studentrepo.Deletecontactstudent(id)
	if err != nil {
		return err
	}
	return nil
}
