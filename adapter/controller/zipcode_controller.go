package controller

import (
	"strconv"

	"github.com/yama-koo/zipcloud-example/usecase"
)

// ZipcodeController struct
type ZipcodeController struct {
	// TODO: controllerレベルで別れる？
	ZipcodeInteractor  usecase.ZipcodeInteractor
	ZipCloudInteractor usecase.ZipCloudInteractor
}

// Search func
func (con *ZipcodeController) Search(c Context) {
	paramZipcode := c.Param("zipcode")

	zipcodes, err := con.ZipCloudInteractor.Search(paramZipcode)
	if err != nil {
		c.JSON(500, map[string]interface{}{"message": err.Error()})
		return
	}

	isNotFound := len(zipcodes) == 1 && zipcodes[0].Prefecture == "" && zipcodes[0].PrefectureCode == ""
	if isNotFound {
		c.JSON(200, nil)
		return
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
		c.JSON(500, map[string]interface{}{"message": err.Error()})
		return
	}

	zipcode, err := con.ZipcodeInteractor.FindByID(id)
	if err != nil {
		c.JSON(500, map[string]interface{}{"message": err.Error()})
		return
	}

	c.JSON(200, zipcode)
}
