package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"avito-tech-backend-test/internal/app/model"

	"github.com/jmoiron/sqlx"
)

type UserRepositoryPostgres struct {
	db *sqlx.DB
}

func NewUserRepositoryPostgres(db *sqlx.DB) *UserRepositoryPostgres {
	return &UserRepositoryPostgres{
		db: db,
	}
}

func (r *UserRepositoryPostgres) UserExists(userID int64) (bool, error) {
	var tmp int64
	query := `SELECT 1 FROM users u WHERE u.id = $1 LIMIT 1`
	err := r.db.Get(&tmp, query, userID)
	switch {
	case err == nil:
		return true, nil
	case errors.Is(err, sql.ErrNoRows):
		return false, nil
	default:
		return false, fmt.Errorf("failed to check user exists (id=%d): %w", userID, err)
	}
}

func (r *UserRepositoryPostgres) AddToSegment(userID int64, segmentID int64) error {
	query := `INSERT INTO user_segments (user_id, segment_id) VALUES ($1, $2)`
	_, err := r.db.Exec(query, userID, segmentID)
	if err != nil {
		return fmt.Errorf("failed to add user (id=%d) to segment (id=%d): %w", userID, segmentID, err)
	}
	log.Printf("user (id=%d) added to segment (id=%d)", userID, segmentID)
	return nil
}

func (r *UserRepositoryPostgres) RemoveFromSegment(userID int64, segmentID int64) error {
	query := `DELETE FROM user_segments us WHERE us.user_id = $1 AND us.segment_id = $2`
	_, err := r.db.Exec(query, userID, segmentID)
	if err != nil {
		return fmt.Errorf("failed to delete user (id=%d) from segment (id=%d): %w", userID, segmentID, err)
	}
	log.Printf("user (id=%d) deleted from segment (id=%d)", userID, segmentID)
	return nil
}

func (r *UserRepositoryPostgres) FindUserSegments(userID int64) ([]*model.Segment, error) {
	var segments []*model.Segment
	query := `SELECT s.id, s.slug FROM segments s JOIN user_segments us ON s.id = us.segment_id WHERE us.user_id = $1`
	err := r.db.Select(&segments, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to select segments by user (id=%d): %w", userID, err)
	}
	log.Printf("found segments by user (id=%d)", userID)
	return segments, nil
}

func (r *UserRepositoryPostgres) UpdateUserSegment(userID int64, segmentIDsAdd []int64, segmentIDsDel []int64) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}
	if len(segmentIDsAdd) == 0 && len(segmentIDsDel) == 0 {
		return nil
	}
	queryAdd := `INSERT INTO user_segments (user_id, segment_id) VALUES ($1, $2)`
	for _, segmentID := range segmentIDsAdd {
		_, err = tx.Exec(queryAdd, userID, segmentID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to add user (id=%d) to segment (id=%d): %w", userID, segmentID, err)
		}
		log.Printf("user (id=%d) added to segment (id=%d)", userID, segmentID)
	}
	queryDel := `DELETE FROM user_segments us WHERE us.user_id = $1 AND us.segment_id = $2`
	for _, segmentID := range segmentIDsDel {
		result, err := tx.Exec(queryDel, userID, segmentID)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to delete user (id=%d) from segment (id=%d): %w", userID, segmentID, err)
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to delete user (id=%d) from segment (id=%d): %w", userID, segmentID, err)
		}
		if rowsAffected == 0 {
			tx.Rollback()
			return fmt.Errorf("failed to delete user (id=%d) from segment (id=%d): user is not in segment", userID, segmentID)
		}
		log.Printf("user (id=%d) deleted from segment (id=%d)", userID, segmentID)
	}
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	return nil
}
