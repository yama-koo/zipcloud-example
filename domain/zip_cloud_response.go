package domain

// ZipCloudResponse struct
type ZipCloudResponse struct {
	Message string    `json:"message"`
	Results []Zipcode `json:"results"`
	Status  int       `json:"status"`
}
