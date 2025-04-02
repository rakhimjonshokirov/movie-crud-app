package config

import (
	"context"

	"github.com/sethvargo/go-envconfig"
	"go.uber.org/fx"
)

const moduleName = "config"

var Module = fx.Module(
	moduleName,
	fx.Provide(newConfig),
)

func newConfig(ctx context.Context) Config {
	var cfg Config
	if err := envconfig.Process(ctx, &cfg); err != nil {
		panic(err)
	}
	return cfg
}
