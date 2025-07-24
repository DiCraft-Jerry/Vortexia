package repository

import (
	"database/sql"
	"fmt"
	"time"

	"simple-ci/internal/model"
)

type projectRepository struct {
	db *sql.DB
}

// NewProjectRepository 创建项目仓库实例
func NewProjectRepository(db *sql.DB) ProjectRepository {
	return &projectRepository{db: db}
}

// Create 创建项目
func (r *projectRepository) Create(project *model.Project) error {
	query := `
		INSERT INTO projects (name, description, repo_url, branch, owner_id, is_active, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id`
	
	now := time.Now()
	err := r.db.QueryRow(
		query,
		project.Name,
		project.Description,
		project.RepoURL,
		project.Branch,
		project.OwnerID,
		project.IsActive,
		now,
		now,
	).Scan(&project.ID)
	
	if err != nil {
		return fmt.Errorf("failed to create project: %w", err)
	}
	
	project.CreatedAt = now
	project.UpdatedAt = now
	return nil
}

// GetByID 根据ID获取项目
func (r *projectRepository) GetByID(id int) (*model.Project, error) {
	query := `
		SELECT id, name, description, repo_url, branch, owner_id, is_active, created_at, updated_at
		FROM projects
		WHERE id = $1`
	
	project := &model.Project{}
	err := r.db.QueryRow(query, id).Scan(
		&project.ID,
		&project.Name,
		&project.Description,
		&project.RepoURL,
		&project.Branch,
		&project.OwnerID,
		&project.IsActive,
		&project.CreatedAt,
		&project.UpdatedAt,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get project by id: %w", err)
	}
	
	return project, nil
}

// GetByOwner 根据所有者获取项目列表
func (r *projectRepository) GetByOwner(ownerID int) ([]*model.Project, error) {
	query := `
		SELECT id, name, description, repo_url, branch, owner_id, is_active, created_at, updated_at
		FROM projects
		WHERE owner_id = $1 AND is_active = true
		ORDER BY created_at DESC`
	
	rows, err := r.db.Query(query, ownerID)
	if err != nil {
		return nil, fmt.Errorf("failed to get projects by owner: %w", err)
	}
	defer rows.Close()
	
	var projects []*model.Project
	for rows.Next() {
		project := &model.Project{}
		err := rows.Scan(
			&project.ID,
			&project.Name,
			&project.Description,
			&project.RepoURL,
			&project.Branch,
			&project.OwnerID,
			&project.IsActive,
			&project.CreatedAt,
			&project.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan project: %w", err)
		}
		projects = append(projects, project)
	}
	
	return projects, nil
}

// Update 更新项目
func (r *projectRepository) Update(project *model.Project) error {
	query := `
		UPDATE projects 
		SET name = $1, description = $2, repo_url = $3, branch = $4, is_active = $5, updated_at = $6
		WHERE id = $7`
	
	_, err := r.db.Exec(
		query,
		project.Name,
		project.Description,
		project.RepoURL,
		project.Branch,
		project.IsActive,
		time.Now(),
		project.ID,
	)
	
	if err != nil {
		return fmt.Errorf("failed to update project: %w", err)
	}
	
	return nil
}

// Delete 删除项目
func (r *projectRepository) Delete(id int) error {
	query := `UPDATE projects SET is_active = false WHERE id = $1`
	
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete project: %w", err)
	}
	
	return nil
}

// List 获取项目列表
func (r *projectRepository) List(offset, limit int) ([]*model.Project, int, error) {
	// 获取总数
	var total int
	countQuery := `SELECT COUNT(*) FROM projects WHERE is_active = true`
	err := r.db.QueryRow(countQuery).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count projects: %w", err)
	}
	
	// 获取列表
	query := `
		SELECT id, name, description, repo_url, branch, owner_id, is_active, created_at, updated_at
		FROM projects
		WHERE is_active = true
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2`
	
	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to list projects: %w", err)
	}
	defer rows.Close()
	
	var projects []*model.Project
	for rows.Next() {
		project := &model.Project{}
		err := rows.Scan(
			&project.ID,
			&project.Name,
			&project.Description,
			&project.RepoURL,
			&project.Branch,
			&project.OwnerID,
			&project.IsActive,
			&project.CreatedAt,
			&project.UpdatedAt,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to scan project: %w", err)
		}
		projects = append(projects, project)
	}
	
	return projects, total, nil
} 