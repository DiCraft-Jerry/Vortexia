package handlers

import (
	"net/http"

	"Vortexia/internal/model"
	"Vortexia/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService service.AuthService
}

// NewAuthHandler 创建认证处理器
func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// Login 用户登录
// @Summary 用户登录
// @Description 用户登录接口
// @Tags 认证
// @Accept json
// @Produce json
// @Param request body model.LoginRequest true "登录请求"
// @Success 200 {object} model.APIResponse{data=model.LoginResponse}
// @Failure 400 {object} model.APIResponse
// @Failure 401 {object} model.APIResponse
// @Router /api/v1/auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.APIResponse{
			Code:    http.StatusBadRequest,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	// 执行登录
	response, err := h.authService.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.APIResponse{
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.APIResponse{
		Code:    http.StatusOK,
		Message: "登录成功",
		Data:    response,
	})
}
