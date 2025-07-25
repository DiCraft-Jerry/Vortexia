package service

import (
	"errors"
	"math"

	"Vortexia/internal/model"
	"Vortexia/internal/repository"
)

type userService struct {
	userRepo repository.UserRepository
}

// NewUserService 创建用户服务实例
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

// Create 创建用户
func (s *userService) Create(req *model.CreateUserRequest) (*model.User, error) {
	// 检查用户名是否已存在
	existingUser, err := s.userRepo.GetByUsername(req.Username)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("用户名已存在")
	}

	// 检查邮箱是否已存在
	existingUser, err = s.userRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("邮箱已存在")
	}

	// 哈希密码
	hashedPassword, err := HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// 创建用户
	user := &model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
		Role:     req.Role,
		IsActive: true,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

// GetByID 根据ID获取用户
func (s *userService) GetByID(id int) (*model.User, error) {
	return s.userRepo.GetByID(id)
}

// GetByUsername 根据用户名获取用户
func (s *userService) GetByUsername(username string) (*model.User, error) {
	return s.userRepo.GetByUsername(username)
}

// Update 更新用户
func (s *userService) Update(user *model.User) error {
	// 检查用户是否存在
	existingUser, err := s.userRepo.GetByID(user.ID)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return errors.New("用户不存在")
	}

	// 检查用户名是否被其他用户使用
	if existingUser.Username != user.Username {
		userByUsername, err := s.userRepo.GetByUsername(user.Username)
		if err != nil {
			return err
		}
		if userByUsername != nil && userByUsername.ID != user.ID {
			return errors.New("用户名已被使用")
		}
	}

	// 检查邮箱是否被其他用户使用
	if existingUser.Email != user.Email {
		userByEmail, err := s.userRepo.GetByEmail(user.Email)
		if err != nil {
			return err
		}
		if userByEmail != nil && userByEmail.ID != user.ID {
			return errors.New("邮箱已被使用")
		}
	}

	return s.userRepo.Update(user)
}

// Delete 删除用户
func (s *userService) Delete(id int) error {
	// 检查用户是否存在
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("用户不存在")
	}

	return s.userRepo.Delete(id)
}

// List 获取用户列表
func (s *userService) List(page, pageSize int) (*model.PaginationResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	offset := (page - 1) * pageSize
	users, total, err := s.userRepo.List(offset, pageSize)
	if err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))

	return &model.PaginationResponse{
		Items:      users,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}, nil
}
