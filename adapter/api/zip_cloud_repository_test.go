package api_test

import (
	"fmt"
	"testing"

	"github.com/yama-koo/zipcloud-example/adapter/api"
	"github.com/yama-koo/zipcloud-example/infrastructure"
)

func TestSearch(t *testing.T) {
	repo := api.ZipCloudRepository{
		ZipCloudAPIHandler: &infrastructure.ZipCloudAPIHandler{},
	}

	zipcodes, err := repo.Search("3700536")
	if err != nil {
		t.Fatalf("%+v\n", err.Error())
	}

	fmt.Printf("%+v\n", zipcodes)
}
