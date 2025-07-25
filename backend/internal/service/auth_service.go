package service

import (
	"errors"
	"time"

	"Vortexia/internal/config"
	"Vortexia/internal/model"
	"Vortexia/internal/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	userRepo repository.UserRepository
}

// NewAuthService 创建认证服务实例
func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepo: userRepo}
}

// Login 用户登录
func (s *authService) Login(username, password string) (*model.LoginResponse, error) {
	// 根据用户名获取用户
	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("用户名或密码错误")
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("用户名或密码错误")
	}

	// 检查用户状态
	if !user.IsActive {
		return nil, errors.New("用户已被禁用")
	}

	// 生成JWT token
	token, err := s.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	return &model.LoginResponse{
		Token: token,
		User:  *user,
	}, nil
}

// ValidateToken 验证JWT token
func (s *authService) ValidateToken(tokenString string) (*model.User, error) {
	// 解析token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("无效的签名方法")
		}
		cfg, _ := config.Load()
		return []byte(cfg.JWT.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	// 验证claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := int(claims["user_id"].(float64))

		// 从数据库获取用户信息
		user, err := s.userRepo.GetByID(userID)
		if err != nil {
			return nil, err
		}
		if user == nil {
			return nil, errors.New("用户不存在")
		}
		if !user.IsActive {
			return nil, errors.New("用户已被禁用")
		}

		return user, nil
	}

	return nil, errors.New("无效的token")
}

// GenerateToken 生成JWT token
func (s *authService) GenerateToken(user *model.User) (string, error) {
	cfg, err := config.Load()
	if err != nil {
		return "", err
	}

	// 创建claims
	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"email":    user.Email,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Duration(cfg.JWT.Expire) * time.Second).Unix(),
		"iat":      time.Now().Unix(),
	}

	// 生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWT.Secret))
}

// HashPassword 密码哈希
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
