package routers

import (
	"Go_study/liteBlog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
	beego.Include(&controllers.IndexController{})
	beego.ErrorController(&controllers.ErrorController{})
}
