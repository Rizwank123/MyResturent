package api

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/rizwank123/myResturent/internal/http/controller"
	"github.com/rizwank123/myResturent/internal/pkg/config"
)

type ResturnetApi struct {
	cfg                 config.ResturantConfig
	MenuCardController  controller.MeanuCardController
	ResturentController controller.ResturentController
	RatingController    controller.RatingController
	UserController      controller.UserController
}

// NewResturnetApi creates a new ResturnetApi instance
//
//	@title						Resturnet API
//	@version					1.0
//	@description				Resturnet's set of APIs
//	@termsOfService				http://example.com/terms/
//	@contact.name				API Support
//	@contact.url				https://rizwank123.github.io
//	@contact.email				rizwank431@gmail.com
//	@host						localhost:7700
//	@BasePath					/api/v1
//	@schemes					http
//	@securityDefinitions.apiKey	JWT
//	@in							header
//
//	@name						Authorization

func NewResturnetApi(cfg config.ResturantConfig, mcs controller.MeanuCardController, rc controller.ResturentController, rtc controller.RatingController, uc controller.UserController) *ResturnetApi {
	return &ResturnetApi{
		cfg:                 cfg,
		MenuCardController:  mcs,
		ResturentController: rc,
		RatingController:    rtc,
		UserController:      uc,
	}
}
func (r *ResturnetApi) SetupRoutes(e *echo.Echo) {
	apiV1 := e.Group("/api/v1")
	auth := echojwt.JWT([]byte(r.cfg.AuthSecret))
	// menu card api
	menuCardApi := apiV1.Group("/menu-card")
	menuCardSecureApi := apiV1.Group("/menu-card")
	menuCardSecureApi.Use(auth)
	menuCardSecureApi.POST("", r.MenuCardController.CreateMenuCard)
	menuCardApi.GET("/:id", r.MenuCardController.GetMenuCard)
	menuCardApi.POST("/filter", r.MenuCardController.Filter)
	menuCardSecureApi.PATCH("/:id", r.MenuCardController.UpdateMenuCard)
	menuCardSecureApi.DELETE("/:id", r.MenuCardController.DeleteMenuCard)

	// resturent api
	resturentApi := apiV1.Group("/resturent")
	resturentSecureApi := apiV1.Group("/resturent")
	resturentSecureApi.Use(auth)
	resturentSecureApi.POST("", r.ResturentController.CreateResturent)
	resturentApi.GET("/:id", r.ResturentController.FindById)
	resturentApi.POST("/filter", r.ResturentController.Filter)
	resturentSecureApi.PUT("/:id", r.ResturentController.UpdateResturent)
	resturentSecureApi.DELETE("/:id", r.ResturentController.DeleteResturent)

	// rating api
	ratingApi := apiV1.Group("/rating")
	ratingApi.POST("", r.RatingController.CreateRating)
	ratingApi.POST("/filter", r.RatingController.Filter)
	ratingApi.PATCH("/:id", r.RatingController.UpdateRating)
	ratingApi.DELETE("/:id", r.RatingController.DeleteRating)

	// user api
	userApi := apiV1.Group("/user")
	userSecureApi := apiV1.Group("/user")
	userSecureApi.Use(auth)
	userApi.POST("", r.UserController.Register)
	userSecureApi.GET("/:id", r.UserController.FindById)
	userSecureApi.POST("/filter", r.UserController.Filter)
	userSecureApi.PATCH("/:id", r.UserController.UpdateUser)
	userApi.POST("/login", r.UserController.Login)

}
