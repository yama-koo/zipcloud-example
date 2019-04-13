package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/yama-koo/zipcloud-example/infrastructure"
)

func main() {
	r := infrastructure.Router{}
	r.Initialize()
	r.Router.Run()
}
