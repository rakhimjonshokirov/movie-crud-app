package users

import "go.uber.org/fx"

const moduleName = "repositories/users"

var Module = fx.Module(
	moduleName,
	fx.Provide(
		NewUserRepository,
	),
)
