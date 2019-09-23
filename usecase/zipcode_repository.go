package usecase

import (
	"github.com/yama-koo/zipcloud-example/domain"
)

// ZipcodeRepository interface
type ZipcodeRepository interface {
	FindByID(int) (domain.Zipcode, error)
	Create(domain.Zipcode) error
}
