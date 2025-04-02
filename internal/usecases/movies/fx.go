package movies

import (
	"rakhimjonshokirov/movie-crud-app/internal/repositories/movies"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

const moduleName = "usecases/movie"

var Module = fx.Module(
	moduleName,
	fx.Provide(
		func(movieRepo *movies.MovieRepository) movieRepo {
			return movieRepo
		},
		newUseCase,
	),
	fx.Decorate(func(logger *zap.Logger) *zap.Logger {
		return logger.With(zap.String("module", moduleName))
	}),
)
