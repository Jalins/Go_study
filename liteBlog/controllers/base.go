package controllers

import (
	"github.com/astaxie/beego"
)
// 约定：如果子controller 存在NestPrepare()方法，就实现了该接口，
type NestPreparer interface {
	NestPrepare()
}
type BaseController struct {
	beego.Controller
}
func (ctx *BaseController) Prepare() {

	ctx.Data["Path"] = ctx.Ctx.Request.RequestURI


	// 判断子类是否实现了NestPreparer接口，如果实现了就调用接口方法。
	if app, ok := ctx.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}