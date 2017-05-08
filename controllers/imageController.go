package controllers

import (
	"github.com/astaxie/beego"
	"time"
	"strconv"
	"os"
	"opms/utils"
	"demo/models/images"
	"strings"
)

type ImageController struct {
	beego.Controller
}

func (this *ImageController)Post() {
	imageFile, _, err := this.GetFile("upload")
	if err == nil {
		defer imageFile.Close()
		now := time.Now()
		dir := "./static/img/" + strconv.Itoa(now.Year()) + "-" + strconv.Itoa(int(now.Month())) + "/" + strconv.Itoa(now.Day())
		err1 := os.MkdirAll(dir, 0755)
		if err1 != nil {
			this.Data["json"] = map[string]interface{}{"code": 1, "message": "目录权限不够"}
			this.ServeJSON()
			return
		}
		fileName := dir + "/" + utils.GetGuid()
		this.SaveToFile("upload", fileName)
		filepath := strings.Replace(fileName, ".", "", 1)
		image, _ := images.Insert(filepath)
		id := map[string]interface{}{"id":image}

		this.Data["json"] = map[string]interface{}{"upload":id}

		this.ServeJSON()
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": err}
		this.ServeJSON()
	}
}

