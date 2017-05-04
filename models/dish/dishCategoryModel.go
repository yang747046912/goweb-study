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

func GetDishCategories(search string, column string, dir string, pageSize int, pageNo int) ([]AsCategoryDishes, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(&AsCategoryDishes{})
	if len(search) != 0 {
		qs = qs.Filter("category_name__contains", search)
	}
	if len(column) != 0&& len(dir) != 0 {
		if dir == "asc" {
			qs = qs.OrderBy(column)
		} else if dir == "desc" {
			qs = qs.OrderBy("-" + column)
		}
	}
	offset := (pageNo - 1) * pageSize
	qs = qs.Limit(pageSize, offset)
	var categoryDishes []AsCategoryDishes
	_, err := qs.All(&categoryDishes)
	return categoryDishes, err
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

func Delete(id int) {
	o := orm.NewOrm()
	qs := o.QueryTable(&AsCategoryDishes{})
	qs.Filter("id", id).Delete()
}

func Update(id int, categoryName string, dishSummary string) AsCategoryDishes {
	var dish = AsCategoryDishes{Id:id, CategoryName:categoryName,
		DishSummary:dishSummary, DishModifyTime:time.Now()}
	o := orm.NewOrm()
	o.Update(&dish, "category_name", "dish_summary", "dish_modify_time")
	o.Read(&dish)
	return dish
}