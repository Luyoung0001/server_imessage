package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"server_imessage/models"
)

// 这里就是将模板导入数据库的一个简单案例
// 前提是创建好了一个数据库,数据库的名字暂定义为 "ginchat"
// 这个是用来实现模型绑定到数据的功能,只运行一次

var (
	DBGinChat *gorm.DB
)

func initMySQLGinChat() (err error) {
	// 连接数据库
	dsn := "root:791975457@qqCom@tcp(rm-cn-pe33bz99b000emxo.rwlb.rds.aliyuncs.com:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"
	DBGinChat, err = gorm.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 判断是否连通
	return DBGinChat.DB().Ping()
}
func main() {
	// 连接数据库
	err := initMySQLGinChat()
	if err != nil {
		panic(err)
	}
	// 模型绑定
	DBGinChat.AutoMigrate(&models.UserBasic{})
	DBGinChat.AutoMigrate(&models.Contact{})
	DBGinChat.AutoMigrate(&models.Message{})
	DBGinChat.AutoMigrate(&models.GroupBasic{})
	DBGinChat.AutoMigrate(&models.Community{})
	defer func(DB *gorm.DB) {
		err := DB.Close()
		if err != nil {
			panic(err)
		}
	}(DBGinChat)
}
