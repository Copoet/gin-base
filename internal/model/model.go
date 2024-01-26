package model

import (
	"fmt"
	"gin-base/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

type BaseModel struct {
	ID         int    `gorm:"primary_key" json:"id"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

func Init() {
	//sql日志
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出到标准输出）
		logger.Config{
			SlowThreshold: time.Second, // 慢SQL阈值
			LogLevel:      logger.Info, // 日志级别
			Colorful:      true,        // 彩色打印
		},
	)

	var err error
	dsn := buildDSN()

	switch config.Config.Database.Type {
	case "mysql":
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: newLogger,
		})
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
