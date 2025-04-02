package movies

import "go.uber.org/fx"

const moduleName = "repositories/movie"

var Module = fx.Module(
	moduleName,
	fx.Provide(
		NewMovieRepository,
	),
)
