package infrastructure_test

import (
	"fmt"
	"testing"

	"github.com/yama-koo/zipcloud-example/infrastructure"
)

func TestSearch(t *testing.T) {
	handler := infrastructure.ZipCloudAPIHandler{}

	zipcodes, err := handler.Search("3700536")
	if err != nil {
		t.Fatalf("%+v\n", err.Error())
	}

	if len(zipcodes) == 0 {
		t.Fatalf("zipcode not found")
	}

	fmt.Printf("%+v\n", zipcodes)
}
