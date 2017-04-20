package controllers

import (
	"github.com/astaxie/beego"
	"demo/models/user"
	"github.com/astaxie/beego/logs"
)

type UserController struct {
	beego.Controller
}

func (this *UserController)Post() {

	userName := this.GetString("username", "")
	password := this.GetString("password", "")
	if len(userName) < 6 {
		this.Data["json"] = "登录失败，请输入六位用户名"
		this.ServeJSON()
		return
	}
	if len(password) < 6 {
		this.Data["json"] = "登录失败，请输入密码"
		this.ServeJSON()
		return
	}
	err, _ := user.LoginUser(userName, password)
	if err != nil {
		logs.Debug(err)
		this.Data["json"] = "登录失败，用户名或密码错误"
		this.ServeJSON()
		return
	}
	this.Data["json"] = "登录成功"
	this.ServeJSON()

}

func (this *UserController)Get() {
	this.TplName = "user/register.tpl"
}