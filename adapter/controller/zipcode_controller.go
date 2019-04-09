package controller

import (
	"strconv"

	"github.com/yama-koo/zipcloud-example/usecase"
)

// ZipcodeController struct
type ZipcodeController struct {
	ZipcodeInteractor  usecase.ZipcodeInteractor
	ZipCloudInteractor usecase.ZipCloudInteractor
}

// Search func
func (con *ZipcodeController) Search(c Context) {
	paramZipcode := c.Param("zipcode")

	zipcodes, err := con.ZipCloudInteractor.Search(paramZipcode)
	if err != nil {
		c.JSON(500, nil)
	}

	c.JSON(200, zipcodes)
}

// Create func
func (con *ZipcodeController) Create() {

}

// FindByID func
func (con *ZipcodeController) FindByID(c Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, nil)
	}

	zipcode, err := con.ZipcodeInteractor.FindByID(id)
	if err != nil {
		c.JSON(500, nil)
	}
	c.JSON(200, zipcode)
}
