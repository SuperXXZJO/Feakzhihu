package router

import (
	"github.com/labstack/echo"
	"zhihu/control"
)

func Api (api *echo.Group){
	api.POST("/questionadd",control.QuestionAdd) //提问
	api.POST("/responseadd",control.ResponseAdd)  //回答
	api.POST("/write",control.ArticleAdd) //写文章
	api.GET("/writeadd",control.ArticleView) //写文章的详情页
	api.GET("/index.html/Hot",control.QuestionSelect10)//主页面热榜
	api.GET("/index.html",control.QuestionView) //主页面推荐
	api.GET("/personal/:userid",control.Peopleview)//个人页面
	api.GET("/personal/:userid/question",control.FindQByQid) //个人页面的问题
	api.GET("/personal/:userid/response",control.ResponseFindById) //个人页面的回答
	api.GET("/personal/:userid/article",control.FindArticleByUID) //个人页面的回答
	api.GET("/personal/:userid/fans",control.Fans) //查找粉丝
	api.GET("/personal/:userid/fanscount",control.FansCount)//查找粉丝数量
	api.GET("/personal/:userid/follows",control.Follows) //查找关注者
	api.GET("/personal/:userid/followscount",control.FollowsCount)//查找关注者数量
	api.POST("/personal/:userid/modify",control.ModifyName) //修改昵称

	questionselect := api.Group("/Q",Hits)
	questionselect.GET("/TopicM/:id",control.QuestionSelect1)//问题详情页面
	questionselect.GET("/TopicM/:id/response",control.ResponseSelectAll)//问题页面的所有回答
	questionselect.GET("/TopicM/:id/response/:response_id",control.ResponseFind1)//查询问题的某一回答详情
	questionselect.POST("/TopicM/:id/response/:response_id/likecountadd",control.LikecountAdd) //点赞
	questionselect.POST("/TopicM/:id/response/:response_id/likecountcancel",control.LikecountCancel) //取消点赞
	questionselect.POST("/TopicM/:id/response/:response_id/dislikecountadd",control.DislikecountAdd) //踩
	questionselect.POST("/TopicM/:id/response/:response_id/dislikecountcancel",control.DislikecountCancel) //取消踩

	questionselect.POST("TopicM/:id/response/:response_id/comment/:comment_id",control.CommentAdd)//写评论
	questionselect.GET("TopicM/:id/response/:response_id/commentAll/:reply_id",control.CommentFindAll)//查找某一回答的所有根评论
	questionselect.GET("TopicM/:id/response/:response_id/commentAll/:reply_id/:comment_id",control.CommentFindSONALL)//查找某一根的所有子评论
}

