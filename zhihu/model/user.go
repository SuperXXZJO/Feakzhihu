package model

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)
type LOGIN struct {
	Phone int
	Password string
}

type User struct {
	Userid int       //id
	Username string //昵称
	Phone int        //手机号
	Password string   //密码

}

//token
type UserToken struct {
	Userid int
	Username string
	jwt.StandardClaims
}

//关注者 (粉丝)
type Fans struct {
	userid  int
	Fansid  int
}

//关注了
type  Follows struct {
	Userid    int
	Followsid  int
}


//登录
func Login(phone int) (User,error){
	mod := User{}
	err := Db.Get(&mod,"select * from user where phone = ?",phone)
	return mod,err
}

//注册 signup
func Signup(mod *User)error{
	//开启事务
	sp,err :=Db.Begin()
	if err != nil{
		return err
	}
	result,err := sp.Exec("insert into user(phone,password)values (?,?)",mod.Phone,mod.Password)
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

//查找所有粉丝 FansSelect
func  FansSelect (userid int)([]Fans,error){
	mods := make([]Fans,0)
	err := Db.Select(&mods,"select * from fans where userid = ?",userid)
	return mods,err
}

//查找粉丝数量 FansCountSelect
func FansCountSelect (userid int)([]int,error){
	mods:=make([]int,0)
	err := Db.Select(&mods,"select count (fansid) from fans where userid = ?",userid)
	return mods,err
}


//查找所有关注的人
func FollowsSelect (userid int)([]Follows,error){
	mods := make([]Follows,0)
	err := Db.Select(&mods,"select * from follows where userid = ?",userid)
	return mods,err
}

//查找关注的人的数量
func FollowsCountSelect (userid int)(int,error){
	var mods int
	err := Db.Select(&mods,"select count (followsid) from follows where userid = ?",userid)
	return mods,err
}



//通过userid查找个人信息
func FindByUserid (userid int )( []User,error){
	mod := make([]User,0)
	err := Db.Select(&mod,"select  * from user where userid = ? limit 1",userid)
	return mod,err
}

//修改昵称
func ModifyName (mod *User)error {
	//开启事务
	sp, err := Db.Begin()
	if err != nil {
		return err
	}
	result, err := sp.Exec("update user set username = ? where userid = ?", mod.Username, mod.Userid)
	if err != nil {
		//回滚
		sp.Rollback()
		return err
	}
	rows, _ := result.RowsAffected()
	if rows < 1 {
		//回滚
		sp.Rollback()
		return errors.New("rows affected < 1")
	}
	//提交
	sp.Commit()
	return nil
}
