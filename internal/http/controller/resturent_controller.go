package controller

import (
	"net/http"

	"github.com/gofrs/uuid/v5"
	"github.com/labstack/echo/v4"
	"github.com/rizwank123/myResturent/internal/domain"
	"github.com/rizwank123/myResturent/internal/http/transport"
)

type ResturentController struct {
	rs domain.ResturentService
}

func NewResturentController(rs domain.ResturentService) ResturentController {
	return ResturentController{
		rs: rs,
	}
}

// CreateResturent creates a new resturent
//
//	@Summary		Create a resturent
//	@Description	Create a resturent
//	@Tags			Resturent
//	@ID				createResturent
//	@Accept			json
//	@Produce		json
//	@Security		JWT
//	@Param			Authorization	header		string						true	"Bearer "
//	@Param			in				body		domain.CreateResturentInput	true	"Payload"
//	@Success		201				{object}	domain.BaseResponse{data=domain.Resturent}
//	@Failure		400				{object}	domain.InvalidRequestError
//	@Failure		401				{object}	domain.UnauthorizedError
//	@Failure		403				{object}	domain.ForbiddenAccessError
//	@Failure		500				{object}	domain.SystemError
//	@Router			/resturent [post]
func (c *ResturentController) CreateResturent(ctx echo.Context) error {
	var in domain.CreateResturentInput
	err := transport.DecodeAndValidateRequestBody(ctx, &in)
	if err != nil {
		return err
	}
	result, err := c.rs.Create(in)
	if err != nil {
		return err
	}
	return transport.SendResponse(ctx, http.StatusCreated, result)
}

// FindResturentById finds a resturent by id
//
//	@Summary		Find a resturent by id
//	@Description	Find a resturent by id
//	@Tags			Resturent
//	@ID				findResturentById
//	@Accept			json
//	@Produce		json
//	@Security		JWT
//	@Param			Authorization	header		string	true	"Bearer "
//	@Param			id				path		string	true	"Resturent id"
//	@Success		200				{object}	domain.BaseResponse{data=domain.Resturent}
//	@Failure		400				{object}	domain.InvalidRequestError
//	@Failure		401				{object}	domain.UnauthorizedError
//	@Failure		403				{object}	domain.ForbiddenAccessError
//	@Failure		404				{object}	domain.DataNotFoundError
//	@Failure		500				{object}	domain.SystemError
//	@Router			/resturent/{id} [get]
func (c *ResturentController) FindById(ctx echo.Context) error {
	id, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		return err
	}
	result, err := c.rs.FindById(id)
	if err != nil {
		return err
	}
	return transport.SendResponse(ctx, http.StatusOK, result)
}

// UpdateResturent updates a resturent
//
//	@Summary		Update a resturent
//	@Description	Update a resturent
//	@Tags			Resturent
//	@ID				updateResturent
//	@Accept			json
//	@Produce		json
//	@Security		JWT
//	@Param			Authorization	header		string						true	"Bearer "
//	@Param			id				path		string						true	"Resturent id"
//	@Param			in				body		domain.UpdateResturentInput	true	"Payload"
//	@Success		200				{object}	domain.BaseResponse{data=domain.Resturent}
//	@Failure		400				{object}	domain.InvalidRequestError
//	@Failure		401				{object}	domain.UnauthorizedError
//	@Failure		403				{object}	domain.ForbiddenAccessError
//	@Failure		404				{object}	domain.DataNotFoundError
//	@Failure		500				{object}	domain.SystemError
//	@Router			/resturent/{id} [put]
func (c *ResturentController) UpdateResturent(ctx echo.Context) error {
	id, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		return err
	}
	var in domain.UpdateResturentInput
	err = transport.DecodeAndValidateRequestBody(ctx, &in)
	if err != nil {
		return err
	}
	result, err := c.rs.Update(id, in)
	if err != nil {
		return err
	}
	return transport.SendResponse(ctx, http.StatusOK, result)
}

// FilterResturents filters resturents by criteria
//
//	@Summary		Filter resturents
//	@Description	Filter resturents by criteria
//	@Tags			Resturent
//	@ID				filterResturents
//	@Accept			json
//	@Produce		json
//	@Param			in	body		domain.FilterInput	true	"Filter criteria"
//	@Success		200	{object}	domain.PaginationResponse{data=[]domain.Resturent}
//	@Failure		400	{object}	domain.InvalidRequestError
//	@Failure		401	{object}	domain.UnauthorizedError
//	@Failure		403	{object}	domain.ForbiddenAccessError
//	@Failure		500	{object}	domain.SystemError
//	@Router			/resturent/filter [post]
func (c *ResturentController) Filter(ctx echo.Context) error {
	var in domain.FilterInput
	err := transport.DecodeAndValidateRequestBody(ctx, &in)
	if err != nil {
		return err
	}
	opt := transport.DecodeQueryOptions(ctx)
	result, total, err := c.rs.Filter(in, opt)
	if err != nil {
		return err
	}
	return transport.SendPaginationResponse(ctx, http.StatusOK, result, total)
}

// DeleteResturent deletes a resturent
//
//	@Summary		Delete a resturent
//	@Description	Delete a resturent by id
//	@Tags			Resturent
//	@ID				deleteResturent
//	@Accept			json
//	@Produce		json
//	@Security		JWT
//	@Param			Authorization	header		string	true	"Bearer "
//	@Param			id				path		string	true	"Resturent id"
//	@Success		204				{object}	domain.BaseResponse
//	@Failure		400				{object}	domain.InvalidRequestError
//	@Failure		401				{object}	domain.UnauthorizedError
//	@Failure		403				{object}	domain.ForbiddenAccessError
//	@Failure		404				{object}	domain.DataNotFoundError
//	@Failure		500				{object}	domain.SystemError
//	@Router			/resturent/{id} [delete]
func (c *ResturentController) DeleteResturent(ctx echo.Context) error {
	id, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		return err
	}
	err = c.rs.Delete(id)
	if err != nil {
		return err
	}
	return transport.SendResponse(ctx, http.StatusNoContent, nil)
}
