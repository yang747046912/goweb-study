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
	email := this.GetString("email")
	if "" == email {
		this.Data["json"] = "邮箱不能为空"
		this.ServeJSON()
		return
	}
	err, msg := user.FindPasswordByEmail(email)
	if err != nil {
		this.Data["json"] = "修改密码失败"
		this.ServeJSON()
		return
	}
	this.Data["json"] = msg
	this.ServeJSON()
	return
}

type SignController struct {
	beego.Controller
}

func (this *SignController)Post() {
	email := this.GetString("email")
	if email == "" {
		this.Data["json"] = "请正确输入邮箱"
		this.ServeJSON()
		return
	}
	username := this.GetString("username")
	if username == "" {
		this.Data["json"] = "请至少正确输入六位用户名"
		this.ServeJSON()
		return
	}
	password := this.GetString("password")
	if password == "" {
		this.Data["json"] = "请至少正确输入六位密码"
		this.ServeJSON()
		return
	}
	msg := user.SignUser(username, password, email)
	this.Data["json"] = msg
	this.ServeJSON()
}