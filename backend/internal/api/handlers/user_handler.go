package handlers

import (
	"net/http"
	"strconv"

	"simple-ci/internal/middleware"
	"simple-ci/internal/model"
	"simple-ci/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

// NewUserHandler 创建用户处理器
func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// GetProfile 获取当前用户信息
// @Summary 获取当前用户信息
// @Description 获取当前登录用户的详细信息
// @Tags 用户
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} model.APIResponse{data=model.User}
// @Failure 401 {object} model.APIResponse
// @Router /api/v1/users/profile [get]
func (h *UserHandler) GetProfile(c *gin.Context) {
	user, exists := middleware.GetCurrentUser(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, model.APIResponse{
			Code:    http.StatusUnauthorized,
			Message: "用户信息不存在",
		})
		return
	}

	c.JSON(http.StatusOK, model.APIResponse{
		Code:    http.StatusOK,
		Message: "获取成功",
		Data:    user,
	})
}

// UpdateProfile 更新当前用户信息
// @Summary 更新当前用户信息
// @Description 更新当前登录用户的信息
// @Tags 用户
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body model.User true "用户信息"
// @Success 200 {object} model.APIResponse{data=model.User}
// @Failure 400 {object} model.APIResponse
// @Failure 401 {object} model.APIResponse
// @Router /api/v1/users/profile [put]
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	currentUser, exists := middleware.GetCurrentUser(c)
	if !exists {
		c.JSON(http.StatusUnauthorized, model.APIResponse{
			Code:    http.StatusUnauthorized,
			Message: "用户信息不存在",
		})
		return
	}

	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.APIResponse{
			Code:    http.StatusBadRequest,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	// 只允许更新自己的信息
	req.ID = currentUser.ID
	req.Role = currentUser.Role // 不允许修改角色

	if err := h.userService.Update(&req); err != nil {
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

// Create 创建用户（管理员）
// @Summary 创建用户
// @Description 管理员创建新用户
// @Tags 用户
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body model.CreateUserRequest true "创建用户请求"
// @Success 201 {object} model.APIResponse{data=model.User}
// @Failure 400 {object} model.APIResponse
// @Failure 403 {object} model.APIResponse
// @Router /api/v1/users [post]
func (h *UserHandler) Create(c *gin.Context) {
	var req model.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.APIResponse{
			Code:    http.StatusBadRequest,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	user, err := h.userService.Create(&req)
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
		Data:    user,
	})
}

// GetByID 根据ID获取用户（管理员）
// @Summary 根据ID获取用户
// @Description 管理员根据ID获取用户信息
// @Tags 用户
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "用户ID"
// @Success 200 {object} model.APIResponse{data=model.User}
// @Failure 400 {object} model.APIResponse
// @Failure 403 {object} model.APIResponse
// @Failure 404 {object} model.APIResponse
// @Router /api/v1/users/{id} [get]
func (h *UserHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.APIResponse{
			Code:    http.StatusBadRequest,
			Message: "无效的用户ID",
		})
		return
	}

	user, err := h.userService.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.APIResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, model.APIResponse{
			Code:    http.StatusNotFound,
			Message: "用户不存在",
		})
		return
	}

	c.JSON(http.StatusOK, model.APIResponse{
		Code:    http.StatusOK,
		Message: "获取成功",
		Data:    user,
	})
}

// Update 更新用户（管理员）
// @Summary 更新用户
// @Description 管理员更新用户信息
// @Tags 用户
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "用户ID"
// @Param request body model.User true "用户信息"
// @Success 200 {object} model.APIResponse{data=model.User}
// @Failure 400 {object} model.APIResponse
// @Failure 403 {object} model.APIResponse
// @Router /api/v1/users/{id} [put]
func (h *UserHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.APIResponse{
			Code:    http.StatusBadRequest,
			Message: "无效的用户ID",
		})
		return
	}

	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.APIResponse{
			Code:    http.StatusBadRequest,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	req.ID = id
	if err := h.userService.Update(&req); err != nil {
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

// Delete 删除用户（管理员）
// @Summary 删除用户
// @Description 管理员删除用户
// @Tags 用户
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "用户ID"
// @Success 200 {object} model.APIResponse
// @Failure 400 {object} model.APIResponse
// @Failure 403 {object} model.APIResponse
// @Router /api/v1/users/{id} [delete]
func (h *UserHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.APIResponse{
			Code:    http.StatusBadRequest,
			Message: "无效的用户ID",
		})
		return
	}

	if err := h.userService.Delete(id); err != nil {
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

// List 获取用户列表（管理员）
// @Summary 获取用户列表
// @Description 管理员获取用户列表
// @Tags 用户
// @Produce json
// @Security ApiKeyAuth
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页大小" default(20)
// @Success 200 {object} model.APIResponse{data=model.PaginationResponse}
// @Failure 403 {object} model.APIResponse
// @Router /api/v1/users [get]
func (h *UserHandler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	result, err := h.userService.List(page, pageSize)
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