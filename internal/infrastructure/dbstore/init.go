package dbstore

import (
	"log"
	"rakhimjonshokirov/movie-crud-app/internal/config"
	"rakhimjonshokirov/movie-crud-app/internal/entities"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Database struct
type DB struct {
	logger *zap.Logger
	config config.Config

	db *gorm.DB
}

func NewDB(params Params) *DB {
	dbNew := &DB{
		logger: params.Logger,
		config: params.Config,
	}

	initialized, err := dbNew.Init()
	if err != nil {
		log.Fatalf("❌ Failed to init DB: %v", err)
	}
	return initialized
}

// NewDatabase creates a new database connection
func (d *DB) Init() (*DB, error) {
	// Connect to database
	db, err := gorm.Open(postgres.Open(d.config.Postgres.PostgresURL()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil, err
	}

	// Run all needed migrations here
	if err := db.AutoMigrate(&entities.Movie{}); err != nil {
		log.Fatalf("❌ AutoMigrate failed: %v", err)
	}

	log.Println("Database connection established.")

	// Return database
	return &DB{
		db: db,
	}, nil
}
