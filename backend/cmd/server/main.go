package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"

	"simple-ci/internal/api/routes"
	"simple-ci/internal/config"
	"simple-ci/internal/repository"
	"simple-ci/internal/service"
	"simple-ci/pkg/logger"

	"github.com/gin-gonic/gin"
)

// @title Simple CI/CD API
// @version 1.0
// @description A simple CI/CD pipeline tool API
// @host localhost:8080
// @BasePath /api/v1
// @schemes http
func main() {
	// 设置内存优化
	debug.SetGCPercent(20)

	// 初始化日志
	logger := logger.New()
	defer logger.Sync()

	// 加载配置
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// 初始化数据库连接
	db, err := repository.NewPostgresDB(cfg.Database)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// 初始化Redis连接
	redisClient, err := repository.NewRedisClient(cfg.Redis)
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}
	defer redisClient.Close()

	// 初始化仓库层
	repos := repository.NewRepositories(db, redisClient)

	// 初始化服务层
	services := service.NewServices(repos)

	// 设置Gin模式
	gin.SetMode(cfg.Server.Mode)

	// 初始化路由
	router := routes.SetupRoutes(services, logger)

	// 创建HTTP服务器
	srv := &http.Server{
		Addr:    ":" + cfg.Server.Port,
		Handler: router,
	}

	// 启动服务器
	go func() {
		logger.Info("Starting server on port " + cfg.Server.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutting down server...")

	// 5秒的超时上下文
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 优雅关闭
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	logger.Info("Server exiting")
} 