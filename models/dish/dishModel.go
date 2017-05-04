package dish

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type AsDishes struct {
	Id              int              `json:"id"`
	DishName        string           `json:"dish_name"`
	DishPrice       float32          `json:"dish_price"`
	DishUnit        string           `json:"dish_unit"`
	DishDescription string           `json:"dish_description"`
	DishCreateTime  time.Time        `json:"dish_create_time"`
	DishModifyTime  time.Time        `json:"dish_modify_time"`
	DishCategoryId  int              `json:"dish_category_id"`
}

func init() {
	orm.RegisterModel(new(AsDishes))
}