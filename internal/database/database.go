package database

import (
	"fmt"
	"gin-base/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Init() {
	var err error
	dsn := buildDSN()

	switch config.Config.Database.Type {
	case "mysql":
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		//可添加不同的数据库驱动
	default:
		log.Fatalf("不支持的数据库类型: %s", config.Config.Database.Type)
	}
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}
}

func buildDSN() string {
	dbConfig := config.Config.Database
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Name)
}
