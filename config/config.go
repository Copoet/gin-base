package config

import (
	"gopkg.in/ini.v1"
	"log"
)

// AppConfig 存储应用配置
type AppConfig struct {
	App      ApplicationConfig
	Database DatabaseConfig
}

//应用配置信息
type ApplicationConfig struct {
	Port string
}

//数据库配置信息
type DatabaseConfig struct {
	Type     string
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

var Config AppConfig

func LoadConfig(path string) {
	err := ini.MapTo(&Config, path)
	if err != nil {
		log.Fatalf("Fail to read file: %v", err)
	}
}
