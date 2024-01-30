package config

import (
	"gopkg.in/ini.v1"
	"log"
	"time"
)

// AppConfig 存储应用配置
type AppConfig struct {
	App       ApplicationConfig
	Database  DatabaseConfig
	JwtConfig JwtConfig
}

// 应用配置信息
type ApplicationConfig struct {
	Port      string
	PrefixUrl string
	RunMode   string
}

// 数据库配置信息
type DatabaseConfig struct {
	Type     string
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

type JwtConfig struct {
	Secret string
	Expire time.Duration
}

var Config AppConfig

func LoadConfig(path string) {
	err := ini.MapTo(&Config, path)
	if err != nil {
		log.Fatalf("Fail to read file: %v", err)
	}
}
