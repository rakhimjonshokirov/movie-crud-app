package rest

import (
	"context"
	"rakhimjonshokirov/movie-crud-app/internal/entities"
	"rakhimjonshokirov/movie-crud-app/internal/usecases/movies"
	"rakhimjonshokirov/movie-crud-app/internal/usecases/users"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

const moduleName = "rest/setup"

// movieGetter is an interface
type movieUscInterface interface {
	GetMovieByID(ctx context.Context, id uint) (*entities.Movie, error)
	GetAllMovies(ctx context.Context) ([]entities.Movie, error)
	CreateMovie(ctx context.Context, movie *entities.Movie) error
	UpdateMovie(ctx context.Context, movie *entities.Movie) error
	DeleteMovie(ctx context.Context, id uint) error
}

// userInterface is an interface
type usersUscInterface interface {
	Register(email, username, password string) error
	Login(email, password string) (string, error)
	ValidateToken(token string) (uint, error)
}

var Module = fx.Module(
	moduleName,
	fx.Provide(
		func(uc *movies.UseCase) movieUscInterface {
			return uc
		},
		func(uc *users.UserUseCase) usersUscInterface {
			return uc
		},
		newHandler,
	),
	fx.Decorate(func(logger *zap.Logger) *zap.Logger {
		return logger.With(zap.String("handlers", moduleName))
	}),
)
