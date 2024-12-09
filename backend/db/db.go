package db

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func NewDatabase() (*Database, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("Error loading environment variables: %v", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbUsn := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=disable",
		dbHost, dbUsn, dbPass, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Create database
	createDbCmd := fmt.Sprintf("CREATE DATABASE \"%s\";", dbName)
	if res := db.Exec(createDbCmd); res.Error != nil {
		if !strings.Contains(res.Error.Error(), "already exists") {
			return nil, fmt.Errorf("Error creating database: %v", res.Error)
		}
	}

	// Close first connection
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.Close()

	// Connect with database name
	dsnWithDB := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable",
		dbHost, dbUsn, dbPass, dbPort, dbName)
	dbWithName, err := gorm.Open(postgres.Open(dsnWithDB), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Database{db: dbWithName}, nil
}

func (d *Database) GetDB() *gorm.DB {
	return d.db
}
