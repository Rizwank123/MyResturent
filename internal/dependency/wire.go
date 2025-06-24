//go:build wireinject

package dependency

import (
	"github.com/google/wire"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rizwank123/myResturent/internal/database"
	"github.com/rizwank123/myResturent/internal/http/api"
	"github.com/rizwank123/myResturent/internal/http/controller"
	"github.com/rizwank123/myResturent/internal/pkg/config"
	"github.com/rizwank123/myResturent/internal/pkg/security"
	"github.com/rizwank123/myResturent/internal/repository"
	"github.com/rizwank123/myResturent/internal/service"
)

// NewConfig returns ResturantConfig
func NewConfig(options config.Options) (config.ResturantConfig, error) {
	wire.Build(config.NewResturantConfig)
	return config.ResturantConfig{}, nil
}

// NewDatabase returns a new database connection pool
func NewDatabase(cfg config.ResturantConfig) (*pgxpool.Pool, error) {
	wire.Build(database.NewDB)
	return &pgxpool.Pool{}, nil
}

func NewResturnetApi(cfg config.ResturantConfig, db *pgxpool.Pool) (*api.ResturnetApi, error) {
	wire.Build(
		security.NewJWTSecurityManager,
		repository.NewTransactioner,
		repository.NewMenuCardRepository,
		repository.NewResturentRepository,
		repository.NewRatingRepository,
		repository.NewUserRepository,

		service.NewMenuCardService,
		service.NewResturentService,
		service.NewRatingService,
		service.NewUserService,

		controller.NewMenuCardController,
		controller.NewResturentController,
		controller.NewRatingController,
		controller.NewUserController,

		api.NewResturnetApi,
	)
	return &api.ResturnetApi{}, nil
}
