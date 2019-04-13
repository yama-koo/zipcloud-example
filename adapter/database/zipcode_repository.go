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
