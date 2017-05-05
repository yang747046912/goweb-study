package controllers

import (
	"github.com/astaxie/beego"
	"demo/models/dish"
	"strconv"
)

type DishCategoryController struct {
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

func (this *DishCategoryController)Get() {
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
	count := dish.GetDishCategoryCount()
	tatall.RecordsFiltered = count
	tatall.RecordsTotal = count
	this.Data["json"] = tatall
	this.ServeJSON()
}
func (this *DishCategoryController)Post() {
	var result reslutData
	category_name := this.GetString("category_name", "")
	dish_summary := this.GetString("dish_summary", "")
	if len(category_name) < 3 {
		errField := errorsField{"category_name", "菜品分类名称必须长度需大于2"}
		result.FieldErrors = append(result.FieldErrors, errField)
	}
	if len(dish_summary) < 3 {
		errField := errorsField{"dish_summary", "菜品分类简介必须长度需大于2"}
		result.FieldErrors = append(result.FieldErrors, errField)
	}
	if len(result.FieldErrors) != 0 {
		this.Data["json"] = result
		this.ServeJSON()
		return
	}

	dishExit := dish.ExistDishCategory("category_name", category_name)
	if dishExit {
		errField := errorsField{"category_name", "菜品名称分类已经存在"}
		result.FieldErrors = append(result.FieldErrors, errField)
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	dishCategory, success := dish.CreateDishCategory(category_name, dish_summary)
	if !success {
		errField := errorsField{"category_name", "系统错误"}
		result.FieldErrors = append(result.FieldErrors, errField)
	} else {
		result.Data = append(result.Data, dishCategory)
	}
	this.Data["json"] = result
	this.ServeJSON()
}
func (this *DishCategoryController)Delete() {
	id := this.Ctx.Input.Param(":id")
	iid, _ := strconv.Atoi(id)
	dish.DeleteDishCategory(iid)
	type result struct {
		Data [] string `json:"data"`
	}
	this.Data["json"] = &result{}
	this.ServeJSON()
}
func (this *DishCategoryController)Put() {
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
	exit := dish.ExistDishCategory("category_name", category_name)
	if exit {
		errField := errorsField{"category_name", "菜品名称已经存在"}
		result.FieldErrors = append(result.FieldErrors, errField)
		this.Data["json"] = result
		this.ServeJSON()
		return
	}

	id := this.Ctx.Input.Param(":id")
	iid, _ := strconv.Atoi(id)
	category, err := dish.UpdateDishCategory(iid, category_name, dish_summary)
	if err != nil {
		errField := errorsField{"category_name", "系统错误"}
		result.FieldErrors = append(result.FieldErrors, errField)
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	result.Data = append(result.Data, category)
	this.Data["json"] = result
	this.ServeJSON()
}