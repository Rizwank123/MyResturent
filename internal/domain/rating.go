package domain

import (
	"context"

	"github.com/gofrs/uuid/v5"
)

type (
	// Rating defines the model for rating
	Rating struct {
		Base
		ResturentID uuid.UUID `json:"resturent_id" db:"resturent_id" example:"c816d9e0-63e5-11ec-90d6-0242ac120003"`
		Name        string    `json:"name,omitempty" db:"name"  example:"Alex t"`
		Rating      float64   `json:"rating" db:"rating" example:"5"`
		Review      *string   `json:"review,omitempty" db:"review" example:"review"`
		Suggestion  *string   `json:"suggestion,omitempty" db:"suggestion" example:"suggestion"`
		BaseAudit
	} // @name Rating

	// CreateRatingInput define the model for creating new rating
	CreateRatingInput struct {
		ResturentID uuid.UUID `json:"resturent_id" db:"resturent_id" example:"c816d9e0-63e5-11ec-90d6-0242ac120003"`
		Name        string    `json:"name,omitempty" db:"name"  example:"Alex t"`
		Rating      float64   `json:"rating" db:"rating" example:"5"`
		Review      string    `json:"review,omitempty" db:"review" example:"review"`
		Suggestion  string    `json:"suggestion,omitempty" db:"suggestion" example:"Make sure reduce serving time"`
	} // @name CreateRatingInput
	// UpdateRatingInput define the model for updating rating
	UpdateRatingInput struct {
		Name       string  `json:"name,omitempty" db:"name" example:"Alex t"`
		Rating     float64 `json:"rating" db:"rating" example:"5"`
		Review     string  `json:"review,omitempty" db:"review" example:"Food was delicious"`
		Suggestion string  `json:"suggestion,omitempty" db:"suggestion" example:"Make sure reduce serving time"`
	}
)

type (
	// RatingRepository defines the methods that any rating repository should implement
	RatingRepository interface {
		// CreateRating creates a new rating
		CreateRating(ctx context.Context, in *Rating) (err error)
		// GetRatingByResturentID gets rating by resturent id
		GetRatingByResturentID(ctx context.Context, resturentID uuid.UUID) (result []Rating, err error)
		// FindByID finds a rating by id
		FindByID(ctx context.Context, id uuid.UUID) (result Rating, err error)
		// UpdateRating updates a rating
		UpdateRating(ctx context.Context, in *Rating) (err error)
		// DeleteRating deletes a rating
		DeleteRating(ctx context.Context, resturentID uuid.UUID) (err error)
		// Filter filters rating by criteria.
		// limit and offset are used for pagination.
		// total is the total number of entities.
		Filter(ctx context.Context, in FilterInput, opt QueryOptions) (result []Rating, total int64, err error)
	}

	// RatingService defines the methods that any rating service should implement
	RatingService interface {
		// CreateRating creates a new rating
		CreateRating(in CreateRatingInput) (result Rating, err error)
		// GetRatingByResturentID gets rating by resturent id
		GetRatingByResturentID(resturentID uuid.UUID) (result []Rating, err error)
		// FindByID finds a rating by id
		FindByID(id uuid.UUID) (result Rating, err error)
		// UpdateRating updates a rating
		UpdateRating(id uuid.UUID, in UpdateRatingInput) (result Rating, err error)
		// DeleteRating deletes a rating
		DeleteRating(ID uuid.UUID) (err error)
		// Filter filters rating by criteria.
		// limit and offset are used for pagination.
		// total is the total number of entities.
		Filter(in FilterInput, opt QueryOptions) (result []Rating, total int64, err error)
	}
)
