package repository

import (
	"database/sql"
	"fmt"

	"simple-ci/internal/config"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
)

// NewPostgresDB 创建PostgreSQL数据库连接
func NewPostgresDB(cfg config.DatabaseConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// 连接池配置（针对1核2GB优化）
	db.SetMaxOpenConns(5)  // 最大连接数
	db.SetMaxIdleConns(2)  // 最大空闲连接数
	db.SetConnMaxLifetime(300) // 连接最大生存时间（5分钟）

	return db, nil
}

// NewRedisClient 创建Redis客户端
func NewRedisClient(cfg config.RedisConfig) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: 10, // 连接池大小
	})

	// 测试连接
	_, err := client.Ping(client.Context()).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return client, nil
} 