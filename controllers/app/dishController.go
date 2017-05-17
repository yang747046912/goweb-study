package app

import (
	"github.com/astaxie/beego"
	"demo/models/dish"
	"demo/models/images"
)

type AppDishCategoriesController struct {
	beego.Controller
}

type  AppCategoriesJsonOut struct {
	Id           int `json:"id"`
	CategoryName string `json:"category_name"`
}

func (this *AppDishCategoriesController)Get() {
	categories, err := dish.GetAllDishCateGories()
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code":-1, "message":"服务器异常", "data":nil}
	} else {
		var outDatas []AppCategoriesJsonOut
		for _, value := range categories {
			tmp := AppCategoriesJsonOut{value.Id, value.CategoryName}
			outDatas = append(outDatas, tmp)
		}
		this.Data["json"] = map[string]interface{}{"code":0, "message":nil, "data":outDatas}
	}
	this.ServeJSON()
}

type AppDishCOntroller struct {
	beego.Controller
}

type AppDishJsonOut struct {
	Id        int              `json:"id"`
	DishName  string           `json:"dish_name"`
	DishPrice float64          `json:"dish_price"`
	DishUnit  string           `json:"dish_unit"`
	Images    []images.AsImages `json:"images"`
}

func (this *AppDishCOntroller)Get() {
	dish_category_id, err1 := this.GetInt("dish_category_id")
	if err1 != nil {
		this.Data["json"] = map[string]interface{}{"code":-1, "message":"dish_category_id 参数异常", "data":nil}
		this.ServeJSON()
		return
	}
	categories, err := dish.GetDishesByDishCategoryId(dish_category_id)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code":-1, "message":"服务器异常", "data":nil}
		this.ServeJSON()
		return
	}
	var outDatas []AppDishJsonOut
	for _, value := range categories {
		tmp := AppDishJsonOut{value.Id, value.DishName, value.DishPrice, value.DishUnit, nil}
		imageIds := images.GetImageIDbyDishID(value.Id)
		if len(imageIds) != 0 {
			imageDishes := images.GetImageUrlByImageIDs(imageIds)
			tmp.Images = imageDishes
		}
		outDatas = append(outDatas, tmp)
	}
	this.Data["json"] = map[string]interface{}{"code":0, "message":nil, "data":outDatas}
	this.ServeJSON()
}

func (this *AppDishCOntroller )Post() {
	dish_category_id, _ := this.GetInt("dish_category_id")
	dish_name := this.GetString("dish_name", "")
	image_url := this.GetString("image_url", "")
	dish, _ := dish.CreateDish(dish_name, 100.00, "元/份", dish_name, dish_category_id)
	img, _ := images.Insert(image_url)
	images.InReferencesDishImages(img.Id, dish.Id)
	this.Data["json"] = map[string]interface{}{"code":0, "message":nil, "data":nil}
	this.ServeJSON()
}
