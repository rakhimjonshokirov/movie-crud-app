package users

import (
	"rakhimjonshokirov/movie-crud-app/internal/repositories/users"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

const moduleName = "usecases/users"

var Module = fx.Module(
	moduleName,
	fx.Provide(
		func(usersRepo *users.UserRepository) UserRepo {
			return usersRepo
		},
		NewUserUseCase,
	),
	fx.Decorate(func(logger *zap.Logger) *zap.Logger {
		return logger.With(zap.String("module", moduleName))
	}),
)
