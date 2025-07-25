package handlers

import (
	"net/http"
	"strconv"

	"Vortexia/internal/middleware"
	"Vortexia/internal/model"
	"Vortexia/internal/service"

	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {
	projectService service.ProjectService
}

// NewProjectHandler 创建项目处理器
func NewProjectHandler(projectService service.ProjectService) *ProjectHandler {
	return &ProjectHandler{projectService: projectService}
}

// Create 创建项目
// @Summary 创建项目
// @Description 创建新项目
// @Tags 项目
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body model.CreateProjectRequest true "创建项目请求"
// @Success 201 {object} model.APIResponse{data=model.Project}
// @Failure 400 {object} model.APIResponse
// @Failure 401 {object} model.APIResponse
// @Router /api/v1/projects [post]
func (h *ProjectHandler) Create(c *gin.Context) {
	var req model.CreateProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.APIResponse{
			Code:    http.StatusBadRequest,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	// 获取当前用户ID
	userID, exists := middleware.GetCurrentUserID(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, model.APIResponse{
			Code:    http.StatusUnauthorized,
			Message: "用户信息不存在",
		})
		return
	}

	project, err := h.projectService.Create(&req, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.APIResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, model.APIResponse{
		Code:    http.StatusCreated,
		Message: "创建成功",
		Data:    project,
	})
}

// GetByID 根据ID获取项目
// @Summary 根据ID获取项目
// @Description 根据ID获取项目详情
// @Tags 项目
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "项目ID"
// @Success 200 {object} model.APIResponse{data=model.Project}
// @Failure 400 {object} model.APIResponse
// @Failure 404 {object} model.APIResponse
// @Router /api/v1/projects/{id} [get]
func (h *ProjectHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.APIResponse{
			Code:    http.StatusBadRequest,
			Message: "无效的项目ID",
		})
		return
	}

	project, err := h.projectService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.APIResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	if project == nil {
		c.JSON(http.StatusNotFound, model.APIResponse{
			Code:    http.StatusNotFound,
			Message: "项目不存在",
		})
		return
	}

	c.JSON(http.StatusOK, model.APIResponse{
		Code:    http.StatusOK,
		Message: "获取成功",
		Data:    project,
	})
}

// Update 更新项目
// @Summary 更新项目
// @Description 更新项目信息
// @Tags 项目
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "项目ID"
// @Param request body model.Project true "项目信息"
// @Success 200 {object} model.APIResponse{data=model.Project}
// @Failure 400 {object} model.APIResponse
// @Router /api/v1/projects/{id} [put]
func (h *ProjectHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.APIResponse{
			Code:    http.StatusBadRequest,
			Message: "无效的项目ID",
		})
		return
	}

	var req model.Project
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.APIResponse{
			Code:    http.StatusBadRequest,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	req.ID = id
	if err := h.projectService.Update(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.APIResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.APIResponse{
		Code:    http.StatusOK,
		Message: "更新成功",
		Data:    req,
	})
}

// Delete 删除项目
// @Summary 删除项目
// @Description 删除项目
// @Tags 项目
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "项目ID"
// @Success 200 {object} model.APIResponse
// @Failure 400 {object} model.APIResponse
// @Router /api/v1/projects/{id} [delete]
func (h *ProjectHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.APIResponse{
			Code:    http.StatusBadRequest,
			Message: "无效的项目ID",
		})
		return
	}

	if err := h.projectService.Delete(id); err != nil {
		c.JSON(http.StatusBadRequest, model.APIResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.APIResponse{
		Code:    http.StatusOK,
		Message: "删除成功",
	})
}

// List 获取项目列表
// @Summary 获取项目列表
// @Description 获取项目列表
// @Tags 项目
// @Produce json
// @Security ApiKeyAuth
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页大小" default(20)
// @Success 200 {object} model.APIResponse{data=model.PaginationResponse}
// @Router /api/v1/projects [get]
func (h *ProjectHandler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	result, err := h.projectService.List(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.APIResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.APIResponse{
		Code:    http.StatusOK,
		Message: "获取成功",
		Data:    result,
	})
}

// GetMyProjects 获取当前用户的项目
// @Summary 获取我的项目
// @Description 获取当前用户拥有的项目
// @Tags 项目
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} model.APIResponse{data=[]model.Project}
// @Router /api/v1/projects/my [get]
func (h *ProjectHandler) GetMyProjects(c *gin.Context) {
	userID, exists := middleware.GetCurrentUserID(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, model.APIResponse{
			Code:    http.StatusUnauthorized,
			Message: "用户信息不存在",
		})
		return
	}

	projects, err := h.projectService.GetByOwner(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.APIResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.APIResponse{
		Code:    http.StatusOK,
		Message: "获取成功",
		Data:    projects,
	})
}
