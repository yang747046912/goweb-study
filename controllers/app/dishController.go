package app

import (
	"github.com/astaxie/beego"
	"demo/models/dish"
)

type AppDishController struct {
	beego.Controller
}

type  jsonOut struct {
	Id           int `json:"id"`
	CategoryName string `json:"category_name"`
}

func (this *AppDishController)Get() {
	categories, err := dish.GetAllDishCateGories()
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code":-1, "message":"服务器异常", "data":nil}
	} else {
		var outDatas []jsonOut
		for _, value := range categories {
			tmp := jsonOut{value.Id, value.CategoryName}
			outDatas = append(outDatas, tmp)
		}
		this.Data["json"] = map[string]interface{}{"code":0, "message":nil, "data":outDatas}
	}
	this.ServeJSON()
}
