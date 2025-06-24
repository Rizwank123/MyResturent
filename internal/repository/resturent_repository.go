package repository

import (
	"context"
	"errors"

	sq "github.com/Masterminds/squirrel"
	"github.com/gofrs/uuid/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rizwank123/myResturent/internal/domain"
)

type pgxResturentRepository struct {
	db  *pgxpool.Pool
	sqt sq.StatementBuilderType
}

func NewResturentRepository(db *pgxpool.Pool) domain.ResturentRepository {
	return &pgxResturentRepository{
		db:  db,
		sqt: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

// Create implements domain.ResturentRepository.
func (r *pgxResturentRepository) Create(ctx context.Context, in *domain.Resturent) (err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	txVal := ctx.Value(TxKey)
	q := `INSERT INTO resturent (name, address, license) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at`
	args := []interface{}{in.Name, in.Address, in.License}
	if txVal != nil {
		tx := txVal.(pgx.Tx)
		err = tx.QueryRow(ctx, q, args...).Scan(&in.ID, &in.CreatedAt, &in.UpdatedAt)
	} else {
		err = r.db.QueryRow(ctx, q, args...).Scan(&in.ID, &in.CreatedAt, &in.UpdatedAt)
	}
	return err
}

// Delete implements domain.ResturentRepository.
func (r *pgxResturentRepository) Delete(ctx context.Context, id uuid.UUID) (err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	txVal := ctx.Value(TxKey)
	q := `UPDATE resturent SET deleted_at = now() WHERE id = $1`
	args := []interface{}{id}
	if txVal != nil {
		tx := txVal.(pgx.Tx)
		_, err = tx.Exec(ctx, q, args...)
	} else {
		_, err = r.db.Exec(ctx, q, args...)
	}
	return err
}

// Filter implements domain.ResturentRepository.
func (r *pgxResturentRepository) Filter(ctx context.Context, in domain.FilterInput, opt domain.QueryOptions) (result []domain.Resturent, total int64, err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	txVal := ctx.Value(TxKey)
	// Retrieve the record counts
	cq := `SELECT count(*) FROM resturent WHERE deleted_at IS NULL`
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
	q := `SELECT * FROM resturent WHERE deleted_at IS NULL`
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

	// Collect the results
	result, err = pgx.CollectRows(rows, pgx.RowToStructByNameLax[domain.Resturent])
	return result, total, err
}

// FindById implements domain.ResturentRepository.
func (r *pgxResturentRepository) FindById(ctx context.Context, id uuid.UUID) (result domain.Resturent, err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	txVal := ctx.Value(TxKey)
	q := `SELECT * FROM resturent WHERE id = $1 AND deleted_at IS NULL LIMIT 1`
	args := []interface{}{id}
	var rows pgx.Rows
	if txVal != nil {
		tx := txVal.(pgx.Tx)
		rows, err = tx.Query(ctx, q, args...)
	} else {
		rows, err = r.db.Query(ctx, q, args...)
	}
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return result, domain.DataNotFoundError{}
		}
		return result, err
	}

	// Collect the results
	result, err = pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[domain.Resturent])
	return result, err
}

// Update implements domain.ResturentRepository.
func (r *pgxResturentRepository) Update(ctx context.Context, in *domain.Resturent) (err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	txVal := ctx.Value(TxKey)
	q := `UPDATE resturent SET name = $2, address = $3, license = $4, updated_at = now() WHERE id = $1`
	args := []interface{}{in.ID, in.Name, in.Address, in.License}
	if txVal != nil {
		tx := txVal.(pgx.Tx)
		_, err = tx.Exec(ctx, q, args...)
	} else {
		_, err = r.db.Exec(ctx, q, args...)
	}
	return err
}
