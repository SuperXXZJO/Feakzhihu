package router

import (
	"github.com/labstack/echo"
	"zhihu/control"
)

var debug = true

func Run(){
	zhihu :=echo.New()
	zhihu.Static("/static", "static")
	zhihu.HideBanner =true
	zhihu.Renderer = renderer
	zhihu.GET("/login.html",control.Loginview)//登录页面
	zhihu.POST("/login",control.Login) //登陆的post
	zhihu.POST("/signup",control.Signup) //注册的post
	api := zhihu.Group("/api",ServerHeader)
	Api(api)
	zhihu.Start(":8080")
}
