package routers

import (
	"github.com/scrpgil/localfront/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/*.*", &controllers.MainController{}, "get:GetFile")
}
