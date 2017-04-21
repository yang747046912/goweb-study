package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"fmt"
)

type MainViewController struct {
	beego.Controller
}

func (this *MainViewController)Get() {
	url := this.Ctx.Input.URI()
	logs.Debug(url)
	page := this.Input().Get("page")
	logs.Debug(page)
	if page == "" {
		this.TplName = "main.html"
		return
	}
	tpl := fmt.Sprintf("content/%s.html", page)
	logs.Debug(tpl)
	this.TplName = tpl
}