package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Storage DBStorage

type DBStorage struct {
	DB     *gorm.DB
	Opened bool
}

func (s *DBStorage) Connect() (*gorm.DB, error) {
	var err error
	// Get database details from environment variables
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	DBName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	s.DB, err = gorm.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			user, password, host, DBName,
		),
	)
	return s.DB, err
}
func (s *DBStorage) Close() error {
	s.Opened = false
	return s.DB.Close()
}
