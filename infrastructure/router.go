package infrastructure

import (
	"github.com/yama-koo/zipcloud-example/adapter/api"
	"github.com/yama-koo/zipcloud-example/adapter/controller"
	"github.com/yama-koo/zipcloud-example/adapter/database"
	"github.com/yama-koo/zipcloud-example/domain"
	"github.com/yama-koo/zipcloud-example/usecase"
	gin "gopkg.in/gin-gonic/gin.v1"
)

// Router struct
type Router struct {
	Router *gin.Engine
}

// Initialize func
func (r *Router) Initialize() {
	router := gin.Default()

	db := GormConnect()
	db.AutoMigrate(&domain.Zipcode{})

	controller := controller.ZipcodeController{
		ZipcodeInteractor: usecase.ZipcodeInteractor{
			ZipcodeRepository: &database.ZipcodeRepository{
				SQLHandler: &SQLHandler{
					DB: db,
				},
			},
		},
		ZipCloudInteractor: usecase.ZipCloudInteractor{
			ZipCloudRepository: &api.ZipCloudRepository{
				ZipCloudAPIHandler: &ZipCloudAPIHandler{},
			},
		},
	}

	router.GET("/zipcodes/search/:zipcode", func(c *gin.Context) { controller.Search(c) })
	// router.GET("/zipcodes", func(c *gin.Context) { controller.FindAll(c) })
	router.GET("/zipcodes/find/:id", func(c *gin.Context) { controller.FindByID(c) })
	// router.POST("/zipcodes", func(c *gin.Context) { controller.Create(c) })
	// router.DELETE("/zipcodes/:id", func(c *gin.Context) { controller.Delete(c) })

	r.Router = router
}
