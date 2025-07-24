package service

import (
	"simple-ci/internal/model"
	"simple-ci/internal/repository"
)

// Services 包含所有服务接口
type Services struct {
	Auth     AuthService
	User     UserService
	Project  ProjectService
	Pipeline PipelineService
	Build    BuildService
}

// NewServices 创建服务集合
func NewServices(repos *repository.Repositories) *Services {
	return &Services{
		Auth:     NewAuthService(repos.User),
		User:     NewUserService(repos.User),
		Project:  NewProjectService(repos.Project),
		Pipeline: NewPipelineService(repos.Pipeline),
		Build:    NewBuildService(repos.Build, repos.Pipeline),
	}
}

// AuthService 认证服务接口
type AuthService interface {
	Login(username, password string) (*model.LoginResponse, error)
	ValidateToken(token string) (*model.User, error)
	GenerateToken(user *model.User) (string, error)
}

// UserService 用户服务接口
type UserService interface {
	Create(req *model.CreateUserRequest) (*model.User, error)
	GetByID(id int) (*model.User, error)
	GetByUsername(username string) (*model.User, error)
	Update(user *model.User) error
	Delete(id int) error
	List(page, pageSize int) (*model.PaginationResponse, error)
}

// ProjectService 项目服务接口
type ProjectService interface {
	Create(req *model.CreateProjectRequest, ownerID int) (*model.Project, error)
	GetByID(id int) (*model.Project, error)
	GetByOwner(ownerID int) ([]*model.Project, error)
	Update(project *model.Project) error
	Delete(id int) error
	List(page, pageSize int) (*model.PaginationResponse, error)
}

// PipelineService 流水线服务接口
type PipelineService interface {
	Create(req *model.CreatePipelineRequest) (*model.Pipeline, error)
	GetByID(id int) (*model.Pipeline, error)
	GetByProject(projectID int) ([]*model.Pipeline, error)
	Update(pipeline *model.Pipeline) error
	Delete(id int) error
	List(page, pageSize int) (*model.PaginationResponse, error)
}

// BuildService 构建服务接口
type BuildService interface {
	Create(req *model.TriggerBuildRequest, triggerBy int) (*model.Build, error)
	GetByID(id int) (*model.Build, error)
	GetByPipeline(pipelineID int, page, pageSize int) (*model.PaginationResponse, error)
	UpdateStatus(id int, status string) error
	List(page, pageSize int) (*model.PaginationResponse, error)
	
	// 构建步骤相关
	GetSteps(buildID int) ([]*model.BuildStep, error)
	UpdateStepStatus(stepID int, status string, output string) error
	ExecuteBuild(buildID int) error
} 