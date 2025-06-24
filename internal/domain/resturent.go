package domain

import (
	"context"

	"github.com/gofrs/uuid/v5"
)

type (
	// Resturent defines the model for resturent
	Resturent struct {
		Base
		Name    string  `json:"name" db:"name"  example:"Alamgiri Resturent"`
		Address Address `json:"address,omitempty" db:"address" `
		License string  `json:"license,omitempty" db:"license" example:"license"`
		BaseAudit
	} // @name Resturent

	// CreateResturentInput define the model for creating new resturent
	CreateResturentInput struct {
		Name    string  `json:"name" db:"name"  example:"Alamgiri Resturent"`
		Address Address `json:"address,omitempty" db:"address" `
		License string  `json:"license,omitempty" db:"license" example:"license"`
	} // @name CreateResturentInput

	// UpdateResturentInput define the model for updating resturent
	UpdateResturentInput struct {
		Name    string  `json:"name" db:"name"  example:"Alamgiri Resturent"`
		Address Address `json:"address,omitempty" db:"address" `
		License string  `json:"license,omitempty" db:"license" example:"license"`
	} // @name UpdateResturentInput
)

type (
	// ResturentRepository defines the methods which a resturent repository must implement
	ResturentRepository interface {
		// CreateResturent creates a new resturent
		Create(ctx context.Context, in *Resturent) (err error)
		// FindById returns a resturent by id
		FindById(ctx context.Context, id uuid.UUID) (result Resturent, err error)
		// UpdateResturent updates a resturent
		Update(ctx context.Context, in *Resturent) (err error)
		// Filter returns a list of resturent
		// limit and offset are used for pagination.
		// total is the total number of entities.
		Filter(ctx context.Context, in FilterInput, opt QueryOptions) (result []Resturent, total int64, err error)
		// DeleteResturent deletes a resturent
		Delete(ctx context.Context, id uuid.UUID) (err error)
	}

	// ResturentService defines the methods which a resturent service must implement
	ResturentService interface {
		// CreateResturent creates a new resturent
		Create(in CreateResturentInput) (result Resturent, err error)
		// FindById returns a resturent by id
		FindById(id uuid.UUID) (result Resturent, err error)
		// UpdateResturent updates a resturent
		Update(id uuid.UUID, in UpdateResturentInput) (result Resturent, err error)
		// Filter returns a list of resturent
		// limit and offset are used for pagination.
		// total is the total number of entities.
		Filter(in FilterInput, opt QueryOptions) (result []Resturent, total int64, err error)
		// DeleteResturent deletes a resturent
		Delete(id uuid.UUID) (err error)
	}
)
