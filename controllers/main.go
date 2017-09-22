package controllers

import (
 
	"github.com/astaxie/beego"
)

// Operations about Users
type MainController struct {
	beego.Controller
}
 
// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (c *MainController) Get() {
 	c.Data["Website"] = "test"
	c.Data["Email"] = "test@test.com"
	c.TplName = "index.tpl"
}
 