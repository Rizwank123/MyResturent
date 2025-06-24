package controller

import (
	"net/http"

	"github.com/gofrs/uuid/v5"
	"github.com/labstack/echo/v4"
	"github.com/rizwank123/myResturent/internal/domain"
	"github.com/rizwank123/myResturent/internal/http/transport"
)

type MeanuCardController struct {
	mcs domain.MenuCardService
}

func NewMenuCardController(mcs domain.MenuCardService) MeanuCardController {
	return MeanuCardController{
		mcs: mcs,
	}
}

// CreateMenuCard creates a new menu card
//
//	@Summary		Create a menu card
//	@Description	Create a menu card
//	@Tags			MenuCard
//	@ID				createMenuCard
//	@Accept			json
//	@Produce		json
//	@Security		JWT
//	@Param			Authorization	header		string						true	"Bearer "
//	@Param			in				body		domain.CreateMenuCardInput	true	"Payload"
//	@Success		201				{object}	domain.BaseResponse{data=domain.MenuCard}
//	@Failure		400				{object}	domain.InvalidRequestError
//	@Failure		401				{object}	domain.UnauthorizedError
//	@Failure		403				{object}	domain.ForbiddenAccessError
//	@Failure		500				{object}	domain.SystemError
//	@Router			/menu_card [post]
func (c *MeanuCardController) CreateMenuCard(ctx echo.Context) error {
	var in domain.CreateMenuCardInput
	err := transport.DecodeAndValidateRequestBody(ctx, &in)
	if err != nil {
		return err
	}
	result, err := c.mcs.Create(in)
	if err != nil {
		return err
	}
	return transport.SendResponse(ctx, http.StatusCreated, result)
}

// GetMenuCard retrieves a menu card
//
//	@Summary		Retrieve a menu card
//	@Description	Retrieve a menu card
//	@Tags			MenuCard
//	@ID				getMenuCard
//	@Accept			json
//	@Produce		json
//	@Security		JWT
//	@Param			Authorization	header		string		true	"Bearer "
//	@Param			id				path		uuid.UUID	true	"Menu card id"
//	@Success		200				{object}	domain.BaseResponse{data=domain.MenuCard}
//	@Failure		400				{object}	domain.InvalidRequestError
//	@Failure		401				{object}	domain.UnauthorizedError
//	@Failure		403				{object}	domain.ForbiddenAccessError
//	@Failure		404				{object}	domain.NotFoundError
//	@Failure		500				{object}	domain.SystemError
//	@Router			/menu_card/{id} [get]
func (c *MeanuCardController) GetMenuCard(ctx echo.Context) error {
	id, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		return err
	}
	result, err := c.mcs.FindById(id)
	if err != nil {
		return err
	}
	return transport.SendResponse(ctx, http.StatusOK, result)
}

// FilterMenuCards filters menu cards by criteria
//
//	@Summary		Filter menu cards
//	@Description	Filter menu cards by criteria
//	@Tags			MenuCard
//	@ID				filterMenuCards
//	@Accept			json
//	@Produce		json
//	@Param			page	query		int					false	"Page number"
//	@Param			size	query		int					false	"Page size"
//	@Param			in		body		domain.FilterInput	true	"Filter criteria"
//	@Success		200		{object}	domain.PaginationResponse{data=[]domain.MenuCard}
//	@Failure		400		{object}	domain.InvalidRequestError
//	@Failure		401		{object}	domain.UnauthorizedError
//	@Failure		403		{object}	domain.ForbiddenAccessError
//	@Failure		500		{object}	domain.SystemError
//	@Router			/menu_card/filter [post]
func (c *MeanuCardController) Filter(ctx echo.Context) error {
	var in domain.FilterInput
	err := transport.DecodeAndValidateRequestBody(ctx, &in)
	if err != nil {
		return err
	}
	opt := transport.DecodeQueryOptions(ctx)
	result, total, err := c.mcs.Filter(in, opt)
	if err != nil {
		return err
	}
	return transport.SendPaginationResponse(ctx, http.StatusOK, result, total)
}

// UpdateMenuCard updates a menu card
//
//	@Summary		Update a menu card
//	@Description	Update a menu card
//	@Tags			MenuCard
//	@ID				updateMenuCard
//	@Accept			json
//	@Produce		json
//	@Security		JWT
//	@Param			Authorization	header		string						true	"Bearer "
//	@Param			id				path		uuid.UUID					true	"Menu card id"
//	@Param			in				body		domain.UpdateMenuCardInput	true	"Payload"
//	@Success		200				{object}	domain.BaseResponse{data=domain.MenuCard}
//	@Failure		400				{object}	domain.InvalidRequestError
//	@Failure		401				{object}	domain.UnauthorizedError
//	@Failure		403				{object}	domain.ForbiddenAccessError
//	@Failure		404				{object}	domain.NotFoundError
//	@Failure		500				{object}	domain.SystemError
//	@Router			/menu_card/{id} [patch]
func (c *MeanuCardController) UpdateMenuCard(ctx echo.Context) error {
	var in domain.UpdateMenuCardInput
	err := transport.DecodeAndValidateRequestBody(ctx, &in)
	if err != nil {
		return err
	}
	id, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		return err
	}
	result, err := c.mcs.Update(id, in)
	if err != nil {
		return err
	}
	return transport.SendResponse(ctx, http.StatusOK, result)
}

// DeleteMenuCard deletes a menu card
//
//	@Summary		Delete a menu card
//	@Description	Delete a menu card
//	@Tags			MenuCard
//	@ID				deleteMenuCard
//	@Accept			json
//	@Produce		json
//	@Security		JWT
//	@Param			Authorization	header		string		true	"Bearer "
//	@Param			id				path		uuid.UUID	true	"Menu card id"
//	@Success		204				{object}	domain.BaseResponse{data=domain.MenuCard}
//	@Failure		400				{object}	domain.InvalidRequestError
//	@Failure		401				{object}	domain.UnauthorizedError
//	@Failure		403				{object}	domain.ForbiddenAccessError
//	@Failure		404				{object}	domain.NotFoundError
//	@Failure		500				{object}	domain.SystemError
//	@Router			/menu_card/{id} [delete]
func (c *MeanuCardController) DeleteMenuCard(ctx echo.Context) error {
	id, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		return err
	}
	err = c.mcs.Delete(id)
	if err != nil {
		return err
	}
	return transport.SendResponse(ctx, http.StatusNoContent, nil)
}
