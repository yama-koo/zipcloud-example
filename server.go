package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/yama-koo/zipcloud-example/infrastructure"
)

func main() {
	router := infrastructure.Router{}
	router.Initialize()
	router.Router.Run()
}
