package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//gorm demo4
//1.定义模型
type User struct {
	gorm.Model
	Name string
	Age  int64
}

func main() {
	// 连接Mysql数据库
	db, err := gorm.Open("mysql", "root:qwe2476979217@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//2. 把模型与数据库中的表对应起来
	db.AutoMigrate(&User{})
	//3. 创建
	u1 := User{Name: "皮肤", Age: 18}
	db.Create(&u1)
}
