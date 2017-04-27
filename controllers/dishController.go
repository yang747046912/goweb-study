package controllers

import (
	"github.com/astaxie/beego"
	"os"
	"io/ioutil"
	"github.com/astaxie/beego/logs"
	"encoding/json"
	"net/url"
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
	RecordsTotal int `json:"recordsTotal"`
	RecordsFiltered int `json:"recordsFiltered"`
	Draw int `json:"draw"`
	Rows[] jsonData `json:"data"`
}
func (this *DishController)Get() {
	uri := this.Ctx.Input.URI()
	logs.Debug(url.PathUnescape(uri))
	bytes, _ := ReadAll("/home/yangcai/go/src/demo/static/data/data1.json")
	//data := string(bytes)
	 draw,_:=this.GetInt("draw", 1)
	var  tatall tatallData
	json.Unmarshal(bytes,&tatall)
	tatall.Draw = draw
	logs.Debug(tatall)
	this.Data["json"] = tatall
	this.ServeJSON()
}

func ReadAll(filePth string) ([]byte, error) {
	f, err := os.Open(filePth)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}