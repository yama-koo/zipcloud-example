package database

import (
	"github.com/yama-koo/zipcloud-example/domain"
)

// ZipcodeRepository struct
type ZipcodeRepository struct {
	SQLHandler
}

// FindByID func
func (repo *ZipcodeRepository) FindByID(id int) (zipcode domain.Zipcode, err error) {
	return repo.SQLHandler.FindByID(id)
}

// Create func
func (repo *ZipcodeRepository) Create(zipcode domain.Zipcode) error {
	return repo.SQLHandler.Create(zipcode)
}
