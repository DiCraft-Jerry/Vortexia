package service

import (
	"Vortexia/internal/model"
	"Vortexia/internal/repository"
)

type buildService struct {
	buildRepo    repository.BuildRepository
	pipelineRepo repository.PipelineRepository
}

// NewBuildService 创建构建服务实例
func NewBuildService(buildRepo repository.BuildRepository, pipelineRepo repository.PipelineRepository) BuildService {
	return &buildService{
		buildRepo:    buildRepo,
		pipelineRepo: pipelineRepo,
	}
}

// Create 创建构建
func (s *buildService) Create(req *model.TriggerBuildRequest, triggerBy int) (*model.Build, error) {
	// TODO: 实现构建创建逻辑
	return nil, nil
}

// GetByID 根据ID获取构建
func (s *buildService) GetByID(id int) (*model.Build, error) {
	return s.buildRepo.GetByID(id)
}

// GetByPipeline 根据流水线获取构建列表
func (s *buildService) GetByPipeline(pipelineID int, page, pageSize int) (*model.PaginationResponse, error) {
	// TODO: 实现分页逻辑
	return nil, nil
}

// UpdateStatus 更新构建状态
func (s *buildService) UpdateStatus(id int, status string) error {
	return s.buildRepo.UpdateStatus(id, status)
}

// List 获取构建列表
func (s *buildService) List(page, pageSize int) (*model.PaginationResponse, error) {
	// TODO: 实现分页逻辑
	return nil, nil
}

// GetSteps 获取构建步骤
func (s *buildService) GetSteps(buildID int) ([]*model.BuildStep, error) {
	return s.buildRepo.GetStepsByBuild(buildID)
}

// UpdateStepStatus 更新步骤状态
func (s *buildService) UpdateStepStatus(stepID int, status string, output string) error {
	return s.buildRepo.UpdateStepStatus(stepID, status, output)
}

// ExecuteBuild 执行构建
func (s *buildService) ExecuteBuild(buildID int) error {
	// TODO: 实现构建执行逻辑
	return nil
}
