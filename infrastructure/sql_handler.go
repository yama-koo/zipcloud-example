package infrastructure

import (
	"github.com/jinzhu/gorm"
	"github.com/yama-koo/zipcloud-example/domain"
)

// SQLHandler struct
type SQLHandler struct {
	DB *gorm.DB
}

// GormConnect func
func GormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "pass"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "zipcloud"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

// FindByID func
func (handler *SQLHandler) FindByID(id int) (zipcode domain.Zipcode, err error) {
	err = handler.DB.Where(&domain.Zipcode{ID: id}).First(&zipcode).Error
	if err != nil {
		return domain.Zipcode{}, err
	}

	return zipcode, nil
}

// Create func
func (handler *SQLHandler) Create(zipcode domain.Zipcode) error {
	err := handler.DB.Create(&zipcode).Error
	if err != nil {
		return err
	}
	return nil
}
