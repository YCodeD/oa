package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/oa?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	//fmt.Println("连接数据库成功")
	db.AutoMigrate(
		&Announcement{},
		&EGMember{},
		&ExerciseGroup{},
		&ExerciseClock{},
		&Resignation{},
		&Reimbursement{},
		&Vehicle{},
		&Travel{},
		&Birthday{},
		&Training{})

	//设置连接池
	//空闲
	db.DB().SetMaxIdleConns(50)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	db.LogMode(true)
}
