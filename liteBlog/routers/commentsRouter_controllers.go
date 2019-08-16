package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["Go_study/liteBlog/controllers:IndexController"] = append(beego.GlobalControllerRouter["Go_study/liteBlog/controllers:IndexController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Go_study/liteBlog/controllers:IndexController"] = append(beego.GlobalControllerRouter["Go_study/liteBlog/controllers:IndexController"],
		beego.ControllerComments{
			Method: "GetAbout",
			Router: `/about`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["Go_study/liteBlog/controllers:IndexController"] = append(beego.GlobalControllerRouter["Go_study/liteBlog/controllers:IndexController"],
		beego.ControllerComments{
			Method: "GetMessage",
			Router: `/message`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
