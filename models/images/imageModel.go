package images

import "github.com/astaxie/beego/orm"

type AsImages struct {
	Id       int `json:"id"`
	ImageUrl string `json:"image_url"`
}

func init() {
	orm.RegisterModel(new(AsImages))
	orm.RegisterModel(new(AsImagesDishes))
}

type AsImagesDishes struct {
	Id      int
	ImageId int
	DishId  int
}

func Insert(imageUrl string) (AsImages, error) {
	o := orm.NewOrm()
	image := AsImages{ImageUrl:imageUrl}
	unm, err := o.Insert(&image)
	if err != nil {
		return AsImages{}, err
	}
	image.Id = int(unm)
	return image, nil
}

func InReferencesDishImages(imageId int, dishId int) error {
	o := orm.NewOrm()
	image := AsImagesDishes{ImageId:imageId, DishId:dishId}
	_, inErrTmp := o.Insert(&image)
	return inErrTmp
}

func GetImageIDbyDishID(dishId int) [] int {
	o := orm.NewOrm()
	qs := o.QueryTable(&AsImagesDishes{})
	com := orm.NewCondition()
	com = com.And("dish_id", dishId)
	qs = qs.SetCond(com)
	var image []*AsImagesDishes
	//	qs.All(&users)
	qs.All(&image)
	var imageIds []int
	for _,value:=range image {
		imageIds = append(imageIds,value.ImageId)
	}
	return imageIds
}

func GetImageUrlByImageIDs(imageIDs []int) []AsImages {
	o := orm.NewOrm()
	qs := o.QueryTable(&AsImages{})
	var image []AsImages
	qs.Filter("id__in" ,imageIDs).All(&image)
	return image
}