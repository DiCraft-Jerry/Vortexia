package service

import (
	"simple-ci/internal/model"
	"simple-ci/internal/repository"
)

type pipelineService struct {
	pipelineRepo repository.PipelineRepository
}

// NewPipelineService 创建流水线服务实例
func NewPipelineService(pipelineRepo repository.PipelineRepository) PipelineService {
	return &pipelineService{pipelineRepo: pipelineRepo}
}

// Create 创建流水线
func (s *pipelineService) Create(req *model.CreatePipelineRequest) (*model.Pipeline, error) {
	// TODO: 实现流水线创建逻辑
	return nil, nil
}

// GetByID 根据ID获取流水线
func (s *pipelineService) GetByID(id int) (*model.Pipeline, error) {
	return s.pipelineRepo.GetByID(id)
}

// GetByProject 根据项目获取流水线列表
func (s *pipelineService) GetByProject(projectID int) ([]*model.Pipeline, error) {
	return s.pipelineRepo.GetByProject(projectID)
}

// Update 更新流水线
func (s *pipelineService) Update(pipeline *model.Pipeline) error {
	return s.pipelineRepo.Update(pipeline)
}

// Delete 删除流水线
func (s *pipelineService) Delete(id int) error {
	return s.pipelineRepo.Delete(id)
}

// List 获取流水线列表
func (s *pipelineService) List(page, pageSize int) (*model.PaginationResponse, error) {
	// TODO: 实现分页逻辑
	return nil, nil
} 