package repository

import (
	"context"
	"errors"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/gofrs/uuid/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rizwank123/myResturent/internal/domain"
)

type pgxRatingRepository struct {
	db  *pgxpool.Pool
	sqt sq.StatementBuilderType
}

func NewRatingRepository(db *pgxpool.Pool) domain.RatingRepository {
	return &pgxRatingRepository{
		db:  db,
		sqt: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

// CreateRating implements domain.RatingRepository.
func (r *pgxRatingRepository) CreateRating(ctx context.Context, in *domain.Rating) (err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	txVal := ctx.Value(TxKey)
	q := `INSERT INTO rating (name, rating, resturent_id, review, suggestion) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at, updated_at`
	args := []interface{}{in.Name, in.Rating, in.ResturentID, in.Review, in.Suggestion}
	if txVal != nil {
		tx := txVal.(pgx.Tx)
		err = tx.QueryRow(ctx, q, args...).Scan(&in.ID, &in.CreatedAt, &in.UpdatedAt)
	} else {
		err = r.db.QueryRow(ctx, q, args...).Scan(&in.ID, &in.CreatedAt, &in.UpdatedAt)
	}
	return err
}

// DeleteRating implements domain.RatingRepository.
func (r *pgxRatingRepository) DeleteRating(ctx context.Context, resturentID uuid.UUID) (err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	txVal := ctx.Value(TxKey)
	q := `UPDATE rating SET deleted_at = now() WHERE resturent_id = $1`
	args := []interface{}{resturentID}
	if txVal != nil {
		tx := txVal.(pgx.Tx)
		_, err = tx.Exec(ctx, q, args...)
	} else {
		_, err = r.db.Exec(ctx, q, args...)
	}
	return err
}

// Filter implements domain.RatingRepository.
func (r *pgxRatingRepository) Filter(ctx context.Context, in domain.FilterInput, opt domain.QueryOptions) (result []domain.Rating, total int64, err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	txVal := ctx.Value(TxKey)
	// Retrieve the record counts
	cq := `SELECT count(*) FROM rating WHERE deleted_at IS NULL`
	cq, cargs := buildQueryForFilter(in, cq)
	if txVal != nil {
		tx := txVal.(pgx.Tx)
		err = tx.QueryRow(ctx, cq, cargs...).Scan(&total)
	} else {
		err = r.db.QueryRow(ctx, cq, cargs...).Scan(&total)
	}
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return result, total, nil
		}
		return result, total, err
	}

	// Retrive the records
	q := `SELECT * FROM rating WHERE deleted_at IS NULL`
	q, args := buildQueryForFilter(in, q)
	q = buildSortKeysForFilter(in, q)
	q = applyLimitAndOffset(q, opt)
	q = buildSelectorForQuery(q, opt)
	var rows pgx.Rows
	if txVal != nil {
		tx := txVal.(pgx.Tx)
		rows, err = tx.Query(ctx, q, args...)
	} else {
		rows, err = r.db.Query(ctx, q, args...)
	}
	if err != nil {
		return result, total, err
	}
	result, err = pgx.CollectRows(rows, pgx.RowToStructByPos[domain.Rating])
	return result, total, err
}

// GetRatingByResturentID implements domain.RatingRepository.
func (r *pgxRatingRepository) GetRatingByResturentID(ctx context.Context, resturentID uuid.UUID) (result []domain.Rating, err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	txVal := ctx.Value(TxKey)
	q := `SELECT * FROM rating WHERE deleted_at IS NULL AND resturent_id = $1`
	args := []interface{}{resturentID}
	var rows pgx.Rows
	if txVal != nil {
		tx := txVal.(pgx.Tx)
		rows, err = tx.Query(ctx, q, args...)
	} else {
		rows, err = r.db.Query(ctx, q, args...)
	}
	if err != nil {
		return result, err
	}
	result, err = pgx.CollectRows(rows, pgx.RowToStructByPos[domain.Rating])
	return result, err
}

// UpdateRating implements domain.RatingRepository.
func (r *pgxRatingRepository) UpdateRating(ctx context.Context, in *domain.Rating) (err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	txVal := ctx.Value(TxKey)
	q := `UPDATE rating SET name = $1, rating = $2, resturent_id = $3, review = $4, suggestion = $5 WHERE id = $6 RETURNING  updated_at`
	args := []interface{}{in.Name, in.Rating, in.ResturentID, in.Review, in.Suggestion, in.ID}
	if txVal != nil {
		tx := txVal.(pgx.Tx)
		err = tx.QueryRow(ctx, q, args...).Scan(&in.UpdatedAt)
	} else {
		err = r.db.QueryRow(ctx, q, args...).Scan(&in.UpdatedAt)
	}
	log.Printf("updated_at: %v", in)
	return err
}

func (r *pgxRatingRepository) FindByID(ctx context.Context, id uuid.UUID) (result domain.Rating, err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	txVal := ctx.Value(TxKey)
	q := `SELECT * FROM rating WHERE deleted_at IS NULL AND id = $1`
	args := []interface{}{id}
	var rows pgx.Rows
	if txVal != nil {
		tx := txVal.(pgx.Tx)
		rows, err = tx.Query(ctx, q, args...)
	} else {
		rows, err = r.db.Query(ctx, q, args...)
	}
	if err != nil {
		return result, err
	}
	result, err = pgx.CollectOneRow(rows, pgx.RowToStructByPos[domain.Rating])
	if err != nil {
		return result, err
	}
	return result, err
}
