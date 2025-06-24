package controller

import (
	"net/http"

	"github.com/gofrs/uuid/v5"
	"github.com/labstack/echo/v4"
	"github.com/rizwank123/myResturent/internal/domain"
	"github.com/rizwank123/myResturent/internal/http/transport"
)

type UserController struct {
	us domain.UserService
}

func NewUserController(us domain.UserService) UserController {
	return UserController{
		us: us,
	}
}

// RegisterUser registers a new user
//
//	@Summary		Registers a new user
//	@Description	Registers a new user
//	@Tags			User
//	@ID				registerUser
//	@Accept			json
//	@Produce		json
//	@Param			in	body		domain.CreateUserInput	true	"Payload"
//	@Success		201	{object}	domain.BaseResponse{data=domain.User}
//	@Failure		400	{object}	domain.InvalidRequestError
//	@Failure		401	{object}	domain.UnauthorizedError
//	@Failure		403	{object}	domain.ForbiddenAccessError
//	@Failure		500	{object}	domain.SystemError
//	@Router			/user [post]
func (c *UserController) Register(ctx echo.Context) error {
	var in domain.CreateUserInput

	err := transport.DecodeAndValidateRequestBody(ctx, &in)
	if err != nil {
		return err
	}
	result, err := c.us.Create(in)
	if err != nil {
		return err
	}
	return transport.SendResponse(ctx, http.StatusCreated, result)
}

// Login logs in a user
//
//	@Summary		Logs in a user
//	@Description	Logs in a user
//	@Tags			User
//	@ID				loginUser
//	@Accept			json
//	@Produce		json
//	@Param			in	body		domain.LoginUserInput	true	"Payload"
//	@Success		200	{object}	domain.LoginResponse
//	@Failure		400	{object}	domain.InvalidRequestError
//	@Failure		401	{object}	domain.UnauthorizedError
//	@Failure		500	{object}	domain.SystemError
//	@Router			/user/login [post]
func (c *UserController) Login(ctx echo.Context) error {
	var in domain.LoginUserInput
	err := transport.DecodeAndValidateRequestBody(ctx, &in)
	if err != nil {
		return err
	}
	result, err := c.us.Login(in)
	if err != nil {
		return err
	}
	return transport.SendResponse(ctx, http.StatusOK, result)
}

// FindUserById Find user by id
//
//	@Summary		Find user by id
//	@Description	Find user by id
//	@Tags			User
//	@ID				findUserById
//	@Accept			json
//	@Produce		json
//	@Security		JWT
//	@Param			Authorization	header		string	true	"Bearer "
//	@Param			id				path		string	true	"User id"
//	@Success		200				{object}	domain.BaseResponse{data=domain.User}
//	@Failure		400				{object}	domain.InvalidRequestError
//	@Failure		401				{object}	domain.UnauthorizedError
//	@Failure		403				{object}	domain.ForbiddenAccessError
//	@Failure		404				{object}	domain.NotFoundError
//	@Failure		500				{object}	domain.SystemError
//	@Router			/user/{id} [get]
func (c *UserController) FindById(ctx echo.Context) error {
	id, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		return err
	}
	result, err := c.us.FindById(id)
	if err != nil {
		return err
	}
	return transport.SendResponse(ctx, http.StatusOK, result)
}

// FilterUser filters user by criteria
//
//	@Summary		Filter user
//	@Description	Filter user by criteria
//	@Tags			User
//	@ID				filterUser
//	@Accept			json
//	@Produce		json
//	@Security		JWT
//	@Param			Authorization	header		string				true	"Bearer "
//	@Param			page			query		int					false	"Page number"
//	@Param			size			query		int					false	"Page size"
//	@Param			in				body		domain.FilterInput	true	"Filter criteria"
//	@Success		200				{object}	domain.PaginationResponse{data=[]domain.User}
//	@Failure		400				{object}	domain.InvalidRequestError
//	@Failure		401				{object}	domain.UnauthorizedError
//	@Failure		403				{object}	domain.ForbiddenAccessError
//	@Failure		404				{object}	domain.NotFoundError
//	@Failure		500				{object}	domain.SystemError
//	@Router			/user/filter [post]
func (c *UserController) Filter(ctx echo.Context) error {
	var in domain.FilterInput
	err := transport.DecodeAndValidateRequestBody(ctx, &in)
	if err != nil {
		return err
	}
	opt := transport.DecodeQueryOptions(ctx)
	result, total, err := c.us.Filter(in, opt)
	if err != nil {
		return err
	}
	return transport.SendPaginationResponse(ctx, http.StatusOK, result, total)
}

// UpdateUser updates a user
//
//	@Summary		Updates a user
//	@Description	Updates a user
//	@Tags			User
//	@ID				updateUser
//	@Accept			json
//	@Produce		json
//	@Security		JWT
//	@Param			Authorization	header		string					true	"Bearer "
//	@Param			id				path		string					true	"User id"
//	@Param			in				body		domain.UpdateUserInput	true	"Payload"
//	@Success		200				{object}	domain.BaseResponse{data=domain.User}
//	@Failure		400				{object}	domain.InvalidRequestError
//	@Failure		401				{object}	domain.UnauthorizedError
//	@Failure		403				{object}	domain.ForbiddenAccessError
//	@Failure		404				{object}	domain.DataNotFoundError
//	@Failure		500				{object}	domain.SystemError
//	@Router			/user/{id} [put]
func (c *UserController) UpdateUser(ctx echo.Context) error {
	id, err := uuid.FromString(ctx.Param("id"))
	if err != nil {
		return err
	}
	var in domain.UpdateUserInput
	err = transport.DecodeAndValidateRequestBody(ctx, &in)
	if err != nil {
		return err
	}
	result, err := c.us.Update(id, in)
	if err != nil {
		return err
	}
	return transport.SendResponse(ctx, http.StatusOK, result)
}
