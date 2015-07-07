package routers

import (
	"github.com/astaxie/beego"
	"github.com/eefret/gastt-web/controllers"
)

func init() {
	beego.Router("/", &controllers.FrontendController{})
	beego.Router("/api", &controllers.ApiController{})
}
