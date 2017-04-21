package user

import "github.com/astaxie/beego/orm"

type UserInfo struct {
	Id       int
	UserName string
	Password string
	Email    string
}

func init() {
	orm.RegisterModel(new(UserInfo))
}

func LoginUser(userName string, password string) (error, UserInfo) {
	o := orm.NewOrm()
	qs := o.QueryTable(&UserInfo{})
	con := orm.NewCondition()
	con = con.And("user_name", userName)
	con = con.And("password", password)
	qs = qs.SetCond(con)
	var user UserInfo
	err := qs.Limit(1).One(&user)
	return err, user
}

func SignUser(userName string, password string, email string) string {
	uExist := exist("user_name", userName)
	if uExist {
		return "用户名已经被其他人使用了"
	}
	uEmail := exist("email", email)
	if uEmail {
		return "邮箱已经被其他人使用了"
	}
	o := orm.NewOrm()
	user := &UserInfo{UserName:userName, Password:password, Email:email}
	o.Insert(user)
	return "注册成功"
}

func exist(colName string, value string) bool {
	o := orm.NewOrm()
	qs := o.QueryTable(&UserInfo{})
	con := orm.NewCondition()
	con = con.And(colName, value)
	qs = qs.SetCond(con)
	return qs.Exist()
}

func FindPasswordByEmail(email string) (error, string) {
	o := orm.NewOrm()
	qs := o.QueryTable(&UserInfo{})
	con := orm.NewCondition()
	con = con.And("email", email)
	qs = qs.SetCond(con)
	var user UserInfo
	err := qs.Limit(1).One(&user)
	return err, user.Password
}