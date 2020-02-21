package model

import (
	"errors"
	"time"
)

type Comment struct {
	Id int  //主键 （没啥用）
	Comment_id int // 根评论id 就是回答(response)的id
	Reply_id int //回复目标用户Id
	Reply_type string //评论的类型 comment 或 response
	Content  string //评论内容
	From_uid int    //回复用户id
	To_uid   int    //目标用户id
	Utime    int64
}


//写评论
func CommentAdd(mod *Comment)error {
	//开启事务
	sp,err :=Db.Begin()
	if err != nil{
		return err
	}
	mod.Utime = time.Now().Unix()
	mod.Comment_id =mod.Comment_id +1
	result,err := sp.Exec("insert into comment(Comment_id,Reply_id,Content,From_uid,to_uid)values (?,?,?,?,?)",mod.Comment_id,mod.Reply_id,mod.Content,mod.From_uid,mod.To_uid,mod.Utime)
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

//查询某一回答的一条根评论
func CommentFind1 (inf *Comment)( *Comment,error){
	mod:= &Comment{}
	err := Db.Select(&mod,"select * from comment where comment_id =? and reply_type = comment",inf.Comment_id)
	return mod,err
}


//查询某回答的所有根评论并按照时间排序
func CommentFindAll (reply_id int)([]Comment,error){
	mod := make([]Comment,0)
	err := Db.Select(&mod,"select * from comment where reply_id = ? and reply_type = comment order by utime desc ",reply_id)
	return mod,err
}

//查询某一根评论的所有子评论
func CommentFindSON (comment_id int)([]Comment,error){
	mods := make([]Comment,0)
	err := Db.Select(&mods,"select * from comment where comment_id = ? ",comment_id)
	return mods,err
}