package api

import (
	"github.com/yama-koo/zipcloud-example/domain"
)

// ZipCloudRepository struct
type ZipCloudRepository struct {
	ZipCloudAPIHandler
}

// Search func
func (repo *ZipCloudRepository) Search(paramZipcode string) ([]domain.Zipcode, error) {
	var zipcodes []domain.Zipcode

	zipcodes, err := repo.ZipCloudAPIHandler.Search(paramZipcode)
	if err != nil {
		return zipcodes, err
	}

	return zipcodes, nil
}
