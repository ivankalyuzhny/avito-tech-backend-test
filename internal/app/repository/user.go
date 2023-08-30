package repository

import (
	"avito-tech-backend-test/internal/app/model"
	"avito-tech-backend-test/internal/app/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	UserExists(userID int64) (bool, error)
	FindUserSegments(userID int64) ([]*model.Segment, error)
	UpdateUserSegment(userID int64, segmentIDsAdd []int64, segmentIDsDel []int64) error
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return postgres.NewUserRepositoryPostgres(db)
}
