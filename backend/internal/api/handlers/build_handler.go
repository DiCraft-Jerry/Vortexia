package handlers

import (
	"net/http"

	"simple-ci/internal/model"
	"simple-ci/internal/service"

	"github.com/gin-gonic/gin"
)

type BuildHandler struct {
	buildService service.BuildService
}

// NewBuildHandler 创建构建处理器
func NewBuildHandler(buildService service.BuildService) *BuildHandler {
	return &BuildHandler{buildService: buildService}
}

// List 获取构建列表
func (h *BuildHandler) List(c *gin.Context) {
	c.JSON(http.StatusOK, model.APIResponse{
		Code:    http.StatusOK,
		Message: "功能开发中",
		Data:    []interface{}{},
	})
}

// Create 创建构建
func (h *BuildHandler) Create(c *gin.Context) {
	c.JSON(http.StatusOK, model.APIResponse{
		Code:    http.StatusOK,
		Message: "功能开发中",
	})
}

// GetByID 根据ID获取构建
func (h *BuildHandler) GetByID(c *gin.Context) {
	c.JSON(http.StatusOK, model.APIResponse{
		Code:    http.StatusOK,
		Message: "功能开发中",
	})
}

// UpdateStatus 更新构建状态
func (h *BuildHandler) UpdateStatus(c *gin.Context) {
	c.JSON(http.StatusOK, model.APIResponse{
		Code:    http.StatusOK,
		Message: "功能开发中",
	})
}

// GetSteps 获取构建步骤
func (h *BuildHandler) GetSteps(c *gin.Context) {
	c.JSON(http.StatusOK, model.APIResponse{
		Code:    http.StatusOK,
		Message: "功能开发中",
		Data:    []interface{}{},
	})
}

// GetByPipeline 根据流水线获取构建
func (h *BuildHandler) GetByPipeline(c *gin.Context) {
	c.JSON(http.StatusOK, model.APIResponse{
		Code:    http.StatusOK,
		Message: "功能开发中",
		Data:    []interface{}{},
	})
}

// WatchLogs WebSocket日志流
func (h *BuildHandler) WatchLogs(c *gin.Context) {
	c.JSON(http.StatusOK, model.APIResponse{
		Code:    http.StatusOK,
		Message: "WebSocket功能开发中",
	})
} 