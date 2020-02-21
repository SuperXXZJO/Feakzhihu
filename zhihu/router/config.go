package router

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"html/template"
	"io"
	"strconv"
	"zhihu/model"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if debug{
	t.templates =template.Must(template.ParseFiles("./views/login.html"))
	}
	return t.templates.ExecuteTemplate(w, name, data)
}

	var renderer = &TemplateRenderer{
	templates: template.Must(template.ParseFiles("./views/login.html")),
		}
//中间件
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
		tokenstring := c.FormValue("token")
		claims := model.UserToken{}
		token,err := jwt.ParseWithClaims(tokenstring,&claims, func(token *jwt.Token) (i interface{}, err error) {
			return []byte("123"),nil
		})
		if err == nil && token.Valid{
			return next(c)
		}else {
			return c.JSON(300,"验证失败！请先登录！")
		}

	}
}
//点击量
func Hits(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ID ,err := strconv.Atoi(c.Param("id"))
		if err != nil{
			return c.JSON(300,"未更新！")
		}
		err2 := model.QuestionHitsAdd(ID)
		if err2 != nil {
			return c.JSON(301,err.Error())
		}
		return next (c)
	}

}