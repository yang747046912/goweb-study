package dish

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type AsCategoryDishes struct {
	Id             int `json:"id"`
	CategoryName   string `json:"category_name"`
	DishCreateTime time.Time        `json:"dish_create_time"`
	DishSummary    string        `json:"dish_summary"`
	DishModifyTime time.Time        `json:"dish_modify_time"`
}

func init() {
	orm.RegisterModel(new(AsCategoryDishes))
}

func CreateDish(categoryName string, dishSummary string) bool {
	o := orm.NewOrm()
	var dishCategory AsCategoryDishes
	dishCategory.CategoryName = categoryName
	dishCategory.DishSummary = dishSummary
	time := time.Now()
	dishCategory.DishCreateTime = time
	dishCategory.DishModifyTime = time
	_, err := o.Insert(&dishCategory)
	if err != nil {
		return false
	}
	return true
}

func GetCount() int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(&AsCategoryDishes{})
	num, err := qs.Count()
	if err != nil {
		return 0
	}
	return num
}

func Exist(colName string, value string) bool {
	o := orm.NewOrm()
	qs := o.QueryTable(&AsCategoryDishes{})
	con := orm.NewCondition()
	con = con.And(colName, value)
	qs = qs.SetCond(con)
	return qs.Exist()
}