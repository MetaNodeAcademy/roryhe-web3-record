package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rory7/task-four/handler"
	"github.com/rory7/task-four/middleware"
)

// SetupRoutes 设置路由
func SetupRoutes() *gin.Engine {
	// 创建 Gin 引擎
	r := gin.Default()

	// 注册全局中间件
	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())
	r.Use(middleware.CORS())

	// 健康检查接口
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// 创建用户处理器
	userHandler := handler.NewUserHandler()

	// API 路由组
	api := r.Group("/api")
	{
		loginApi := api.Group("/login")
		{
			loginApi.POST("", userHandler.Login)
			loginApi.POST("/register", userHandler.CreateUser)
		}

		// 用户相关路由
		users := api.Group("/users")
		users.Use(middleware.JWTAuth())
		{
			users.GET("", userHandler.GetAllUsers) // 获取所有用户
		}

		//文章相关路由
		postRoutes := api.Group("/posts")
		postRoutes.Use(middleware.JWTAuth())
		{
			postRoutes.POST("")    //创建文章
			postRoutes.GET("")     //获取单个文章
			postRoutes.GET("/all") //获取全部文章
			postRoutes.PUT("")     // update 文章
			postRoutes.DELETE("")  // delete 文章
		}
	}

	return r
}
