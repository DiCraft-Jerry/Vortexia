package repository

import (
	"database/sql"
	"fmt"
	"time"

	"simple-ci/internal/model"
)

type userRepository struct {
	db *sql.DB
}

// NewUserRepository 创建用户仓库实例
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

// Create 创建用户
func (r *userRepository) Create(user *model.User) error {
	query := `
		INSERT INTO users (username, email, password_hash, role, is_active, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id`
	
	now := time.Now()
	err := r.db.QueryRow(
		query,
		user.Username,
		user.Email,
		user.Password,
		user.Role,
		user.IsActive,
		now,
		now,
	).Scan(&user.ID)
	
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	
	user.CreatedAt = now
	user.UpdatedAt = now
	return nil
}

// GetByID 根据ID获取用户
func (r *userRepository) GetByID(id int) (*model.User, error) {
	query := `
		SELECT id, username, email, password_hash, role, is_active, created_at, updated_at
		FROM users
		WHERE id = $1`
	
	user := &model.User{}
	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}
	
	return user, nil
}

// GetByUsername 根据用户名获取用户
func (r *userRepository) GetByUsername(username string) (*model.User, error) {
	query := `
		SELECT id, username, email, password_hash, role, is_active, created_at, updated_at
		FROM users
		WHERE username = $1`
	
	user := &model.User{}
	err := r.db.QueryRow(query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get user by username: %w", err)
	}
	
	return user, nil
}

// GetByEmail 根据邮箱获取用户
func (r *userRepository) GetByEmail(email string) (*model.User, error) {
	query := `
		SELECT id, username, email, password_hash, role, is_active, created_at, updated_at
		FROM users
		WHERE email = $1`
	
	user := &model.User{}
	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}
	
	return user, nil
}

// Update 更新用户
func (r *userRepository) Update(user *model.User) error {
	query := `
		UPDATE users 
		SET username = $1, email = $2, role = $3, is_active = $4, updated_at = $5
		WHERE id = $6`
	
	_, err := r.db.Exec(
		query,
		user.Username,
		user.Email,
		user.Role,
		user.IsActive,
		time.Now(),
		user.ID,
	)
	
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	
	return nil
}

// Delete 删除用户
func (r *userRepository) Delete(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	
	return nil
}

// List 获取用户列表
func (r *userRepository) List(offset, limit int) ([]*model.User, int, error) {
	// 获取总数
	var total int
	countQuery := `SELECT COUNT(*) FROM users`
	err := r.db.QueryRow(countQuery).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count users: %w", err)
	}
	
	// 获取列表
	query := `
		SELECT id, username, email, password_hash, role, is_active, created_at, updated_at
		FROM users
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2`
	
	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to list users: %w", err)
	}
	defer rows.Close()
	
	var users []*model.User
	for rows.Next() {
		user := &model.User{}
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.Password,
			&user.Role,
			&user.IsActive,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, user)
	}
	
	return users, total, nil
} 