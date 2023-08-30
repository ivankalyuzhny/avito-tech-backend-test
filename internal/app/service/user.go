package service

import (
	"avito-tech-backend-test/internal/app/model"
	"avito-tech-backend-test/internal/app/repository"
	"avito-tech-backend-test/pkg/utils"
	"fmt"
)

type UserService struct {
	userRepo    *repository.UserRepository
	segmentRepo *repository.SegmentRepository
}

func NewUserService(userRepo *repository.UserRepository, segmentRepo *repository.SegmentRepository) *UserService {
	return &UserService{
		userRepo:    userRepo,
		segmentRepo: segmentRepo,
	}
}

func (s *UserService) UpdateUserSegments(userID int64, segmentSlugsAdd []string, segmentSlugsDel []string) error {
	exists, err := s.userRepo.Exists(userID)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("user (id=%d) not found", userID)
	}
	segmentSlugsAddUnique, segmentSlugsDelUnique := utils.RemoveIntersections(
		utils.RemoveRepeats(segmentSlugsAdd),
		utils.RemoveRepeats(segmentSlugsDel),
	)
	segmentIDsAdd := make([]int64, len(segmentSlugsAddUnique))
	for i, segmentSlug := range segmentSlugsAddUnique {
		segment, err := s.segmentRepo.FindBySlug(segmentSlug)
		if err != nil {
			return err
		}
		segmentIDsAdd[i] = segment.ID
	}
	segmentIDsDel := make([]int64, len(segmentSlugsDelUnique))
	for i, segmentSlug := range segmentSlugsDelUnique {
		segment, err := s.segmentRepo.FindBySlug(segmentSlug)
		if err != nil {
			return err
		}
		segmentIDsDel[i] = segment.ID
	}
	err = s.userRepo.UpdateUserSegment(userID, segmentIDsAdd, segmentIDsDel)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) AddToSegments(userID int64, segmentSlugs []string) error {
	exists, err := s.userRepo.Exists(userID)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("user (id=%d) not found", userID)
	}
	for _, slug := range segmentSlugs {
		segment, err := s.segmentRepo.FindBySlug(slug)
		if err != nil {
			return err
		}
		err = s.userRepo.AddToSegment(userID, segment.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *UserService) RemoveFromSegments(userID int64, segmentSlugs []string) error {
	exists, err := s.userRepo.Exists(userID)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("user (id=%d) not found", userID)
	}
	for _, slug := range segmentSlugs {
		segment, err := s.segmentRepo.FindBySlug(slug)
		if err != nil {
			return err
		}
		err = s.userRepo.RemoveFromSegment(userID, segment.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *UserService) FindUserSegments(userID int64) ([]*model.Segment, error) {
	exists, err := s.userRepo.Exists(userID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("user (id=%d) not found", userID)
	}
	return s.userRepo.FindUserSegments(userID)
}
