package repository

import (
	"database/sql"

	"simple-ci/internal/model"

	"github.com/go-redis/redis/v8"
)

// Repositories 包含所有仓库接口
type Repositories struct {
	User     UserRepository
	Project  ProjectRepository
	Pipeline PipelineRepository
	Build    BuildRepository
}

// NewRepositories 创建仓库集合
func NewRepositories(db *sql.DB, redis *redis.Client) *Repositories {
	return &Repositories{
		User:     NewUserRepository(db),
		Project:  NewProjectRepository(db),
		Pipeline: NewPipelineRepository(db),
		Build:    NewBuildRepository(db, redis),
	}
}

// UserRepository 用户仓库接口
type UserRepository interface {
	Create(user *model.User) error
	GetByID(id int) (*model.User, error)
	GetByUsername(username string) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	Update(user *model.User) error
	Delete(id int) error
	List(offset, limit int) ([]*model.User, int, error)
}

// ProjectRepository 项目仓库接口
type ProjectRepository interface {
	Create(project *model.Project) error
	GetByID(id int) (*model.Project, error)
	GetByOwner(ownerID int) ([]*model.Project, error)
	Update(project *model.Project) error
	Delete(id int) error
	List(offset, limit int) ([]*model.Project, int, error)
}

// PipelineRepository 流水线仓库接口
type PipelineRepository interface {
	Create(pipeline *model.Pipeline) error
	GetByID(id int) (*model.Pipeline, error)
	GetByProject(projectID int) ([]*model.Pipeline, error)
	Update(pipeline *model.Pipeline) error
	Delete(id int) error
	List(offset, limit int) ([]*model.Pipeline, int, error)
}

// BuildRepository 构建仓库接口
type BuildRepository interface {
	Create(build *model.Build) error
	GetByID(id int) (*model.Build, error)
	GetByPipeline(pipelineID int, offset, limit int) ([]*model.Build, int, error)
	UpdateStatus(id int, status string) error
	List(offset, limit int) ([]*model.Build, int, error)
	
	// 构建步骤相关
	CreateStep(step *model.BuildStep) error
	GetStepsByBuild(buildID int) ([]*model.BuildStep, error)
	UpdateStepStatus(id int, status string, output string) error
} 