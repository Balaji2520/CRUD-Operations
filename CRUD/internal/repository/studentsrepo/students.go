package studentsrepo

import (
	"errors"
	"sample/db"
	"sample/internal/models/pagination"
	"sample/internal/models/student"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// type Contacts interface{
// 	CreateContacts(data contact_models.ContactDetails) (interface{}, error)
// }
type Contactstudent interface {
	Listcontactstudent(page *pagination.Paginatevalue) ([]student.StudentDetails, error)
	GetOnecontactstudent(id uint) (interface{}, error)
	Createcontactstudent(data *student.StudentDetails) error
	Updatecontactstudent(query map[string]interface{}, data *student.StudentDetails) error
	Deletecontactstudent(id uint) error
}

type students struct {
	db *gorm.DB
}

func NewRepository() *students {
	db := db.ReturnDatabaseObject()
	return &students{db}
}

// ------------------------contactsdent apps----------------------------------
func (r *students) Listcontactstudent(page *pagination.Paginatevalue) ([]student.StudentDetails, error) {
	var data []student.StudentDetails
	err := r.db.Model(&student.StudentDetails{}).Find(&data)

	if err.Error != nil {
		return nil, err.Error
	}

	return data, nil
}
func (r *students) GetOnecontactstudent(id uint) (interface{}, error) {
	var data student.StudentDetails
	err := r.db.Preload(clause.Associations+"."+clause.Associations).Where("id", id).First(&data)
	if err.RowsAffected == 0 {
		return nil, errors.New("record not found")
	}

	if err.Error != nil {
		return nil, err.Error
	}

	return data, nil
}
func (r *students) Createcontactstudent(data *student.StudentDetails) error {
	err := r.db.Model(&student.StudentDetails{}).Create(data).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *students) Updatecontactstudent(query map[string]interface{}, data *student.StudentDetails) error {
	res := r.db.Model(&student.StudentDetails{}).Where(query).Updates(data)
	if res.Error != nil {

		return res.Error

	}
	return nil
}
func (r *students) Deletecontactstudent(id uint) error {
	// zone := os.Getenv("DB_TZ")
	// loc, _ := time.LoadLocation(zone)
	// data := map[string]interface{}{
	// 	"deleted_by": query["user_id"].(int),
	// 	"deleted_at": time.Now().In(loc),
	// }
	// delete(query, "user_id")
	res := r.db.Model(&student.StudentDetails{}).Where("id", id).Updates(id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
