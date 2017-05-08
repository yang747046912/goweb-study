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