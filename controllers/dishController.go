package controllers

import (
	"github.com/astaxie/beego"
	"os"
	"io/ioutil"
	"github.com/astaxie/beego/logs"
	"net/url"
	"demo/models/dish"
)

type DishController struct {
	beego.Controller
}

type tatallData struct {
	RecordsTotal    int64 `json:"recordsTotal"`
	RecordsFiltered int64 `json:"recordsFiltered"`
	Draw            int `json:"draw"`
	Rows            [] dish.AsCategoryDishes `json:"data"`
}

type errorsField struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

type reslutData struct {
	FieldErrors []errorsField `json:"fieldErrors"`
	Data        []dish.AsCategoryDishes `json:"data"`
}

func (this *DishController)Get() {
	draw, _ := this.GetInt("draw", 1)
	search := this.GetString("search", "")
	column := this.GetString("column", "")
	dir := this.GetString("dir", "")
	pageSize, _ := this.GetInt("pageSize", 10)
	pageNo, _ := this.GetInt("pageNo", 1)
	categoryDishes, err := dish.GetDishCategories(search, column, dir, pageSize, pageNo)
	var tatall tatallData
	tatall.Draw = draw
	if err == nil {
		tatall.Rows = categoryDishes
	}
	count := dish.GetCount()
	tatall.RecordsFiltered = count
	tatall.RecordsTotal = count
	this.Data["json"] = tatall
	this.ServeJSON()
}
func (this *DishController)Post() {
	var result reslutData
	category_name := this.GetString("category_name", "")
	dish_summary := this.GetString("dish_summary", "")
	if len(category_name) < 3 {
		errField := errorsField{"category_name", "菜品名称必须长度需大于2"}
		result.FieldErrors = append(result.FieldErrors, errField)
	}
	if len(dish_summary) < 3 {
		errField := errorsField{"dish_summary", "菜品简介必须长度需大于2"}
		result.FieldErrors = append(result.FieldErrors, errField)
	}
	if len(result.FieldErrors) != 0 {
		this.Data["json"] = result
		this.ServeJSON()
		return
	}

	dishExit := dish.Exist("category_name", category_name)
	if dishExit {
		errField := errorsField{"category_name", "菜品名称已经存在"}
		result.FieldErrors = append(result.FieldErrors, errField)
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	success := dish.CreateDish(category_name, dish_summary)
	if !success {
		errField := errorsField{"category_name", "系统错误"}
		result.FieldErrors = append(result.FieldErrors, errField)
	}
	this.Data["json"] = result
	this.ServeJSON()
}

func (this *DishController)Delete() {
	logs.Debug("delete")
	uri := this.Ctx.Input.URI()
	logs.Debug(url.PathUnescape(uri))
	type result struct {
		Data [] string `json:"data"`
	}

	this.Data["json"] = &result{}
	this.ServeJSON()
}

func ReadAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}