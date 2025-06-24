package service

import (
	"context"

	"github.com/gofrs/uuid/v5"
	"github.com/rizwank123/myResturent/internal/domain"
)

type ratingServiceImpl struct {
	rtr domain.RatingRepository
}

func NewRatingService(rtr domain.RatingRepository) domain.RatingService {
	return &ratingServiceImpl{
		rtr: rtr,
	}
}

// CreateRating implements domain.RatingService.
func (s *ratingServiceImpl) CreateRating(in domain.CreateRatingInput) (result domain.Rating, err error) {
	result.ResturentID = in.ResturentID
	result.Name = in.Name
	result.Rating = in.Rating
	result.Review = &in.Review
	result.Suggestion = &in.Suggestion
	return result, s.rtr.CreateRating(context.Background(), &result)
}

// DeleteRating implements domain.RatingService.
func (s *ratingServiceImpl) DeleteRating(ID uuid.UUID) (err error) {
	return s.rtr.DeleteRating(context.Background(), ID)
}

// Filter implements domain.RatingService.
func (s *ratingServiceImpl) Filter(in domain.FilterInput, opt domain.QueryOptions) (result []domain.Rating, total int64, err error) {
	return s.rtr.Filter(context.Background(), in, opt)
}

// GetRatingByResturentID implements domain.RatingService.
func (s *ratingServiceImpl) GetRatingByResturentID(resturentID uuid.UUID) (result []domain.Rating, err error) {
	return s.rtr.GetRatingByResturentID(context.Background(), resturentID)
}

// UpdateRating implements domain.RatingService.
func (s *ratingServiceImpl) UpdateRating(id uuid.UUID, in domain.UpdateRatingInput) (result domain.Rating, err error) {
	result, err = s.rtr.FindByID(context.Background(), id)
	if err != nil {
		return result, err
	}
	if in.Name != "" {
		result.Name = in.Name
	}
	if in.Rating != result.Rating {
		result.Rating = in.Rating
	}
	if in.Review != "" {
		result.Review = &in.Review
	}
	if in.Suggestion != "" {
		result.Suggestion = &in.Suggestion
	}
	err = s.rtr.UpdateRating(context.Background(), &result)

	return result, err
}

// FindByID implements domain.RatingService.
func (s *ratingServiceImpl) FindByID(id uuid.UUID) (result domain.Rating, err error) {
	return s.rtr.FindByID(context.Background(), id)
}
