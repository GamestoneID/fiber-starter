package providers

import (
	"fmt"
	"gamestone/config"
	"gamestone/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewDBConnection establishes a connection to the database based on provided config and returns the DB instance
func NewDBConnection(cfg config.DatabaseConfig) (*gorm.DB, error) {
	// Initialize the DSN (Data Source Name) and Dialector based on the DB driver
	var dsn string
	var dialector gorm.Dialector

	switch cfg.DBDriver {
	case "postgres":
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
		dialector = postgres.Open(dsn)

	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
		dialector = mysql.Open(dsn)

	default:
		return nil, fmt.Errorf("unsupported database driver: %s", cfg.DBDriver)
	}

	// Open the DB connection
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return nil, err
	}

	// Run database migrations
	if err := db.AutoMigrate(
		&models.Game{},
	); err != nil {
		log.Fatalf("Failed to apply database migrations: %v", err)
		return nil, err
	}

	log.Println("Database connected and migrations applied successfully")
	return db, nil
}
