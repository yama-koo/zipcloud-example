package api

import (
	"github.com/yama-koo/zipcloud-example/domain"
)

// ZipCloudAPIHandler interface
type ZipCloudAPIHandler interface {
	Search(string) ([]domain.Zipcode, error)
}
