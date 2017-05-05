package controllers

import (
	"github.com/astaxie/beego"
	"demo/models/dish"
	"strconv"
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
	var result reslutData
	dish_name := this.GetString("dish_name", "")
	dish_description := this.GetString("dish_description", "")
	dish_price, price_err := this.GetFloat("dish_price", 0)
	dish_unit := this.GetString("dish_unit", "")
	dish_category_id, id_err := this.GetInt("dish_category_id", 0)
	if len(dish_name) < 3 {
		errField := errorsField{"dish_name", "菜品名称必须长度需大于2"}
		result.FieldErrors = append(result.FieldErrors, errField)
	}
	if len(dish_description) < 3 {
		errField := errorsField{"dish_description", "菜品简介必须长度需大于2"}
		result.FieldErrors = append(result.FieldErrors, errField)
	}
	if price_err != nil {
		errField := errorsField{"dish_price", "菜品单价必须为数字,不能包含其他字符(例如:12.45)"}
		result.FieldErrors = append(result.FieldErrors, errField)
	} else if dish_price == 0 {
		errField := errorsField{"dish_price", "菜品单价不能为0"}
		result.FieldErrors = append(result.FieldErrors, errField)
	}
	if len(dish_unit) < 1 {
		errField := errorsField{"dish_unit", "菜品单位长度必须大于1"}
		result.FieldErrors = append(result.FieldErrors, errField)
	}
	if id_err != nil {
		errField := errorsField{"dish_category_id", "请选择正确的菜品分类"}
		result.FieldErrors = append(result.FieldErrors, errField)
	} else if dish_category_id == 0 {
		errField := errorsField{"dish_category_id", "请选择正确的菜品分类"}
		result.FieldErrors = append(result.FieldErrors, errField)
	}
	if len(result.FieldErrors) != 0 {
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	dishExit := dish.ExistDishes("dish_name", dish_name)
	if dishExit {
		errField := errorsField{"dish_name", "菜品名称已经存在"}
		result.FieldErrors = append(result.FieldErrors, errField)
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	success := dish.CreateDish(dish_name, dish_price, dish_unit, dish_description, dish_category_id)
	if !success {
		errField := errorsField{"dish_name", "系统错误"}
		result.FieldErrors = append(result.FieldErrors, errField)
	}
	this.Data["json"] = result
	this.ServeJSON()
}

type reslutDataDish struct {
	FieldErrors []errorsField `json:"fieldErrors"`
	Data        []dish.AsDishes `json:"data"`
}


func (this*DishController)Put() {
	var result reslutDataDish
	dish_name := this.GetString("dish_name", "")
	dish_description := this.GetString("dish_description", "")
	dish_price, price_err := this.GetFloat("dish_price", 0)
	dish_unit := this.GetString("dish_unit", "")
	dish_category_id, id_err := this.GetInt("dish_category_id", 0)
	if len(dish_name) < 3 {
		errField := errorsField{"dish_name", "菜品名称必须长度需大于2"}
		result.FieldErrors = append(result.FieldErrors, errField)
	}
	if len(dish_description) < 3 {
		errField := errorsField{"dish_description", "菜品简介必须长度需大于2"}
		result.FieldErrors = append(result.FieldErrors, errField)
	}
	if price_err != nil {
		errField := errorsField{"dish_price", "菜品单价必须为数字,不能包含其他字符(例如:12.45)"}
		result.FieldErrors = append(result.FieldErrors, errField)
	} else if dish_price == 0 {
		errField := errorsField{"dish_price", "菜品单价不能为0"}
		result.FieldErrors = append(result.FieldErrors, errField)
	}
	if len(dish_unit) < 1 {
		errField := errorsField{"dish_unit", "菜品单位长度必须大于1"}
		result.FieldErrors = append(result.FieldErrors, errField)
	}
	if id_err != nil {
		errField := errorsField{"dish_category_id", "请选择正确的菜品分类"}
		result.FieldErrors = append(result.FieldErrors, errField)
	} else if dish_category_id == 0 {
		errField := errorsField{"dish_category_id", "请选择正确的菜品分类"}
		result.FieldErrors = append(result.FieldErrors, errField)
	}
	if len(result.FieldErrors) != 0 {
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	dishExit := dish.ExistDishes("dish_name", dish_name)
	if dishExit {
		errField := errorsField{"dish_name", "菜品名称已经存在"}
		result.FieldErrors = append(result.FieldErrors, errField)
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	id := this.Ctx.Input.Param(":id")
	iid, _ := strconv.Atoi(id)
	category,err := dish.UpdateDish(iid, dish_name, dish_description,dish_price,dish_unit,dish_category_id)
	if err!=nil {
		errField := errorsField{"dish_name", "系统错误"}
		result.FieldErrors = append(result.FieldErrors, errField)
		this.Data["json"] = result
		this.ServeJSON()
		return
	}
	result.Data = append(result.Data, category)
	this.Data["json"] = result
	this.ServeJSON()
}

func (this*DishController)Delete() {
	id := this.Ctx.Input.Param(":id")
	iid, _ := strconv.Atoi(id)
	dish.DeleteDish(iid)
	type result struct {
		Data [] string `json:"data"`
	}
	this.Data["json"] = &result{}
	this.ServeJSON()
}