package controllers

import (
	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	if beego.BConfig.RunMode == "dev" {
		beego.Warn(u.Ctx.Input.Header("user-agent"))
	}
	u.Ctx.WriteString(u.Ctx.Input.Header("user-agent"))
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	if beego.BConfig.RunMode == "dev" {
		beego.Warn(u.Ctx.Input.Header("user-agent"))
	}
	u.Ctx.WriteString(u.Ctx.Input.Header("user-agent"))
}
