package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgconn"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/rizwank123/myResturent/internal/domain"
	"github.com/rizwank123/myResturent/internal/http/swagger"
	"github.com/rizwank123/myResturent/internal/http/transport"
)

func (b *ResturnetApi) SetupMiddleware(e *echo.Echo) {
	vv10 := validator.New()
	vv10.RegisterValidation("trim", func(fl validator.FieldLevel) bool {
		return len(strings.TrimSpace(fl.Field().String())) != 0
	})

	e.Validator = &transport.CustomValidator{Validator: vv10}

	// Set Up error handler middlerware
	e.HTTPErrorHandler = errorMiddleware

	// Set Up Request Body limit to 10MB
	e.Use(echomiddleware.BodyLimit("10M"))
	// Recovery Middleware recovers from panics
	e.Use(echomiddleware.Recover())

	// Add Request ID Middleware
	e.Use(echomiddleware.RequestID())
	// Add Cors Middleware
	e.Use(echomiddleware.CORS())
	// Add Swagger Middleware
	e.Use(swagger.RedirectSwagger)

	// Add logger Middleware
	e.Use(echomiddleware.LoggerWithConfig(echomiddleware.LoggerConfig{
		Format: "${time_rfc3339} ${id} ${remote_ip} ${method} ${uri} ${latency_human} ${status} ${error}\n",
	}))

}

// errorMiddleware custom error handler for echo
func errorMiddleware(err error, c echo.Context) {
	switch err.(type) {
	case *echo.HTTPError:
		err := err.(*echo.HTTPError)
		switch err.Code {
		case http.StatusUnauthorized:
			_ = c.JSON(err.Code, domain.UnauthorizedError{
				Code:    domain.ErrorCodeUNAUTHORIZED,
				Message: err.Message.(string),
			})
		case http.StatusForbidden:
			_ = c.JSON(err.Code, domain.ForbiddenAccessError{
				Code:    domain.ErrorCodeFORBIDDEN_ACCESS,
				Message: err.Message.(string),
			})
		case http.StatusNotFound:
			_ = c.JSON(err.Code, domain.NotFoundError{})
		case http.StatusBadRequest:
			_ = c.JSON(err.Code, domain.InvalidRequestError{Message: err.Message.(string)})
		default:
			_ = c.JSON(err.Code, domain.SystemError{Code: domain.ErrorCodeINTERNAL_SERVER_ERROR, Message: err.Message.(string)})
		}

	case validator.ValidationErrors:
		var ve error
		fields := make([]string, 0)
		errs := err.(validator.ValidationErrors)

		for _, e := range errs {
			if e.Tag() == "required" {

				fields = append(fields, fmt.Sprintf("%s is required", e.Field()))
				continue
			}

			if e.Tag() == "e164" {
				fields = append(fields, fmt.Sprintf("%s is an invalid mobile number", e.Field()))
				continue
			}

			if e.Tag() == "email" {
				fields = append(fields, fmt.Sprintf("%s is an invalid email address", e.Field()))
				continue
			}

			if e.Tag() == "oneof" {
				fields = append(fields, fmt.Sprintf("%s must be one of %s", e.Field(), e.Param()))
				continue
			}

		}

		ve = domain.ValidationError{
			Code:    domain.ErrorCodeVALIDATION_ERROR,
			Message: domain.MessageVALIDATIONFAILED,
			Fields:  fields,
		}
		_ = c.JSON(http.StatusBadRequest, ve)

	case *pgconn.PgError:
		res := domain.SystemError{
			Code:    domain.ErrorCodeINTERNAL_SERVER_ERROR,
			Message: err.Error(),
		}
		_ = c.JSON(http.StatusInternalServerError, res)

	case domain.DataNotFoundError:
		res := domain.UserError{
			Code:    domain.ErrorCodeINVALID_REQUEST,
			Message: err.Error(),
		}
		_ = c.JSON(http.StatusBadRequest, res)

	case domain.UserError:
		usrErr := err.(domain.UserError)
		res := domain.UserError{
			Code:    usrErr.Code,
			Message: usrErr.Message,
		}
		_ = c.JSON(http.StatusBadRequest, res)

	case domain.UnauthorizedError:
		res := domain.UnauthorizedError{
			Code:    domain.ErrorCodeUNAUTHORIZED,
			Message: domain.MessageUNAUTHORIZEDACCESS,
		}
		_ = c.JSON(http.StatusUnauthorized, res)

	case domain.ForbiddenAccessError:
		res := domain.ForbiddenAccessError{
			Code:    domain.ErrorCodeFORBIDDEN_ACCESS,
			Message: domain.MessageFORBIDDENACCESS,
		}
		_ = c.JSON(http.StatusForbidden, res)

	default:
		res := domain.SystemError{
			Code:    domain.ErrorCodeINTERNAL_SERVER_ERROR,
			Message: err.Error(),
		}
		_ = c.JSON(http.StatusInternalServerError, res)
	}
}
