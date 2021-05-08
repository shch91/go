package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

// UserInfo 用户信息
type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

type Like struct {
	ID        int    `gorm:"primary_key"`
	Ip        string `gorm:"type:varchar(20);not null;index:ip_idx"`
	Ua        string `gorm:"type:varchar(256);not null;"`
	Title     string `gorm:"type:varchar(128);not null;index:title_idx"`
	Hash      uint64 `gorm:"unique_index:hash_idx;"`
	CreatedAt time.Time
}

func main() {
	db, err := gorm.Open("mysql", "root:Yy&sh2021@(127.0.0.1:3306)/Test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.DB().SetMaxIdleConns(2)
	db.DB().SetMaxOpenConns(10)
	// 自动迁移
	db.AutoMigrate(&UserInfo{})


	u1 := UserInfo{Name: "枯藤", Gender: "男", Hobby: "篮球"}

	// 创建记录
	db.Create(&u1)

	// 查询
	var u = new(UserInfo)
	db.First(u)
	fmt.Printf("%#v\n", u)
	var uu UserInfo
	db.Find(&uu, "hobby=?", "足球")
	fmt.Printf("%#v\n", uu)
	// 更新
	db.Model(&u).Update("hobby", "双色球")
	// 删除
	//db.Delete(&u)
}
