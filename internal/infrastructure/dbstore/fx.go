package dbstore

import (
	"rakhimjonshokirov/movie-crud-app/internal/config"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const moduleName = "db"

var Module = fx.Module(
	moduleName,
	fx.Provide(
		func(db *DB) *gorm.DB {
			return db.db
		},
		NewDB,
	),
)

type Params struct {
	fx.In
	Logger *zap.Logger
	Config config.Config
}
