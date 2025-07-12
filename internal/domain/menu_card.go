package domain

import (
	"context"

	"github.com/gofrs/uuid/v5"
)

type (
	// Category defines the type for category
	Category string
	// FoodType defines the type for food type
	FoodType string
	// MealType defines the type for meal type
	MealType string
	// MenuCard defines the model for menu card
	MenuCard struct {
		Base
		ResturentID uuid.UUID `json:"resturent_id" db:"resturent_id" example:"c816d9e0-63e5-11ec-90d6-0242ac120003"`
		Name        string    `json:"name,omitempty" db:"name"  example:"Alex t"`
		Price       float64   `json:"price" db:"price" example:"5"`
		OfferPrice  *float64  `json:"offer_price" db:"offer_price" example:"5"`
		Category    Category  `json:"category,omitempty" db:"category" example:"Veg"`
		Size        string    `json:"size,omitempty" db:"size" example:"half plate"`
		Image       *string   `json:"image,omitempty" db:"image" example:"image"`
		FoodType    FoodType  `json:"food_type,omitempty" db:"food_type" example:"Starter"`
		MealType    *MealType `json:"meal_type,omitempty" db:"meal_type" example:"Dinner"`
		IsAvailable bool      `json:"is_available" db:"is_available" example:"true"`
		Description *string   `json:"description,omitempty" db:"description" example:"description"`
		BaseAudit
	} // @name MenuCard

	// CreateMenuCardInput define the model for creating new menu card
	CreateMenuCardInput struct {
		ResturentID uuid.UUID `json:"resturent_id" db:"resturent_id" example:"c816d9e0-63e5-11ec-90d6-0242ac120003"`
		Name        string    `json:"name,omitempty" db:"name"  example:"Butter Chicken"`
		Price       float64   `json:"price" db:"price" example:"300"`
		OfferPrice  float64   `json:"offer_price,omitempty" db:"offer_price" example:"300"`
		Size        string    `json:"size,omitempty" db:"size"  example:"half plate"`
		Category    Category  `json:"category,omitempty" db:"category" example:"veg"`
		FoodType    FoodType  `json:"food_type,omitempty" db:"food_type" example:"Main Course"`
		MealType    MealType  `json:"meal_type,omitempty" db:"meal_type" example:"DiNNER"`
		Image       string    `json:"image,omitempty" db:"image" example:"image"`
		IsAvailable bool      `json:"is_available" db:"is_available" example:"true"`
		Description string    `json:"description,omitempty" db:"description" example:"description"`
	} // @name CreateMenuCardInput
	// UpdateMenuCardInput define the model for updating menu card
	UpdateMenuCardInput struct {
		ResturentID uuid.UUID `json:"resturent_id" db:"resturent_id" example:"c816d9e0-63e5-11ec-90d6-0242ac120003"`
		Name        string    `json:"name,omitempty" db:"name"  example:"Butter Chicken"`
		Price       float64   `json:"price" db:"price" example:"300"`
		OfferPrice  float64   `json:"offer_price,omitempty" db:"offer_price" example:"300"`
		Size        string    `json:"size,omitempty" db:"size"  example:"half plate"`
		Category    Category  `json:"category,omitempty" db:"category" example:"veg"`
		FoodType    FoodType  `json:"food_type,omitempty" db:"food_type" example:"Main Course"`
		MealType    MealType  `json:"meal_type,omitempty" db:"meal_type" example:"BREAKFAST"`
		Image       string    `json:"image,omitempty" db:"image" example:"image"`
		IsAvailable bool      `json:"is_available" db:"is_available" example:"true"`
		Description string    `json:"description,omitempty" db:"description" example:"description"`
	} // @name UpdateMenuCardInput
)

type (
	// MenuCardRepository defines the methods that any menu card repository should implement
	MenuCardRepository interface {
		// FindByID Find menu card by id
		FindById(ctx context.Context, id uuid.UUID) (result MenuCard, err error)
		// FindByResturentID Find menu card by resturent id
		FindByResturentID(ctx context.Context, resturentID uuid.UUID) (result []MenuCard, err error)
		// Filter filters menu card by criteria.
		// limit and offset are used for pagination.
		// total is the total number of entities.
		Filter(ctx context.Context, in FilterInput, opt QueryOptions) (result []MenuCard, total int64, err error)
		// Create Create new menu card
		Create(ctx context.Context, input *MenuCard) (err error)
		// Update Update menu card
		Update(ctx context.Context, input *MenuCard) (err error)
		// Delete Delete menu card
		Delete(ctx context.Context, id uuid.UUID) (err error)
	}
	// MenuCardService defines the methods that any menu card service should implement
	MenuCardService interface {
		// FindByID Find menu card by id
		FindById(id uuid.UUID) (result MenuCard, err error)
		// FindByResturentID Find menu card by resturent id
		FindByResturentID(resturentID uuid.UUID) (result []MenuCard, err error)
		// Filter filters menu card by criteria.
		// limit and offset are used for pagination.
		// total is the total number of entities.
		Filter(in FilterInput, opt QueryOptions) (result []MenuCard, total int64, err error)
		// Create Create new menu card
		Create(in CreateMenuCardInput) (result MenuCard, err error)
		// Update Update menu card
		Update(id uuid.UUID, in UpdateMenuCardInput) (result MenuCard, err error)
		// Delete Delete menu card
		Delete(id uuid.UUID) (err error)
	}
)
