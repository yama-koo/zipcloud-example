package usecase

import (
	"github.com/yama-koo/zipcloud-example/domain"
)

// ZipCloudInteractor struct
type ZipCloudInteractor struct {
	ZipCloudRepository
}

// Search func
func (interactor *ZipCloudInteractor) Search(paramZipcode string) ([]domain.Zipcode, error) {
	zipcodes := []domain.Zipcode{}

	zipcodes, err := interactor.ZipCloudRepository.Search(paramZipcode)
	if err != nil {
		return zipcodes, err
	}

	return zipcodes, nil
}
