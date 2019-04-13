package usecase

import (
	"github.com/yama-koo/zipcloud-example/domain"
)

// ZipCloudRepository interface
type ZipCloudRepository interface {
	Search(string) ([]domain.Zipcode, error)
}
