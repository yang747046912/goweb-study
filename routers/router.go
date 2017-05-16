package routers

import (
	"github.com/astaxie/beego"
	"demo/controllers"
	"demo/controllers/app"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user/login", &controllers.UserController{})
	beego.Router("/user/password", &controllers.UserController{})
	beego.Router("/user/register", &controllers.SignController{})
	beego.Router("/main/view", &controllers.MainViewController{})
	beego.Router("/dish/category/data/?:id", &controllers.DishCategoryController{})
	beego.Router("/dish/dishes/data/?:id", &controllers.DishController{})
	beego.Router("/image/dishes/data/?:id", &controllers.ImageController{})

	beego.Router("/app/dish/categories",&app.AppDishCategoriesController{})
	beego.Router("/app/dish",&app.AppDishCOntroller{})
}
