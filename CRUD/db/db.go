package db

import (
	"fmt"

	student_models "sample/internal/models/student"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func DBConnection() {
	var (
		DB_HOST    = "localhost"
		DB_USER    = "postgres"
		DB_PASS    = "Balaji@2001"
		DB_NAME    = "student"
		DB_PORT    = "5432"
		DB_SSLMODE = "disable"
		DB_TZ      = "Asia/Kolkata"
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", DB_HOST, DB_USER, DB_PASS, DB_NAME, DB_PORT, DB_SSLMODE, DB_TZ)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect sql database")
	}
	if err != nil {
		fmt.Print("Failed")
	}
	fmt.Println("DB Connected .........", db)

	db.AutoMigrate(student_models.ContactTables...)
	fmt.Println("========>>> Tables Migrated")

}

func ReturnDatabaseObject() *gorm.DB {
	return db
}
