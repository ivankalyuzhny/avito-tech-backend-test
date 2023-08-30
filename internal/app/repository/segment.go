package repository

import (
	"avito-tech-backend-test/internal/app/model"
	"avito-tech-backend-test/internal/app/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type SegmentRepository interface {
	CreateSegment(segment *model.Segment) error
	DeleteSegment(slug string) error
	GetSegment(slug string) (*model.Segment, error)
}

func NewSegmentRepository(db *sqlx.DB) SegmentRepository {
	return postgres.NewSegmentRepositoryPostgres(db)
}
