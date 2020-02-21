package control

import (
	"github.com/labstack/echo"
	"strconv"
	"zhihu/model"
)

func Loginview (c echo.Context)error{
	return c.Render(200,"login.html",nil)
}

//个人信息页
func Peopleview (c echo.Context)error{
	userid,_ := strconv.Atoi(c.Param("userid"))
	mod,err := model.FindByUserid(userid)
	if err != nil {
		return c.JSON(400,err.Error())
	}
	return c.JSON(200,mod)
}


//写文章详情页
func ArticleView (c echo.Context)error{
	return c.JSON(200,"")
}

//推荐

func QuestionView (c echo.Context)error{
	mods,err :=model.QuestionSelectByutime()
	if err !=nil {
		return c.JSON(400,err.Error())
	}
	return c.JSON(200,mods)

}
