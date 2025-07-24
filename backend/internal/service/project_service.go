package service

import (
	"errors"
	"math"

	"simple-ci/internal/model"
	"simple-ci/internal/repository"
)

type projectService struct {
	projectRepo repository.ProjectRepository
}

// NewProjectService 创建项目服务实例
func NewProjectService(projectRepo repository.ProjectRepository) ProjectService {
	return &projectService{projectRepo: projectRepo}
}

// Create 创建项目
func (s *projectService) Create(req *model.CreateProjectRequest, ownerID int) (*model.Project, error) {
	// TODO: 可以添加项目名称重复检查等业务逻辑

	project := &model.Project{
		Name:        req.Name,
		Description: req.Description,
		RepoURL:     req.RepoURL,
		Branch:      req.Branch,
		OwnerID:     ownerID,
		IsActive:    true,
	}

	if err := s.projectRepo.Create(project); err != nil {
		return nil, err
	}

	return project, nil
}

// GetByID 根据ID获取项目
func (s *projectService) GetByID(id int) (*model.Project, error) {
	return s.projectRepo.GetByID(id)
}

// GetByOwner 根据所有者获取项目列表
func (s *projectService) GetByOwner(ownerID int) ([]*model.Project, error) {
	return s.projectRepo.GetByOwner(ownerID)
}

// Update 更新项目
func (s *projectService) Update(project *model.Project) error {
	// 检查项目是否存在
	existingProject, err := s.projectRepo.GetByID(project.ID)
	if err != nil {
		return err
	}
	if existingProject == nil {
		return errors.New("项目不存在")
	}

	return s.projectRepo.Update(project)
}

// Delete 删除项目
func (s *projectService) Delete(id int) error {
	// 检查项目是否存在
	project, err := s.projectRepo.GetByID(id)
	if err != nil {
		return err
	}
	if project == nil {
		return errors.New("项目不存在")
	}

	return s.projectRepo.Delete(id)
}

// List 获取项目列表
func (s *projectService) List(page, pageSize int) (*model.PaginationResponse, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	offset := (page - 1) * pageSize
	projects, total, err := s.projectRepo.List(offset, pageSize)
	if err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))

	return &model.PaginationResponse{
		Items:      projects,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}, nil
} 