package infrastructure

import (
	"rakhimjonshokirov/movie-crud-app/internal/infrastructure/cache"
	"rakhimjonshokirov/movie-crud-app/internal/infrastructure/dbstore"

	"go.uber.org/fx"
)

var Module = fx.Options(
	cache.Module,
	dbstore.Module,
)
