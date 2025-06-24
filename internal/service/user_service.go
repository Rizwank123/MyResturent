package service

import (
	"context"
	"errors"

	"github.com/gofrs/uuid/v5"
	"github.com/rizwank123/myResturent/internal/domain"
	"github.com/rizwank123/myResturent/internal/pkg/security"
	"golang.org/x/crypto/bcrypt"
)

type userServiceImpl struct {
	rs  domain.ResturentRepository
	smc security.Manager
	tr  domain.Transactioner
	ur  domain.UserRepository
}

func NewUserService(rs domain.ResturentRepository, smc security.Manager, tr domain.Transactioner, ur domain.UserRepository) domain.UserService {
	return &userServiceImpl{
		rs:  rs,
		smc: smc,
		tr:  tr,
		ur:  ur,
	}
}

// Create implements domain.UserService.
func (s *userServiceImpl) Create(in domain.CreateUserInput) (result domain.User, err error) {

	ctx := context.Background()

	ctx, err = s.tr.Begin(ctx)
	if err != nil {
		return result, err
	}

	defer func() {
		s.tr.Rollback(ctx, err)
	}()
	resturent := domain.Resturent{
		Name:    in.ResturentName,
		Address: domain.Address{},
		License: "",
	}

	err = s.rs.Create(ctx, &resturent)

	if err != nil {
		return result, err
	}
	result.Name = in.Name
	result.Email = in.Email

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return result, err
	}
	result.Password = string(hashedPassword)
	result.Role = domain.UserRole(in.Role)
	result.Mobile = in.Mobile
	result.ResturentID = resturent.ID
	err = s.ur.Create(ctx, &result)

	if err != nil {
		return result, err
	}

	if err != nil {
		return result, err
	}
	err = s.tr.Commit(ctx)
	if err != nil {
		return result, err
	}
	return result, nil

}

// Filter implements domain.UserService.
func (s *userServiceImpl) Filter(in domain.FilterInput, opt domain.QueryOptions) (result []domain.User, total int64, err error) {
	return s.ur.Filter(context.Background(), in, opt)
}

// FindByEmail implements domain.UserService.
func (s *userServiceImpl) FindByEmail(email string) (result domain.User, err error) {
	return s.ur.FindByEmail(context.Background(), email)
}

// FindById implements domain.UserService.
func (s *userServiceImpl) FindById(id uuid.UUID) (result domain.User, err error) {
	return s.ur.FindById(context.Background(), id)
}

// Login implements domain.UserService.
func (s *userServiceImpl) Login(in domain.LoginUserInput) (result domain.LoginResponse, err error) {
	user, err := s.ur.FindByEmail(context.Background(), in.Email)
	if err != nil {
		return result, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password)); err != nil {
		return result, errors.New("invalid password")
	}
	t := security.TokenMetaData{
		UserID:      user.ID.String(),
		Role:        string(user.Role),
		ResturentID: user.ResturentID.String(),
	}
	token, err := s.smc.GenerateToken(t)
	if err != nil {
		return result, err
	}

	return domain.LoginResponse{Token: token}, err
}

// Update implements domain.UserService.
func (s *userServiceImpl) Update(id uuid.UUID, in domain.UpdateUserInput) (result domain.User, err error) {
	result, err = s.ur.FindById(context.Background(), id)
	if err != nil {
		return result, err
	}

	if in.Name != "" {
		result.Name = in.Name
	}
	if in.Email != "" {
		result.Email = in.Email
	}
	if in.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
		if err != nil {
			return result, errors.New("failed to hash password")
		}
		result.Password = string(hashedPassword)
	}
	if in.Role != "" {
		result.Role = domain.UserRole(in.Role)
	}
	if in.Mobile != "" {
		result.Mobile = in.Mobile
	}
	return result, s.ur.Update(context.Background(), &result)
}
