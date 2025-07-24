package model

import (
	"time"
)

// User 用户模型
type User struct {
	ID        int       `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"-" db:"password_hash"` // 不返回密码
	Role      string    `json:"role" db:"role"`
	IsActive  bool      `json:"is_active" db:"is_active"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Project 项目模型
type Project struct {
	ID          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	RepoURL     string    `json:"repo_url" db:"repo_url"`
	Branch      string    `json:"branch" db:"branch"`
	OwnerID     int       `json:"owner_id" db:"owner_id"`
	IsActive    bool      `json:"is_active" db:"is_active"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// Pipeline 流水线模型
type Pipeline struct {
	ID        int       `json:"id" db:"id"`
	ProjectID int       `json:"project_id" db:"project_id"`
	Name      string    `json:"name" db:"name"`
	Config    string    `json:"config" db:"config"` // YAML配置
	IsActive  bool      `json:"is_active" db:"is_active"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Build 构建模型
type Build struct {
	ID         int        `json:"id" db:"id"`
	PipelineID int        `json:"pipeline_id" db:"pipeline_id"`
	Branch     string     `json:"branch" db:"branch"`
	Commit     string     `json:"commit" db:"commit"`
	Status     string     `json:"status" db:"status"`
	StartedAt  time.Time  `json:"started_at" db:"started_at"`
	FinishedAt *time.Time `json:"finished_at,omitempty" db:"finished_at"`
	Duration   *int       `json:"duration,omitempty" db:"duration"` // 秒
	TriggerBy  int        `json:"trigger_by" db:"trigger_by"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
}

// BuildStep 构建步骤模型
type BuildStep struct {
	ID         int        `json:"id" db:"id"`
	BuildID    int        `json:"build_id" db:"build_id"`
	Name       string     `json:"name" db:"name"`
	Command    string     `json:"command" db:"command"`
	Status     string     `json:"status" db:"status"`
	Output     string     `json:"output" db:"output"`
	StartedAt  time.Time  `json:"started_at" db:"started_at"`
	FinishedAt *time.Time `json:"finished_at,omitempty" db:"finished_at"`
	Duration   *int       `json:"duration,omitempty" db:"duration"`
	StepOrder  int        `json:"step_order" db:"step_order"`
}

// BuildStatus 构建状态常量
const (
	BuildStatusPending  = "pending"
	BuildStatusRunning  = "running"
	BuildStatusSuccess  = "success"
	BuildStatusFailed   = "failed"
	BuildStatusCanceled = "canceled"
)

// StepStatus 步骤状态常量
const (
	StepStatusPending  = "pending"
	StepStatusRunning  = "running"
	StepStatusSuccess  = "success"
	StepStatusFailed   = "failed"
	StepStatusSkipped  = "skipped"
)

// UserRole 用户角色常量
const (
	RoleAdmin = "admin"
	RoleUser  = "user"
)

// CreateUserRequest 创建用户请求
type CreateUserRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"required,oneof=admin user"`
}

// LoginRequest 登录请求
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse 登录响应
type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

// CreateProjectRequest 创建项目请求
type CreateProjectRequest struct {
	Name        string `json:"name" binding:"required,min=1,max=100"`
	Description string `json:"description" binding:"max=500"`
	RepoURL     string `json:"repo_url" binding:"required,url"`
	Branch      string `json:"branch" binding:"required"`
}

// CreatePipelineRequest 创建流水线请求
type CreatePipelineRequest struct {
	ProjectID int    `json:"project_id" binding:"required"`
	Name      string `json:"name" binding:"required,min=1,max=100"`
	Config    string `json:"config" binding:"required"`
}

// TriggerBuildRequest 触发构建请求
type TriggerBuildRequest struct {
	PipelineID int    `json:"pipeline_id" binding:"required"`
	Branch     string `json:"branch" binding:"required"`
	Commit     string `json:"commit"`
}

// APIResponse 统一API响应格式
type APIResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// PaginationResponse 分页响应
type PaginationResponse struct {
	Items      interface{} `json:"items"`
	Total      int         `json:"total"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	TotalPages int         `json:"total_pages"`
} 