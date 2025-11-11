package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rory7/task-four/config"
	"github.com/rory7/task-four/database"
	"github.com/rory7/task-four/routes"
)

func main() {
	// 加载配置
	cfg := config.Load()

	// 设置 Gin 运行模式
	gin.SetMode(cfg.Server.Mode)

	// 初始化数据库
	if err := database.InitDB(cfg.Database.DSN); err != nil {
		log.Fatalf("数据库初始化失败: %v", err)
	}

	// 设置路由
	r := routes.SetupRoutes()

	// 启动服务器
	addr := fmt.Sprintf(":%s", cfg.Server.Port)
	log.Printf("服务器启动在端口 %s", cfg.Server.Port)
	if err := r.Run(addr); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
