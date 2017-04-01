package routers

import (
	"github.com/astaxie/beego"
	"demo/controllers"
)

func init() {
    beego.Router("/",&controllers.MainController{})
}
