package repository

import (
	"database/sql"
	"fmt"
	"time"

	"Vortexia/internal/model"
)

type pipelineRepository struct {
	db *sql.DB
}

// NewPipelineRepository 创建流水线仓库实例
func NewPipelineRepository(db *sql.DB) PipelineRepository {
	return &pipelineRepository{db: db}
}

// Create 创建流水线
func (r *pipelineRepository) Create(pipeline *model.Pipeline) error {
	query := `
		INSERT INTO pipelines (project_id, name, config, is_active, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`

	now := time.Now()
	err := r.db.QueryRow(
		query,
		pipeline.ProjectID,
		pipeline.Name,
		pipeline.Config,
		pipeline.IsActive,
		now,
		now,
	).Scan(&pipeline.ID)

	if err != nil {
		return fmt.Errorf("failed to create pipeline: %w", err)
	}

	pipeline.CreatedAt = now
	pipeline.UpdatedAt = now
	return nil
}

// GetByID 根据ID获取流水线
func (r *pipelineRepository) GetByID(id int) (*model.Pipeline, error) {
	query := `
		SELECT id, project_id, name, config, is_active, created_at, updated_at
		FROM pipelines
		WHERE id = $1`

	pipeline := &model.Pipeline{}
	err := r.db.QueryRow(query, id).Scan(
		&pipeline.ID,
		&pipeline.ProjectID,
		&pipeline.Name,
		&pipeline.Config,
		&pipeline.IsActive,
		&pipeline.CreatedAt,
		&pipeline.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get pipeline by id: %w", err)
	}

	return pipeline, nil
}

// GetByProject 根据项目获取流水线列表
func (r *pipelineRepository) GetByProject(projectID int) ([]*model.Pipeline, error) {
	query := `
		SELECT id, project_id, name, config, is_active, created_at, updated_at
		FROM pipelines
		WHERE project_id = $1 AND is_active = true
		ORDER BY created_at DESC`

	rows, err := r.db.Query(query, projectID)
	if err != nil {
		return nil, fmt.Errorf("failed to get pipelines by project: %w", err)
	}
	defer rows.Close()

	var pipelines []*model.Pipeline
	for rows.Next() {
		pipeline := &model.Pipeline{}
		err := rows.Scan(
			&pipeline.ID,
			&pipeline.ProjectID,
			&pipeline.Name,
			&pipeline.Config,
			&pipeline.IsActive,
			&pipeline.CreatedAt,
			&pipeline.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan pipeline: %w", err)
		}
		pipelines = append(pipelines, pipeline)
	}

	return pipelines, nil
}

// Update 更新流水线
func (r *pipelineRepository) Update(pipeline *model.Pipeline) error {
	query := `
		UPDATE pipelines 
		SET name = $1, config = $2, is_active = $3, updated_at = $4
		WHERE id = $5`

	_, err := r.db.Exec(
		query,
		pipeline.Name,
		pipeline.Config,
		pipeline.IsActive,
		time.Now(),
		pipeline.ID,
	)

	if err != nil {
		return fmt.Errorf("failed to update pipeline: %w", err)
	}

	return nil
}

// Delete 删除流水线
func (r *pipelineRepository) Delete(id int) error {
	query := `UPDATE pipelines SET is_active = false WHERE id = $1`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete pipeline: %w", err)
	}

	return nil
}

// List 获取流水线列表
func (r *pipelineRepository) List(offset, limit int) ([]*model.Pipeline, int, error) {
	// 获取总数
	var total int
	countQuery := `SELECT COUNT(*) FROM pipelines WHERE is_active = true`
	err := r.db.QueryRow(countQuery).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count pipelines: %w", err)
	}

	// 获取列表
	query := `
		SELECT id, project_id, name, config, is_active, created_at, updated_at
		FROM pipelines
		WHERE is_active = true
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2`

	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to list pipelines: %w", err)
	}
	defer rows.Close()

	var pipelines []*model.Pipeline
	for rows.Next() {
		pipeline := &model.Pipeline{}
		err := rows.Scan(
			&pipeline.ID,
			&pipeline.ProjectID,
			&pipeline.Name,
			&pipeline.Config,
			&pipeline.IsActive,
			&pipeline.CreatedAt,
			&pipeline.UpdatedAt,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to scan pipeline: %w", err)
		}
		pipelines = append(pipelines, pipeline)
	}

	return pipelines, total, nil
}
