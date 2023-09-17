package main

import (
	"bubble/dao"
	"bubble/models"
	"bubble/routers"
	"fmt"
)

func main() {
	// 创建数据库
	// sql: CREATE DATABASE bubble;
	//连接数据库
	err := dao.InitMysql()
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer dao.Close() // 程序退出关闭数据库连接

	// 模型绑定
	dao.DB.AutoMigrate(&models.Todo{}) //todos
	r := routers.SetupRouter()
	r.Run()
}
