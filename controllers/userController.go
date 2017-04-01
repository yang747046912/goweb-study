package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type UserController struct {
	beego.Controller
}

func (this *UserController)Post() {
	username := this.GetString("username", "aaa")
	logs.Debug("用户登录 %s ", username)
	this.Data["json"] = "登录失败，用户名或密码不存在"
	this.ServeJSON()
}

func (this *UserController)Get() {
	this.TplName = "user/register.tpl"
}