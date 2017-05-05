package dish

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type AsDishes struct {
	Id              int              `json:"id"`
	DishName        string           `json:"dish_name"`
	DishPrice       float64          `json:"dish_price"`
	DishUnit        string           `json:"dish_unit"`
	DishDescription string           `json:"dish_description"`
	DishCreateTime  time.Time        `json:"dish_create_time"`
	DishModifyTime  time.Time        `json:"dish_modify_time"`
	DishCategoryId  int              `json:"dish_category_id"`
}

func init() {
	orm.RegisterModel(new(AsDishes))
}

func CreateDish(dishName string, dishPrice float64, dishUnit string, dishDescription string, dishCategoryId int) bool {
	o := orm.NewOrm()
	time := time.Now()
	dish := &AsDishes{DishName:dishName, DishPrice:dishPrice,
		DishUnit:dishUnit, DishDescription:dishDescription,
		DishCreateTime:time, DishModifyTime:time, DishCategoryId:dishCategoryId}
	_, err := o.Insert(dish)
	if err != nil {
		return false
	}
	return true
}

func GetDishes(search string, column string, dir string, pageSize int, pageNo int) ([]AsDishes, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(&AsDishes{})
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
	var categoryDishes []AsDishes
	_, err := qs.All(&categoryDishes)
	return categoryDishes, err
}

func GetDishesCount() int64 {
	o := orm.NewOrm()
	qs := o.QueryTable(&AsDishes{})
	unm, err := qs.Count()
	if err != nil {
		unm = 0
	}
	return unm
}

func ExistDishes(colName string, value string) bool {
	o := orm.NewOrm()
	qs := o.QueryTable(&AsDishes{})
	con := orm.NewCondition()
	con = con.And(colName, value)
	qs = qs.SetCond(con)
	return qs.Exist()
}

func DeleteDish(id int) {
	o := orm.NewOrm()
	qs := o.QueryTable(&AsDishes{})
	qs.Filter("id", id).Delete()
}

func UpdateDish(id int, categoryName string, dishSummary string, dishPrice float64, dish_unit string, dish_category_id int) (AsDishes, error) {
	var dish = AsDishes{Id:id, DishName:categoryName,
		DishDescription:dishSummary,
		DishPrice:dishPrice,
		DishUnit:dish_unit,
		DishCategoryId:dish_category_id,
		DishModifyTime:time.Now()}
	o := orm.NewOrm()
	_, err := o.Update(&dish, "dish_name", "dish_description", "dish_price", "dish_unit", "dish_category_id", "dish_modify_time")
	if err != nil {
		return AsDishes{}, err
	}
	o.Read(&dish)
	return dish, nil
}