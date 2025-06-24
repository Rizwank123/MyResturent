package controller

import (
	"net/http"

	"github.com/gofrs/uuid/v5"
	"github.com/labstack/echo/v4"
	"github.com/rizwank123/myResturent/internal/domain"
	"github.com/rizwank123/myResturent/internal/http/transport"
)

type RatingController struct {
	rts domain.RatingService
}

func NewRatingController(rts domain.RatingService) RatingController {
	return RatingController{
		rts: rts,
	}
}

// FilterRatings filters ratings by criteria
//
//	@Summary		Filter ratings
//	@Description	Filter ratings by criteria
//	@Tags			Rating
//	@ID				filterRatings
//	@Accept			json
//	@Produce		json
//	@Param			page	query		int					false	"Page number"
//	@Param			size	query		int					false	"Page size"
//	@Param			in		body		domain.FilterInput	true	"Filter criteria"
//	@Success		200		{object}	domain.PaginationResponse{data=[]domain.Rating}
//	@Failure		400		{object}	domain.InvalidRequestError
//	@Failure		401		{object}	domain.UnauthorizedError
//	@Failure		403		{object}	domain.ForbiddenAccessError
//	@Failure		500		{object}	domain.SystemError
//	@Router			/rating/filter [post]
func (c *RatingController) Filter(ctx echo.Context) error {
	var in domain.FilterInput
	err := transport.DecodeAndValidateRequestBody(ctx, &in)
	if err != nil {
		return err
	}
	opt := transport.DecodeQueryOptions(ctx)
	result, total, err := c.rts.Filter(in, opt)
	if err != nil {
		return err
	}
	return transport.SendPaginationResponse(ctx, http.StatusOK, result, total)
}

// CreateRating creates a new rating
//
//	@Summary		Create a rating
//	@Description	Create a rating
//	@Tags			Rating
//	@ID				createRating
//	@Accept			json
//	@Produce		json
//	@Param			in	body		domain.CreateRatingInput	true	"Payload"
//	@Success		200	{object}	domain.Rating
//	@Failure		400	{object}	domain.InvalidRequestError
//	@Failure		401	{object}	domain.UnauthorizedError
//	@Failure		403	{object}	domain.ForbiddenAccessError
//	@Failure		500	{object}	domain.SystemError
//	@Router			/rating [post]
func (c *RatingController) CreateRating(ctx echo.Context) error {
	var in domain.CreateRatingInput
	err := transport.DecodeAndValidateRequestBody(ctx, &in)
	if err != nil {
		return err
	}
	result, err := c.rts.CreateRating(in)
	if err != nil {
		return err
	}
	return transport.SendResponse(ctx, http.StatusCreated, result)
}

// UpdateRating updates an existing rating
//
//	@Summary		Update a rating
//	@Description	Update a rating with the given ID
//	@Tags			Rating
//	@ID				updateRating
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string						true	"Rating ID"
//	@Param			in	body		domain.UpdateRatingInput	true	"Payload"
//	@Success		200	{object}	domain.Rating
//	@Failure		400	{object}	domain.InvalidRequestError
//	@Failure		401	{object}	domain.UnauthorizedError
//	@Failure		403	{object}	domain.ForbiddenAccessError
//	@Failure		404	{object}	domain.DataNotFoundError
//	@Failure		500	{object}	domain.SystemError
//	@Router			/rating/{id} [put]
func (c *RatingController) UpdateRating(ctx echo.Context) error {
	id, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		return err
	}
	var in domain.UpdateRatingInput
	err = transport.DecodeAndValidateRequestBody(ctx, &in)
	if err != nil {
		return err
	}
	result, err := c.rts.UpdateRating(id, in)
	if err != nil {
		return err
	}
	return transport.SendResponse(ctx, http.StatusOK, result)
}

// FindRatingByID finds a rating by id
//
//	@Summary		Find a rating by id
//	@Description	Find a rating by id
//	@Tags			Rating
//	@ID				findRatingByID
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Rating ID"
//	@Success		200	{object}	domain.Rating
//	@Failure		400	{object}	domain.InvalidRequestError
//	@Failure		401	{object}	domain.UnauthorizedError
//	@Failure		403	{object}	domain.ForbiddenAccessError
//	@Failure		404	{object}	domain.DataNotFoundError
//	@Failure		500	{object}	domain.SystemError
//	@Router			/rating/{id} [get]
func (c *RatingController) FindByID(ctx echo.Context) error {
	id, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		return err
	}
	result, err := c.rts.FindByID(id)
	if err != nil {
		return err
	}
	return transport.SendResponse(ctx, http.StatusOK, result)
}

// DeleteRating deletes a rating
//
//	@Summary		Delete a rating
//	@Description	Delete a rating by id
//	@Tags			Rating
//	@ID				deleteRating
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Rating ID"
//	@Success		204	{object}	domain.BaseResponse
//	@Failure		400	{object}	domain.InvalidRequestError
//	@Failure		401	{object}	domain.UnauthorizedError
//	@Failure		403	{object}	domain.ForbiddenAccessError
//	@Failure		404	{object}	domain.DataNotFoundError
//	@Failure		500	{object}	domain.SystemError
//	@Router			/rating/{id} [delete]
func (c *RatingController) DeleteRating(ctx echo.Context) error {
	id, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		return err
	}
	err = c.rts.DeleteRating(id)
	if err != nil {
		return err
	}
	return transport.SendResponse(ctx, http.StatusNoContent, nil)
}
