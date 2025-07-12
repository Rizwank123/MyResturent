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

type pgxMenuCardRepository struct {
	db  *pgxpool.Pool
	sqt sq.StatementBuilderType
}

func NewMenuCardRepository(db *pgxpool.Pool) domain.MenuCardRepository {
	return &pgxMenuCardRepository{
		db:  db,
		sqt: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

// Create implements domain.MenuCardRepository.
func (r *pgxMenuCardRepository) Create(ctx context.Context, in *domain.MenuCard) (err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	txVal := ctx.Value(TxKey)
	q := `INSERT INTO menu_card (resturent_id, name, price, size, category, food_type, meal_type, image, is_available, offer_price, description) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id, created_at, updated_at`
	args := []interface{}{in.ResturentID, in.Name, in.Price, in.Size, in.Category, in.FoodType, in.MealType, in.Image, in.IsAvailable, in.OfferPrice, in.Description}
	if txVal != nil {
		tx := txVal.(pgx.Tx)
		err = tx.QueryRow(ctx, q, args...).Scan(&in.ID, &in.CreatedAt, &in.UpdatedAt)
	} else {
		err = r.db.QueryRow(ctx, q, args...).Scan(&in.ID, &in.CreatedAt, &in.UpdatedAt)
	}
	return err
}

// Delete implements domain.MenuCardRepository.
func (r *pgxMenuCardRepository) Delete(ctx context.Context, id uuid.UUID) (err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	txVal := ctx.Value(TxKey)
	q := `UPDATE menu_card SET deleted_at = now() WHERE id = $1`
	args := []interface{}{id}
	if txVal != nil {
		tx := txVal.(pgx.Tx)
		_, err = tx.Exec(ctx, q, args...)
	} else {
		_, err = r.db.Exec(ctx, q, args...)
	}
	return err
}

// Filter implements domain.MenuCardRepository.
func (r *pgxMenuCardRepository) Filter(ctx context.Context, in domain.FilterInput, opt domain.QueryOptions) (result []domain.MenuCard, total int64, err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	txVal := ctx.Value(TxKey)
	// Retrieve the record counts
	cq := `SELECT count(*) FROM menu_card WHERE deleted_at IS NULL`
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
	q := `SELECT * FROM menu_card WHERE deleted_at IS NULL`
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
	result, err = pgx.CollectRows(rows, pgx.RowToStructByNameLax[domain.MenuCard])
	return result, total, err
}

// FindById implements domain.MenuCardRepository.
func (r *pgxMenuCardRepository) FindById(ctx context.Context, id uuid.UUID) (result domain.MenuCard, err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	txVal := ctx.Value(TxKey)
	q := `SELECT * FROM menu_card WHERE id = $1 AND deleted_at IS NULL LIMIT 1`
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
	result, err = pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[domain.MenuCard])
	return result, err
}

// FindByResturentID implements domain.MenuCardRepository.
func (r *pgxMenuCardRepository) FindByResturentID(ctx context.Context, resturentID uuid.UUID) (result []domain.MenuCard, err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	txVal := ctx.Value(TxKey)
	q := `SELECT * FROM menu_card WHERE deleted_at IS NULL AND resturent_id = $1`
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
	result, err = pgx.CollectRows(rows, pgx.RowToStructByNameLax[domain.MenuCard])
	return result, err
}

// Update implements domain.MenuCardRepository.
func (r *pgxMenuCardRepository) Update(ctx context.Context, in *domain.MenuCard) (err error) {
	if ctx == nil {
		ctx = context.Background()
	}
	txVal := ctx.Value(TxKey)
	q := `UPDATE menu_card SET resturent_id = $2, name = $3, price = $4, size = $5, category = $6, food_type = $7, meal_type = $8, image = $9, is_available = $10, description = $11, offer_price = $12 WHERE id = $1`
	args := []interface{}{in.ID, in.ResturentID, in.Name, in.Price, in.Size, in.Category, in.FoodType, in.MealType, in.Image, in.IsAvailable, in.Description, in.OfferPrice}
	if txVal != nil {
		tx := txVal.(pgx.Tx)
		_, err = tx.Exec(ctx, q, args...)
	} else {
		_, err = r.db.Exec(ctx, q, args...)
	}
	return err
}
