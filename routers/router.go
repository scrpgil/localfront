package routers

import (
	"localfront/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/*.*", &controllers.MainController{}, "get:GetFile")
}
