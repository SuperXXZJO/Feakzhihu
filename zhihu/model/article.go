package model

import (
	"errors"
	"time"
)

type Article struct {
	Id int
	Userid int
	Title string
	Content string
	Utime int64
}

//写文章

func ArticleAdd (mod *Article)error {
	//开启事务
	sp,err :=Db.Begin()
	if err != nil{
		return err
	}
	mod.Utime = time.Now().Unix()
	result,err := sp.Exec("insert into article (userid,title,content,utime)",mod.Userid,mod.Title,mod.Content,mod.Utime)
	if err != nil{
		//回滚
		sp.Rollback()
		return err
	}
	rows,_ := result.RowsAffected()
	if rows <1{
		//回滚
		sp.Rollback()
		return errors.New("rows affected < 1")
	}
	//提交
	sp.Commit()
	return nil
}

//根据作者查询文章
func FindArticleByAID (userid int)([]Article,error){
	mods := make([]Article,0)
	err := Db.Select(&mods,"select * from article where userid = ?",userid)
	return mods,err
}
