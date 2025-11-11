package config

import (
	"os"
)

// Config 应用配置结构
type Config struct {
	// Server 服务器配置
	Server ServerConfig
	// Database 数据库配置
	Database DatabaseConfig
}

// ServerConfig 服务器配置
type ServerConfig struct {
	// Port 服务端口
	Port string
	// Mode 运行模式：debug, release, test
	Mode string
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	// DSN 数据库连接字符串
	DSN string
}

// Load 加载配置
func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "9080"),
			Mode: getEnv("GIN_MODE", "debug"),
		},
		Database: DatabaseConfig{
			DSN: getEnv("DATABASE_DSN", "task-four.db"),
		},
	}
}

// getEnv 获取环境变量，如果不存在则返回默认值
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
