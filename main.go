package main

import (
	"beego-rest-api/database"
	_ "beego-rest-api/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	database.InitDB()
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
