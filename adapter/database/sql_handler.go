package database

import "github.com/yama-koo/zipcloud-example/domain"

// SQLHandler interface
type SQLHandler interface {
	FindByID(id int) (zipcode domain.Zipcode, err error)
	Create(domain.Zipcode) error
}
