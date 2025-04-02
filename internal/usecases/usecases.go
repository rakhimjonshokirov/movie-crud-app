package usecases

import (
	"rakhimjonshokirov/movie-crud-app/internal/usecases/movies"
	"rakhimjonshokirov/movie-crud-app/internal/usecases/users"

	"go.uber.org/fx"
)

var Module = fx.Options(
	movies.Module,
	users.Module,
)
