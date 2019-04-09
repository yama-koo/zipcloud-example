package infrastructure

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/yama-koo/zipcloud-example/domain"
)

// ZipCloudAPIHandler struct
type ZipCloudAPIHandler struct {
}

// Search func
func (handler *ZipCloudAPIHandler) Search(paramZipcode string) ([]domain.Zipcode, error) {
	var zipCloudResponse domain.ZipCloudResponse

	values := url.Values{}
	values.Add("zipcode", paramZipcode)

	res, err := http.Get("http://zipcloud.ibsnet.co.jp/api/search" + "?" + values.Encode())
	if err != nil {
		return zipCloudResponse.Results, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return zipCloudResponse.Results, err
	}

	json.Unmarshal(body, &zipCloudResponse)

	return zipCloudResponse.Results, nil
}
