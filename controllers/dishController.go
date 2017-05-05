package controllers

import (
	"github.com/astaxie/beego"
	"demo/models/dish"
)

type tatallDataDish struct {
	RecordsTotal    int64 `json:"recordsTotal"`
	RecordsFiltered int64 `json:"recordsFiltered"`
	Draw            int `json:"draw"`
	Rows            [] dish.AsDishes `json:"data"`
}

type DishController struct {
	beego.Controller
}

func (this*DishController)Get() {
	draw, _ := this.GetInt("draw", 1)
	search := this.GetString("search", "")
	column := this.GetString("column", "")
	dir := this.GetString("dir", "")
	pageSize, _ := this.GetInt("pageSize", 10)
	pageNo, _ := this.GetInt("pageNo", 1)
	categoryDishes, err := dish.GetDishes(search, column, dir, pageSize, pageNo)
	var tatall tatallDataDish
	tatall.Draw = draw
	if err == nil {
		tatall.Rows = categoryDishes
	}
	count := dish.GetDishesCount()
	tatall.RecordsFiltered = count
	tatall.RecordsTotal = count
	this.Data["json"] = tatall
	this.ServeJSON()
}

func (this*DishController)Post() {

}

func (this*DishController)Put() {

}

func (this*DishController)Delete() {

}