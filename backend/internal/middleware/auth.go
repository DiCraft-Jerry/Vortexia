package middleware

import (
	"net/http"
	"strings"

	"Vortexia/internal/model"
	"Vortexia/internal/service"

	"github.com/gin-gonic/gin"
)

// JWTAuth JWT认证中间件
func JWTAuth(authService service.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头获取token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, model.APIResponse{
				Code:    http.StatusUnauthorized,
				Message: "未提供认证令牌",
			})
			c.Abort()
			return
		}

		// 检查Bearer格式
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, model.APIResponse{
				Code:    http.StatusUnauthorized,
				Message: "认证令牌格式错误",
			})
			c.Abort()
			return
		}

		// 验证token
		user, err := authService.ValidateToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, model.APIResponse{
				Code:    http.StatusUnauthorized,
				Message: "认证令牌无效: " + err.Error(),
			})
			c.Abort()
			return
		}

		// 将用户信息存储到上下文
		c.Set("user", user)
		c.Set("user_id", user.ID)
		c.Set("user_role", user.Role)

		c.Next()
	}
}

// AdminRequired 管理员权限中间件
func AdminRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, model.APIResponse{
				Code:    http.StatusUnauthorized,
				Message: "用户信息不存在",
			})
			c.Abort()
			return
		}

		u, ok := user.(*model.User)
		if !ok || u.Role != model.RoleAdmin {
			c.JSON(http.StatusForbidden, model.APIResponse{
				Code:    http.StatusForbidden,
				Message: "需要管理员权限",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// GetCurrentUser 从上下文获取当前用户
func GetCurrentUser(c *gin.Context) (*model.User, bool) {
	user, exists := c.Get("user")
	if !exists {
		return nil, false
	}

	u, ok := user.(*model.User)
	return u, ok
}

// GetCurrentUserID 从上下文获取当前用户ID
func GetCurrentUserID(c *gin.Context) (int, bool) {
	userID, exists := c.Get("user_id")
	if !exists {
		return 0, false
	}

	id, ok := userID.(int)
	return id, ok
}
