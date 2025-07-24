package routes

import (
	"simple-ci/internal/api/handlers"
	"simple-ci/internal/middleware"
	"simple-ci/internal/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

// SetupRoutes 配置路由
func SetupRoutes(services *service.Services, logger *zap.Logger) *gin.Engine {
	r := gin.New()

	// 中间件
	r.Use(middleware.Logger(logger))
	r.Use(middleware.Recovery(logger))
	r.Use(middleware.CORS())

	// 初始化handlers
	authHandler := handlers.NewAuthHandler(services.Auth)
	userHandler := handlers.NewUserHandler(services.User)
	projectHandler := handlers.NewProjectHandler(services.Project)
	pipelineHandler := handlers.NewPipelineHandler(services.Pipeline)
	buildHandler := handlers.NewBuildHandler(services.Build)

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "message": "Simple CI/CD is running"})
	})

	// Swagger文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API路由组
	api := r.Group("/api/v1")

	// 认证相关路由（无需认证）
	auth := api.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
	}

	// 需要认证的路由
	protected := api.Group("/")
	protected.Use(middleware.JWTAuth(services.Auth))

	// 用户管理路由
	users := protected.Group("/users")
	{
		users.GET("/profile", userHandler.GetProfile)
		users.PUT("/profile", userHandler.UpdateProfile)
		users.GET("/", middleware.AdminRequired(), userHandler.List)
		users.POST("/", middleware.AdminRequired(), userHandler.Create)
		users.GET("/:id", middleware.AdminRequired(), userHandler.GetByID)
		users.PUT("/:id", middleware.AdminRequired(), userHandler.Update)
		users.DELETE("/:id", middleware.AdminRequired(), userHandler.Delete)
	}

	// 项目管理路由
	projects := protected.Group("/projects")
	{
		projects.GET("/", projectHandler.List)
		projects.POST("/", projectHandler.Create)
		projects.GET("/:id", projectHandler.GetByID)
		projects.PUT("/:id", projectHandler.Update)
		projects.DELETE("/:id", projectHandler.Delete)
		projects.GET("/my", projectHandler.GetMyProjects)
	}

	// 流水线管理路由
	pipelines := protected.Group("/pipelines")
	{
		pipelines.GET("/", pipelineHandler.List)
		pipelines.POST("/", pipelineHandler.Create)
		pipelines.GET("/:id", pipelineHandler.GetByID)
		pipelines.PUT("/:id", pipelineHandler.Update)
		pipelines.DELETE("/:id", pipelineHandler.Delete)
		pipelines.GET("/project/:project_id", pipelineHandler.GetByProject)
	}

	// 构建管理路由
	builds := protected.Group("/builds")
	{
		builds.GET("/", buildHandler.List)
		builds.POST("/", buildHandler.Create)
		builds.GET("/:id", buildHandler.GetByID)
		builds.PUT("/:id/status", buildHandler.UpdateStatus)
		builds.GET("/:id/steps", buildHandler.GetSteps)
		builds.GET("/pipeline/:pipeline_id", buildHandler.GetByPipeline)
	}

	// WebSocket路由（实时日志）
	ws := protected.Group("/ws")
	{
		ws.GET("/builds/:id/logs", buildHandler.WatchLogs)
	}

	return r
} 