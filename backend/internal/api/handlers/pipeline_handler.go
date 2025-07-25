package handlers

import (
	"net/http"

	"Vortexia/internal/model"
	"Vortexia/internal/service"

	"github.com/gin-gonic/gin"
)

type PipelineHandler struct {
	pipelineService service.PipelineService
}

// NewPipelineHandler 创建流水线处理器
func NewPipelineHandler(pipelineService service.PipelineService) *PipelineHandler {
	return &PipelineHandler{pipelineService: pipelineService}
}

// List 获取流水线列表
func (h *PipelineHandler) List(c *gin.Context) {
	c.JSON(http.StatusOK, model.APIResponse{
		Code:    http.StatusOK,
		Message: "功能开发中",
		Data:    []interface{}{},
	})
}

// Create 创建流水线
func (h *PipelineHandler) Create(c *gin.Context) {
	c.JSON(http.StatusOK, model.APIResponse{
		Code:    http.StatusOK,
		Message: "功能开发中",
	})
}

// GetByID 根据ID获取流水线
func (h *PipelineHandler) GetByID(c *gin.Context) {
	c.JSON(http.StatusOK, model.APIResponse{
		Code:    http.StatusOK,
		Message: "功能开发中",
	})
}

// Update 更新流水线
func (h *PipelineHandler) Update(c *gin.Context) {
	c.JSON(http.StatusOK, model.APIResponse{
		Code:    http.StatusOK,
		Message: "功能开发中",
	})
}

// Delete 删除流水线
func (h *PipelineHandler) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, model.APIResponse{
		Code:    http.StatusOK,
		Message: "功能开发中",
	})
}

// GetByProject 根据项目获取流水线
func (h *PipelineHandler) GetByProject(c *gin.Context) {
	c.JSON(http.StatusOK, model.APIResponse{
		Code:    http.StatusOK,
		Message: "功能开发中",
		Data:    []interface{}{},
	})
}
