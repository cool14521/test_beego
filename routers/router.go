// @APIVersion 1.0.0
// @Title 财经报表
// @Description 财经报表通用版本

package routers

import (
	"finance_report/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		//通用报表
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ReportController{},
			),
		),
		//用户信息
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		//用户权限 role
		//用户认证 auth
	)
	beego.AddNamespace(ns)
}
