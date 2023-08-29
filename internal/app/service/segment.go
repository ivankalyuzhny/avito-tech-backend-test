package service

import (
	"avito-tech-backend-test/internal/app/model"
	"avito-tech-backend-test/internal/app/repository"
)

type SegmentService struct {
	repo *repository.SegmentRepository
}

func NewSegmentService(repo *repository.SegmentRepository) *SegmentService {
	return &SegmentService{
		repo: repo,
	}
}

func (s *SegmentService) Create(slug string) (*model.Segment, error) {
	segment := model.NewSegment(0, slug)
	insertID, err := s.repo.Create(segment)
	if err != nil {
		return nil, err
	}
	segment.ID = insertID
	return segment, nil
}

func (s *SegmentService) Delete(slug string) error {
	err := s.repo.Delete(slug)
	return err
}
