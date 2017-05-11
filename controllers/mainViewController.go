package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"fmt"
	"demo/models/dish"
)

type MainViewController struct {
	beego.Controller
}

func (this *MainViewController)Get() {
	url := this.Ctx.Input.URI()
	logs.Debug(url)
	page := this.Input().Get("page")
	logs.Debug(page)
	if page == "" {
		this.TplName = "main.html"
		return
	}
	if page =="index"{
		page="dish-category-manager"
	}
	tpl := fmt.Sprintf("content/%s.html", page)
	logs.Debug(tpl)
	if page == "dish-manager" {
		categoryDishes, err := dish.GetAllDishCateGories()
		if err == nil {
			type outData struct {
				ID   int`json:"id"`
				Name string`json:"name"`
			}
			var outDatas [] outData
			for _, value := range categoryDishes {
				var outTmp outData
				outTmp.ID = value.Id
				outTmp.Name = value.CategoryName
				outDatas = append(outDatas, outTmp)
			}
			this.Data["options"] = &outDatas
		}
	}
	this.TplName = tpl
}