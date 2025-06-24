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

type pgxUserRepository struct {
	db  *pgxpool.Pool
	sqt sq.StatementBuilderType
}

func NewUserRepository(db *pgxpool.Pool) domain.UserRepository {
	return &pgxUserRepository{
		db:  db,
		sqt: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

// Create implements domain.UserRepository.
func (r *pgxUserRepository) Create(ctx context.Context, user *domain.User) (err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	txVal := ctx.Value(TxKey)

	q := `INSERT INTO users (name, email, password, role, mobile, resturent_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, created_at, updated_at`
	args := []interface{}{user.Name, user.Email, user.Password, user.Role, user.Mobile, user.ResturentID}
	if txVal != nil {
		tx := txVal.(pgx.Tx)
		err = tx.QueryRow(ctx, q, args...).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	} else {
		err = r.db.QueryRow(ctx, q, args...).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	}

	if err != nil {
		return err
	}
	return nil
}

// Delete implements domain.UserRepository.
func (r *pgxUserRepository) Delete(ctx context.Context, id uuid.UUID) (err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	txVal := ctx.Value(TxKey)
	q := `UPDATE users SET deleted_at = now() WHERE id = $1`
	args := []interface{}{id}
	if txVal != nil {
		tx := txVal.(pgx.Tx)
		_, err = tx.Exec(ctx, q, args...)
	} else {
		_, err = r.db.Exec(ctx, q, args...)
	}
	return err
}

// Filter implements domain.UserRepository.
func (r *pgxUserRepository) Filter(ctx context.Context, in domain.FilterInput, opt domain.QueryOptions) (result []domain.User, total int64, err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	txVal := ctx.Value(TxKey)
	// Retrieve the record counts
	cq := `SELECT count(*) FROM users WHERE deleted_at IS NULL`
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
	q := `SELECT * FROM users WHERE deleted_at IS NULL`
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
	defer rows.Close()

	if err != nil {
		return result, total, err
	}
	if rows == nil {
		return result, total, nil
	}

	// Collect the results
	result, err = pgx.CollectRows(rows, pgx.RowToStructByNameLax[domain.User])
	return result, total, err

}

// FindByEmail implements domain.UserRepository.
func (r *pgxUserRepository) FindByEmail(ctx context.Context, email string) (result domain.User, err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	txVal := ctx.Value(TxKey)
	q := `SELECT * FROM users WHERE email = $1 AND deleted_at IS NULL LIMIT 1`
	args := []interface{}{email}
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
	result, err = pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[domain.User])
	return result, err
}

// FindById implements domain.UserRepository.
func (r *pgxUserRepository) FindById(ctx context.Context, id uuid.UUID) (result domain.User, err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	txVal := ctx.Value(TxKey)
	q := `SELECT * FROM users WHERE id = $1 AND deleted_at IS NULL LIMIT 1`
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
	result, err = pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[domain.User])
	return result, err
}

// Update implements domain.UserRepository.
func (r *pgxUserRepository) Update(ctx context.Context, user *domain.User) (err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	txVal := ctx.Value(TxKey)
	q := `UPDATE users SET name = $1, email = $2, password = $3, role = $4, mobile = $5, resturent_id = $6 WHERE id = $7  RETURNING updated_at`
	args := []interface{}{user.Name, user.Email, user.Password, user.Role, user.Mobile, user.ResturentID, user.ID}
	if txVal != nil {
		tx := txVal.(pgx.Tx)
		err = tx.QueryRow(ctx, q, args...).Scan(&user.UpdatedAt)
	} else {
		err = r.db.QueryRow(ctx, q, args...).Scan(&user.UpdatedAt)
	}
	return err
}
