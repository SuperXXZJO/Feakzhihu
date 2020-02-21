package control

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"strconv"
	"time"

	//"strconv"
	"zhihu/model"
)


//login 登录
func Login (c echo.Context)error{
	phone,_:= strconv.Atoi(c.FormValue("phone"))
	password :=c.FormValue("password")
	mod,err :=model.Login(phone)
	if err != nil{
		return c.JSON(300,"请输入正确的手机号！")
	}
	if mod.Password != password{
		return c.JSON(300,"密码错误")
	}
	//生成token
	claims := model.UserToken{
		Userid: mod.Userid,
		Username: mod.Username,
		StandardClaims: jwt.StandardClaims{ExpiresAt:time.Now().Add(2 * time.Hour).Unix()},
	}
	token :=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	ss,err := token.SignedString([]byte("123"))
	return c.JSON(200,ss)
}
//signuo 注册
func Signup (c echo.Context)error{
	inf := model.User{}
	err := c.Bind(&inf)
	if err != nil{
		return c.JSON(300,"输入数据有误")
	}
	if len(strconv.Itoa(inf.Phone)) !=11 {
		return c.JSON(300,"请输入正确的手机号码！")
	}
	if  inf.Password == ""{
		return c.JSON(300,"密码不能为空！")
	}

	err = model.Signup(&inf)
	if err != nil {
		return c.JSON(301,err.Error())
	}
	return c.JSON(200,"注册成功！")
}

//提问
func QuestionAdd (c echo.Context)error{
	inf := model.Question{}
	err := c.Bind(&inf)
	if err != nil{
		return c.JSON(300,"输入数据有误")
	}
	if inf.Question == ""{
		return c.JSON(300,"问题不能为空！")
	}

	err = model.QuestionAdd(&inf)
	if err != nil {
		return c.JSON(301,err.Error())
	}
	return c.JSON(200,"提问成功！")
}
//回答
func ResponseAdd (c echo.Context)error{
	inf := model.Response{}
	err := c.Bind(&inf)
	if err != nil{
		return c.JSON(300,"输入数据有误")
	}
	if inf.Question_id <=0 {
		return c.JSON(300,"请传入正确的问题编号")
	}
	if inf.Content == ""{
		return c.JSON(300,"回答不能为空！")
	}

	err = model.ResponseAdd(&inf)
	if err != nil {
		return c.JSON(301,err.Error())
	}
	return c.JSON(200,"回答成功！")
}

//点赞
func LikecountAdd (c echo.Context)error{
	inf := model.Response{}
	err := c.Bind(&inf)

	if err != nil{
		return c.JSON(300,"错误！")
	}
	err = model.LikecountAdd(&inf)
	if err != nil {
		return c.JSON(301,err.Error())
	}
	return c.JSON(200,"点赞成功！")

}

//取消点赞
func LikecountCancel (c echo.Context)error{
	inf := model.Response{}
	err := c.Bind(&inf)
	if err != nil{
		return c.JSON(300,"错误！")
	}
	err = model.LikecountCancel(&inf)
	if err != nil {
		return c.JSON(301,err.Error())
	}
	return c.JSON(200,"取消成功！")
}

//踩
func DislikecountAdd (c echo.Context)error{
	inf := model.Response{}
	err := c.Bind(&inf)
	if err != nil{
		return c.JSON(300,"错误！")
	}
	err = model.DislikecountAdd(&inf)
	if err != nil {
		return c.JSON(301,err.Error())
	}
	return c.JSON(200,"踩成功！")
}

//取消踩
func DislikecountCancel (c echo.Context)error{
	inf := model.Response{}
	err := c.Bind(&inf)
	if err != nil{
		return c.JSON(300,"错误！")
	}
	err = model.DislikecountCancel(&inf)
	if err != nil {
		return c.JSON(301,err.Error())
	}
	return c.JSON(200,"取消成功！")
}

//写文章
func ArticleAdd (c echo.Context)error {
	inf := model.Article{}
	err :=c.Bind(&inf)
	if err != nil{
		return c.JSON(300,"错误！")
	}
	if inf.Title == ""{
		return c.JSON(300,"标题不能为空！")
	}
	if inf.Content == ""{
		return c.JSON(300,"内容不能为空！")
	}
	err = model.ArticleAdd(&inf)
	return c.JSON(200,"成功！")

}

//查询所有评论
func ResponseSelectAll(c echo.Context)error{
	id,_ := strconv.Atoi(c.Param("id"))
	mods,err :=model.ResponseSelectAll(id)
	if err != nil{
		return c.JSON(300,"未查询到评论！")
	}
	return c.JSON(200,mods)
}

//查询一个问题
func QuestionSelect1(c echo.Context)error{
	id,_ := strconv.Atoi(c.Param("id"))
	mod,err := model.QuestionSelect1(id)
	if err != nil {
		return c.JSON(300,"未查询到问题！")
	}
	return c.JSON(200,mod)
}

//热榜
func QuestionSelect10 (c echo.Context)error{
	mods,err := model.QuestionSelect10()
	if err != nil {
		return c.JSON(300,"未查询到数据")
	}
	return c.JSON(200,mods)
}

//查找所有粉丝
func Fans (c echo.Context)error{
	userid,err:= strconv.Atoi(c.Param("userid"))
	if err != nil{
		return c.JSON(400,"数据有误！")
	}
	mods,err2 := model.FansSelect(userid)
	if err2 !=nil {
		return c.JSON(300,err.Error())
	}
	return c.JSON(200,mods)
}
//查询粉丝数量
func FansCount (c echo.Context)error{
	userid,err:= strconv.Atoi(c.Param("userid"))
	if err != nil{
		return c.JSON(400,"数据有误！")
	}
	mods,err2 := model.FansCountSelect(userid)
	if err2 !=nil {
		return c.JSON(300,"未查询到数据！")
	}
	return c.JSON(200,mods)

}

//查找所有关注者
func Follows (c echo.Context)error{
	userid,err:= strconv.Atoi(c.Param("userid"))
	if err != nil{
		return c.JSON(400,"数据有误！")
	}
	mods,err2 := model.FollowsSelect(userid)
	if err2 !=nil {
		return c.JSON(300,"您没有关注任何人！")
	}
	return c.JSON(200,mods)
}

//查询关注者数量
func FollowsCount (c echo.Context)error{
	userid,err:= strconv.Atoi(c.Param("userid"))
	if err != nil{
		return c.JSON(400,"数据有误！")
	}
	mods,err2 := model.FollowsCountSelect(userid)
	if err2 !=nil {
		return c.JSON(300,"未查询到数据！")
	}
	return c.JSON(200,mods)
}

//查询用户的回答 （个人页面的回答）
func ResponseFindById (c echo.Context)error{
	responser_id,err := strconv.Atoi(c.Param("userid"))
	if err != nil{
		return c.JSON(400,"数据有误！")
	}
	mods,err2 := model.ResponseAllByuserid(responser_id)
	if err2 !=nil {
		return c.JSON(300,"未查询到数据！")
	}
	return c.JSON(200,mods)
}

//查询用户的问题 （个人页面的问题)

func FindQByQid (c echo.Context)error{
	QID,err := strconv.Atoi(c.Param("userid"))
	if err != nil{
		return c.JSON(400,"数据有误！")
	}
	mods,err2 := model.FindQByQid(QID)
	if err2 !=nil {
		return c.JSON(300,err.Error())
	}
	return c.JSON(200,mods)
}

//查询用户的文章
func FindArticleByUID (c echo.Context)error{
	UID,err := strconv.Atoi(c.Param("userid"))
	if err != nil{
		return c.JSON(400,"数据有误！")
	}
	mods,err2 := model.FindArticleByAID(UID)
	if err2 !=nil {
		return c.JSON(300,err.Error())
	}
	return c.JSON(200,mods)
}

//修改用户昵称
func ModifyName (c echo.Context)error {
	inf := model.User{}
	err :=c.Bind(&inf)
	if err != nil{
		 return c.JSON(300,err.Error())
	}
	if inf.Username == ""{
		return c.JSON(300,"昵称不能为空！")
	}
	err = model.ModifyName(&inf)
	return c.JSON(200,"修改成功！")
}

//写评论
func CommentAdd (c echo.Context)error{
	inf := model.Comment{}
	err :=c.Bind(&inf)
	if err != nil{
		return c.JSON(300,err.Error())
	}
	err = model.CommentAdd(&inf)
	return c.JSON(200,"评论成功！")
}

//查询问题的某一回答
func ResponseFind1 (c echo.Context)error{
	QID,err := strconv.Atoi(c.Param("id"))
	if err != nil{
		return c.JSON(400,"数据有误！")
	}
	mod,err2 := model.ResponseSelect1(QID)
	if err2 !=nil {
		return c.JSON(300,"未查询到数据！")
	}
	return c.JSON(200,mod)
}

//查询某一回答的所有根评论
func CommentFindAll (c echo.Context)error{
	RID,err := strconv.Atoi(c.Param("reply_id"))
	if err != nil{
		return c.JSON(400,err.Error())
	}
	mod,err2 := model.CommentFindAll(RID)
	if err2 !=nil {
		return c.JSON(300,"未查询到数据！")
	}
	return c.JSON(200,mod)
}

//查询某一根评论的所有子评论
func  CommentFindSONALL (c echo.Context)error{
	CID,err := strconv.Atoi(c.Param("comment_id"))
	if err != nil{
		return c.JSON(400,err.Error())
	}
	mod,err2 := model.CommentFindSON(CID)
	if err2 !=nil {
		return c.JSON(300,"未查询到数据！")
	}
	return c.JSON(200,mod)
}