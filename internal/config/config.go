package config

import (
	"fmt"
	"time"
)

const ApplicationLabel = "movie-crud-app"

type Config struct {
	Environment              string        `env:"ENVIRONMENT, default=development"`
	ServiceName              string        `env:"SERVICE_NAME"`
	ServerIP                 string        `env:"SERVER_IP"`
	LogLevel                 string        `env:"LOG_LEVEL, default=debug"`
	HTTPPort                 string        `env:"HTTP_PORT, default=:7900"`
	AllowedOrigins           string        `env:"CORS_ALLOWED_ORIGINS"`
	GracefulShutdownInterval time.Duration `env:"GRACEFUL_SHUTDOWN_INTERVAL,default=5s"`

	Postgres *Database `env:",prefix=POSTGRES_"`
}

type Database struct {
	Host                  string        `env:"HOST"`
	Port                  int           `env:"PORT"`
	User                  string        `env:"USER"`
	Password              string        `env:"PASSWORD"`
	Database              string        `env:"DATABASE"`
	SSLMode               string        `env:"SSLMODE,default=disable"`
	MaxIdleConnections    int32         `env:"MAX_IDLE_CONNECTIONS,default=25"`
	MaxOpenConnections    int           `env:"MAX_OPEN_CONNECTIONS,default=25"`
	ConnectionMaxLifetime time.Duration `env:"CONNECTION_MAX_LIFETIME,default=5m"`
}

func (c *Database) PostgresURL() string {
	if c.User == "" {
		return fmt.Sprintf(
			"host=%s port=%d  dbname=%s sslmode=%s",
			c.Host,
			c.Port,
			c.Database,
			c.SSLMode,
		)
	}

	if c.Password == "" {
		return fmt.Sprintf(
			"host=%s port=%d user=%s  dbname=%s sslmode=%s",
			c.Host,
			c.Port,
			c.User,
			c.Database,
			c.SSLMode,
		)
	}

	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Host,
		c.Port,
		c.User,
		c.Password,
		c.Database,
		c.SSLMode,
	)
}
