package usecase

import (
	"github.com/yama-koo/zipcloud-example/domain"
)

// ZipcodeInteractor struct
type ZipcodeInteractor struct {
	ZipcodeRepository
}

// FindByID func
func (interactor *ZipcodeInteractor) FindByID(id int) (domain.Zipcode, error) {
	zipcode := domain.Zipcode{}
	zipcode, err := interactor.ZipcodeRepository.FindByID(id)
	if err != nil {
		return zipcode, err
	}

	return zipcode, nil
}
