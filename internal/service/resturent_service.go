package service

import (
	"context"

	"github.com/gofrs/uuid/v5"
	"github.com/rizwank123/myResturent/internal/domain"
)

type ResturentSeriveImple struct {
	rr domain.ResturentRepository
}

func NewResturentService(rr domain.ResturentRepository) domain.ResturentService {
	return &ResturentSeriveImple{
		rr: rr,
	}
}

// Create implements domain.ResturentService.
func (s *ResturentSeriveImple) Create(in domain.CreateResturentInput) (result domain.Resturent, err error) {
	result.Name = in.Name
	result.Address = in.Address
	result.License = in.License
	return result, s.rr.Create(context.Background(), &result)
}

// Delete implements domain.ResturentService.
func (s *ResturentSeriveImple) Delete(id uuid.UUID) (err error) {
	return s.rr.Delete(context.Background(), id)
}

// Filter implements domain.ResturentService.
func (s *ResturentSeriveImple) Filter(in domain.FilterInput, opt domain.QueryOptions) (result []domain.Resturent, total int64, err error) {
	return s.rr.Filter(context.Background(), in, opt)
}

// FindById implements domain.ResturentService.
func (s *ResturentSeriveImple) FindById(id uuid.UUID) (result domain.Resturent, err error) {
	return s.rr.FindById(context.Background(), id)
}

// Update implements domain.ResturentService.
func (s *ResturentSeriveImple) Update(id uuid.UUID, in domain.UpdateResturentInput) (result domain.Resturent, err error) {
	result, err = s.rr.FindById(context.Background(), id)
	if err != nil {
		return result, err
	}

	if in.Name != "" {
		result.Name = in.Name
	}
	if in.Address.City != "" {
		result.Address.City = in.Address.City
	}
	if in.Address.Street != "" {
		result.Address.Street = in.Address.Street
	}
	if in.Address.State != "" {
		result.Address.State = in.Address.State
	}
	if in.Address.Pincode != "" {
		result.Address.Pincode = in.Address.Pincode
	}
	if in.Address.Country != "" {
		result.Address.Country = in.Address.Country
	}
	if in.License != "" {
		result.License = in.License
	}
	return result, s.rr.Update(context.Background(), &result)
}
