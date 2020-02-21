package model

import (
	"errors"
	"time"
)

type Question struct {
	Id int            //问题的id
	Question string   //问题
	Detail string      //详情
	Hits int        //点击量
	Questioner_id int //提问者
	Utime int64    //更新时间
}
//提问
func QuestionAdd (mod *Question) error{
	//开启事务
	sp,err :=Db.Begin()
	if err != nil{
		return err
	}
	mod.Utime = time.Now().Unix()

	result,err := sp.Exec("insert into questions(question,detail,utime,questioner_id)values (?,?,?,?)",mod.Question,mod.Detail,mod.Utime,mod.Questioner_id)
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

//查询一个问题
func QuestionSelect1 (id int)([]Question,error){
	mod := make([]Question,0)
	err := Db.Select(&mod,"select * from questions where id = ?  ",id)
	return mod,err
}


//热榜 查询多个问题
func QuestionSelect10 ()([]Question,error){
	mods := make([]Question,0)
	err := Db.Select(&mods,"select * from questions order by hits desc limit 10")
	return mods,err
}




//增加点击量
func QuestionHitsAdd (id int)error{
	//开启事务
	sp,err := Db.Begin()
	if err != nil{
		return err
	}
	result,err := sp.Exec("update questions set hits = hits+1 where id = ?",id)
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

//推荐功能(查询问题)
func QuestionSelectByutime ()([]Question,error){
	mods := make([]Question,0)
	err := Db.Select(&mods,"select * from questions order by utime desc  ")

	return mods,err
}


//通过提问者id查询问题
func FindQByQid (QID int)([]Question,error){
	mods := make([]Question,0)
	err := Db.Select(&mods,"select * from questions where questioner_id = ? ",QID)
	return mods,err
}