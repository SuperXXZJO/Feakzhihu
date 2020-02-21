package model

import (
	"errors"
	"log"
	"time"
)

type Response struct{
	Id int         //回答的id
	Responser_id int  //回答者的id
	Content string    //回答内容
	Question_id int   //回答问题的id
	Utime  int64   //回答时间
	Likecount int     //点赞数
	Dislikecount int   //踩数
}
//添加评论
func ResponseAdd (mod *Response)error{
	//开启事务
	sp,err :=Db.Begin()
	if err != nil{
		return err
	}
	mod.Utime = time.Now().Unix()
	result,err := sp.Exec("insert into response(Question_id,content,utime,Responser_id)values (?,?)",mod.Question_id,mod.Content,mod.Utime,mod.Responser_id)
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

//查询问题所有回答
func ResponseSelectAll(question_id int)([]Response,error){
	mods := make([]Response,0)
	err := Db.Select(&mods,"select * from response where question_id = ? order by likecount desc  ",question_id)
	return mods,err
}
//查询问题的某一回答
func ResponseSelect1 (question_id int)([]Response,error){
	mod := make([]Response,0)
	err :=Db.Select(&mod,"select * from response where question_id = ? limit 1",question_id)
	return mod,err
}




//查询用户所有回答 按照时间排序
func ResponseAllByuserid (userid int)([]Response,error){
	mods := make([]Response,0,10)
	err := Db.Select(&mods,"select * from response where responser_id = ? order by utime desc ",userid)
	return mods,err
}


//点赞
func LikecountAdd (mod *Response)error{
	//开启事务
	sp,err := Db.Begin()
	if err != nil{
		return err
	}

	result,err := sp.Exec("update response set likecount=likecount+1 where id = ?",mod.Id)
	if err != nil{
		//回滚
		sp.Rollback()
		return err
	}
	rows,err2:= result.RowsAffected()
	if err2 != nil {
		log.Println(err2)
	}

	if rows <1{
		//回滚
		sp.Rollback()
		return errors.New("rows affected < 1")
	}
	//提交
	sp.Commit()
	return nil
}

//取消点赞
func LikecountCancel (mod *Response)error{
	//开启事务
	sp,err := Db.Begin()
	if err != nil{
		return err
	}
	result,err := sp.Exec("update response set likecount = likecount-1 where id = ?",mod.Id)
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

//踩
func DislikecountAdd (mod *Response)error{
	//开启事务
	sp,err := Db.Begin()
	if err != nil{
		return err
	}
	result,err := sp.Exec("update response set dislikecount = dislikecount+1 where (id = ?)values (?)",mod.Id)
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

//取消踩
func DislikecountCancel (mod *Response)error{
	//开启事务
	sp,err := Db.Begin()
	if err != nil{
		return err
	}
	result,err := sp.Exec("update response set dislikecount = dislikecount-1 where id = ?",mod.Id)
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

