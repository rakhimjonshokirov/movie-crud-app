package repositories

import (
	"rakhimjonshokirov/movie-crud-app/internal/repositories/movies"
	"rakhimjonshokirov/movie-crud-app/internal/repositories/users"

	"go.uber.org/fx"
)

var Module = fx.Options(
	movies.Module,
	users.Module,
)
