package service

import (
	"context"

	"github.com/gofrs/uuid/v5"
	"github.com/rizwank123/myResturent/internal/domain"
)

type MenuCardService struct {
	mr domain.MenuCardRepository
	tr domain.Transactioner
}

func NewMenuCardService(mr domain.MenuCardRepository, tr domain.Transactioner) domain.MenuCardService {
	return &MenuCardService{
		mr: mr,
		tr: tr,
	}
}

// Create implements domain.MenuCardService.
func (s *MenuCardService) Create(in domain.CreateMenuCardInput) (result domain.MenuCard, err error) {
	result.Name = in.Name
	result.ResturentID = in.ResturentID
	result.Price = in.Price
	result.Size = in.Size
	result.Category = in.Category
	result.FoodType = in.FoodType
	result.MealType = &in.MealType
	result.Image = &in.Image
	result.IsAvailable = in.IsAvailable
	result.Description = &in.Description
	return result, s.mr.Create(context.Background(), &result)

}

// Delete implements domain.MenuCardService.
func (s *MenuCardService) Delete(id uuid.UUID) (err error) {
	return s.mr.Delete(context.Background(), id)
}

// Filter implements domain.MenuCardService.
func (s *MenuCardService) Filter(in domain.FilterInput, opt domain.QueryOptions) (result []domain.MenuCard, total int64, err error) {
	if in.ResturentID != uuid.Nil {
		SetupResturentIDFilter(&in)
	}
	return s.mr.Filter(context.Background(), in, opt)
}

// FindById implements domain.MenuCardService.
func (s *MenuCardService) FindById(id uuid.UUID) (result domain.MenuCard, err error) {
	return s.mr.FindById(context.Background(), id)
}

// FindByResturentID implements domain.MenuCardService.
func (s *MenuCardService) FindByResturentID(resturentID uuid.UUID) (result []domain.MenuCard, err error) {
	return s.mr.FindByResturentID(context.Background(), resturentID)
}

// Update implements domain.MenuCardService.
func (s *MenuCardService) Update(id uuid.UUID, in domain.UpdateMenuCardInput) (result domain.MenuCard, err error) {
	result, err = s.mr.FindById(context.Background(), id)
	if err != nil {
		return result, err
	}

	if in.Name != "" {
		result.Name = in.Name
	}
	if in.Price != 0 {
		result.Price = in.Price
	}
	if in.Size != "" {
		result.Size = in.Size
	}
	if in.Category != "" {
		result.Category = in.Category
	}
	if in.FoodType != "" {
		result.FoodType = in.FoodType
	}
	if in.MealType != "" {
		result.MealType = &in.MealType
	}
	if in.Image != "" {
		result.Image = &in.Image
	}
	if in.IsAvailable != result.IsAvailable {
		result.IsAvailable = in.IsAvailable
	}
	if in.Description != "" {
		result.Description = &in.Description
	}
	return result, s.mr.Update(context.Background(), &result)
}
func (c *MenuCardService) DeleteMenuCard(id uuid.UUID) (err error) {
	return c.mr.Delete(context.Background(), id)
}
