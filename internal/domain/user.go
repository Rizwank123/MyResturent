package domain

import (
	"context"

	"github.com/gofrs/uuid/v5"
)

type (
	UserRole string
	// User defines the model for user
	User struct {
		Base
		Name        string    `json:"username" db:"name" example:"username"`
		Password    string    `json:"password" db:"password" example:"password"`
		Email       string    `json:"email" db:"email" example:"email"`
		Role        UserRole  `json:"role" db:"role" example:"role"`
		Mobile      string    `json:"mobile" db:"mobile" example:"mobile"`
		ResturentID uuid.UUID `json:"resturent_id" db:"resturent_id" example:"c816d9e0-63e5-11ec-90d6-0242ac120003"`
		BaseAudit
	} // @name User

	// CreateUserInput define the model for creating new user
	CreateUserInput struct {
		Name          string `json:"username" db:"name" example:"username"`
		ResturentName string `json:"resturent_name" db:"resturent_name" example:"resturent_name"`
		Password      string `json:"password" db:"password" example:"password"`
		Email         string `json:"email" db:"email" example:"email"`
		Role          string `json:"role" db:"role" example:"role"`
		Mobile        string `json:"mobile" db:"mobile" example:"mobile"`
		ResturentID   uuid.UUID
	} // @name CreateUserInput

	// UpdateUserInput define the model for updating user
	UpdateUserInput struct {
		Name        string `json:"username" db:"name" example:"username"`
		Password    string `json:"password" db:"password" example:"password"`
		Email       string `json:"email" db:"email" example:"email"`
		Role        string `json:"role" db:"role" example:"role"`
		Mobile      string `json:"mobile" db:"mobile" example:"mobile"`
		ResturentID uuid.UUID
	} // @name UpdateUserInput
	// LoginUserInput define the model for login user
	LoginUserInput struct {
		Email    string `json:"email" db:"email" example:"email"`
		Password string `json:"password" db:"password" example:"password"`
	}
	// LoginResponse define the model for login response
	LoginResponse struct {
		Token string `json:"token" db:"token" example:"token"`
	} // @name LoginResponse
)

type (
	// UserRepository Defines the methods that any user repository should implement
	UserRepository interface {
		// FindById Find user by id
		FindById(ctx context.Context, id uuid.UUID) (result User, err error)

		// FindByEmail Find user by email
		FindByEmail(ctx context.Context, email string) (result User, err error)

		// Filter filters user by criteria.
		// limit and offset are used for pagination.
		// total is the total number of entities.
		Filter(ctx context.Context, in FilterInput, opt QueryOptions) (result []User, total int64, err error)
		// Create Create user
		Create(ctx context.Context, user *User) (err error)

		// Update Update user
		Update(ctx context.Context, user *User) (err error)
		// Delete Delete user
		Delete(ctx context.Context, id uuid.UUID) (err error)
	}

	// UserService Defines the methods that any user service should implement
	UserService interface {
		// FindById Find user by id
		FindById(id uuid.UUID) (result User, err error)

		// Filter filters users by criteria.
		// limit and offset are used for pagination.
		// total is the total number of entities.
		Filter(in FilterInput, opt QueryOptions) (result []User, total int64, err error)

		// FindByEmail Find user by email
		FindByEmail(email string) (result User, err error)

		// Create Create user
		Create(in CreateUserInput) (result User, err error)

		// Update Update user
		Update(id uuid.UUID, in UpdateUserInput) (result User, err error)

		// Login Login user
		Login(in LoginUserInput) (result LoginResponse, err error)
	}
)
