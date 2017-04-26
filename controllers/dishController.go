package controllers

import (
	"github.com/astaxie/beego"
	"os"
	"io/ioutil"
	"github.com/astaxie/beego/logs"
	"encoding/json"
)

type DishController struct {
	beego.Controller
}

type jsonData struct {
	Id    int `json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

type tatallData struct {
	Total int `json:"total"`
	Rows[] jsonData `json:"rows"`
}
func (this *DishController)Get() {
	bytes, _ := ReadAll("/home/yangcai/go/src/demo/static/data/data1.json")
	//data := string(bytes)
	var  tatall tatallData
	json.Unmarshal(bytes,&tatall)
	logs.Debug(tatall)
	this.Data["json"] = tatall
	this.ServeJSON()
}

func ReadAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}