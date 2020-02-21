package model

import (
	"github.com/jmoiron/sqlx"
	"log"
)
//连接数据库
var Db *sqlx.DB

func init(){
	db,err := sqlx.Open("mysql","root:root@tcp(127.0.0.1:3306)/zhihu?charset=utf8")
	if err != nil{
		log.Fatal(err.Error())
	}
	if err = db.Ping() ;err != nil{
		log.Fatal(err.Error())
	}
	Db=db

}