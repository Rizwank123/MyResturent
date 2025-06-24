package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type HealthController struct{}

func NewHealthController() *HealthController {
	return &HealthController{}
}

// GET /health
// @Summary		Find service health status
// @Description	Find service health status
// @Tags			Health
// @ID				healthCheck
// @Accept			json
// @Produce		json
// @Success		200	{object}	map[string]string
// @Failure		500	{object}	domain.SystemError
// @Router			/health [get]
func (hc *HealthController) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status": "ok",
	})
}

// GET /metrics
// @Summary		Find service metrics
// @Description	Find service metrics
// @Tags			Health
// @ID				metrics
// @Accept			json
// @Produce		json
// @Success		200	{object}	prometheus metrics
// @Failure		500	{object}	domain.SystemError
// @Router			/metrics [get]
func (hc *HealthController) Metrics(c echo.Context) error {
	promhttp.Handler().ServeHTTP(c.Response(), c.Request())
	return nil
}
