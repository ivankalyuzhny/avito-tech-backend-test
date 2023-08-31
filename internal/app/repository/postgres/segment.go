package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"

	"avito-tech-backend-test/internal/app/model"
)

type SegmentRepositoryPostgres struct {
	db *sqlx.DB
}

func NewSegmentRepositoryPostgres(db *sqlx.DB) *SegmentRepositoryPostgres {
	return &SegmentRepositoryPostgres{
		db: db,
	}
}

func (r *SegmentRepositoryPostgres) CreateSegment(segment *model.Segment) error {
	query := `INSERT INTO segments (slug) VALUES ($1) RETURNING id`
	var insertID int64
	err := r.db.QueryRow(query, segment.Slug).Scan(&insertID)
	if err != nil {
		return fmt.Errorf("failed to create segment (slug=%s):%w", segment.Slug, err)
	}
	log.Printf("created segment (id=%d slug=%s)", insertID, segment.Slug)
	return nil
}

func (r *SegmentRepositoryPostgres) DeleteSegment(slug string) error {
	query := `DELETE FROM segments s WHERE s.slug = $1`
	_, err := r.db.Exec(query, slug)
	if err != nil {
		return fmt.Errorf("failed to delete segment (slug=%s): %w", slug, err)
	}
	return nil
}

func (r *SegmentRepositoryPostgres) GetSegment(slug string) (*model.Segment, error) {
	var segment model.Segment
	query := `SELECT id, slug FROM segments s WHERE s.slug = $1`
	err := r.db.Get(&segment, query, slug)
	if err != nil {
		return nil, fmt.Errorf("failed to select segment by (slug=%s): %w", slug, err)
	}
	return &segment, nil
}
