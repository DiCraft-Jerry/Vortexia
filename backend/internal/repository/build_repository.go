package repository

import (
	"database/sql"
	"fmt"
	"time"

	"simple-ci/internal/model"

	"github.com/go-redis/redis/v8"
)

type buildRepository struct {
	db    *sql.DB
	redis *redis.Client
}

// NewBuildRepository 创建构建仓库实例
func NewBuildRepository(db *sql.DB, redis *redis.Client) BuildRepository {
	return &buildRepository{db: db, redis: redis}
}

// Create 创建构建
func (r *buildRepository) Create(build *model.Build) error {
	query := `
		INSERT INTO builds (pipeline_id, branch, commit, status, started_at, trigger_by, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id`
	
	now := time.Now()
	err := r.db.QueryRow(
		query,
		build.PipelineID,
		build.Branch,
		build.Commit,
		build.Status,
		build.StartedAt,
		build.TriggerBy,
		now,
	).Scan(&build.ID)
	
	if err != nil {
		return fmt.Errorf("failed to create build: %w", err)
	}
	
	build.CreatedAt = now
	return nil
}

// GetByID 根据ID获取构建
func (r *buildRepository) GetByID(id int) (*model.Build, error) {
	query := `
		SELECT id, pipeline_id, branch, commit, status, started_at, finished_at, duration, trigger_by, created_at
		FROM builds
		WHERE id = $1`
	
	build := &model.Build{}
	err := r.db.QueryRow(query, id).Scan(
		&build.ID,
		&build.PipelineID,
		&build.Branch,
		&build.Commit,
		&build.Status,
		&build.StartedAt,
		&build.FinishedAt,
		&build.Duration,
		&build.TriggerBy,
		&build.CreatedAt,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get build by id: %w", err)
	}
	
	return build, nil
}

// GetByPipeline 根据流水线获取构建列表
func (r *buildRepository) GetByPipeline(pipelineID int, offset, limit int) ([]*model.Build, int, error) {
	// 获取总数
	var total int
	countQuery := `SELECT COUNT(*) FROM builds WHERE pipeline_id = $1`
	err := r.db.QueryRow(countQuery, pipelineID).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count builds: %w", err)
	}
	
	// 获取列表
	query := `
		SELECT id, pipeline_id, branch, commit, status, started_at, finished_at, duration, trigger_by, created_at
		FROM builds
		WHERE pipeline_id = $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3`
	
	rows, err := r.db.Query(query, pipelineID, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get builds by pipeline: %w", err)
	}
	defer rows.Close()
	
	var builds []*model.Build
	for rows.Next() {
		build := &model.Build{}
		err := rows.Scan(
			&build.ID,
			&build.PipelineID,
			&build.Branch,
			&build.Commit,
			&build.Status,
			&build.StartedAt,
			&build.FinishedAt,
			&build.Duration,
			&build.TriggerBy,
			&build.CreatedAt,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to scan build: %w", err)
		}
		builds = append(builds, build)
	}
	
	return builds, total, nil
}

// UpdateStatus 更新构建状态
func (r *buildRepository) UpdateStatus(id int, status string) error {
	var query string
	var args []interface{}
	
	if status == model.BuildStatusSuccess || status == model.BuildStatusFailed || status == model.BuildStatusCanceled {
		// 完成状态，更新结束时间和持续时间
		query = `
			UPDATE builds 
			SET status = $1, finished_at = $2, 
			    duration = EXTRACT(EPOCH FROM ($2 - started_at))::int
			WHERE id = $3`
		args = []interface{}{status, time.Now(), id}
	} else {
		// 其他状态，只更新状态
		query = `UPDATE builds SET status = $1 WHERE id = $2`
		args = []interface{}{status, id}
	}
	
	_, err := r.db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("failed to update build status: %w", err)
	}
	
	return nil
}

// List 获取构建列表
func (r *buildRepository) List(offset, limit int) ([]*model.Build, int, error) {
	// 获取总数
	var total int
	countQuery := `SELECT COUNT(*) FROM builds`
	err := r.db.QueryRow(countQuery).Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count builds: %w", err)
	}
	
	// 获取列表
	query := `
		SELECT id, pipeline_id, branch, commit, status, started_at, finished_at, duration, trigger_by, created_at
		FROM builds
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2`
	
	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to list builds: %w", err)
	}
	defer rows.Close()
	
	var builds []*model.Build
	for rows.Next() {
		build := &model.Build{}
		err := rows.Scan(
			&build.ID,
			&build.PipelineID,
			&build.Branch,
			&build.Commit,
			&build.Status,
			&build.StartedAt,
			&build.FinishedAt,
			&build.Duration,
			&build.TriggerBy,
			&build.CreatedAt,
		)
		if err != nil {
			return nil, 0, fmt.Errorf("failed to scan build: %w", err)
		}
		builds = append(builds, build)
	}
	
	return builds, total, nil
}

// CreateStep 创建构建步骤
func (r *buildRepository) CreateStep(step *model.BuildStep) error {
	query := `
		INSERT INTO build_steps (build_id, name, command, status, output, started_at, step_order)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id`
	
	err := r.db.QueryRow(
		query,
		step.BuildID,
		step.Name,
		step.Command,
		step.Status,
		step.Output,
		step.StartedAt,
		step.StepOrder,
	).Scan(&step.ID)
	
	if err != nil {
		return fmt.Errorf("failed to create build step: %w", err)
	}
	
	return nil
}

// GetStepsByBuild 根据构建ID获取步骤列表
func (r *buildRepository) GetStepsByBuild(buildID int) ([]*model.BuildStep, error) {
	query := `
		SELECT id, build_id, name, command, status, output, started_at, finished_at, duration, step_order
		FROM build_steps
		WHERE build_id = $1
		ORDER BY step_order ASC`
	
	rows, err := r.db.Query(query, buildID)
	if err != nil {
		return nil, fmt.Errorf("failed to get build steps: %w", err)
	}
	defer rows.Close()
	
	var steps []*model.BuildStep
	for rows.Next() {
		step := &model.BuildStep{}
		err := rows.Scan(
			&step.ID,
			&step.BuildID,
			&step.Name,
			&step.Command,
			&step.Status,
			&step.Output,
			&step.StartedAt,
			&step.FinishedAt,
			&step.Duration,
			&step.StepOrder,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan build step: %w", err)
		}
		steps = append(steps, step)
	}
	
	return steps, nil
}

// UpdateStepStatus 更新步骤状态
func (r *buildRepository) UpdateStepStatus(id int, status string, output string) error {
	var query string
	var args []interface{}
	
	if status == model.StepStatusSuccess || status == model.StepStatusFailed || status == model.StepStatusSkipped {
		// 完成状态，更新结束时间和持续时间
		query = `
			UPDATE build_steps 
			SET status = $1, output = $2, finished_at = $3, 
			    duration = EXTRACT(EPOCH FROM ($3 - started_at))::int
			WHERE id = $4`
		args = []interface{}{status, output, time.Now(), id}
	} else {
		// 其他状态，只更新状态和输出
		query = `UPDATE build_steps SET status = $1, output = $2 WHERE id = $3`
		args = []interface{}{status, output, id}
	}
	
	_, err := r.db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("failed to update step status: %w", err)
	}
	
	return nil
} 