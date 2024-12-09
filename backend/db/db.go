package db

import (
	"fmt"
	"os"
	// "strings"

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
        return nil, fmt.Errorf("error loading environment variables: %v", err)
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
        return nil, fmt.Errorf("error connecting to database server: %v", err)
    }

    var dbExists bool
    checkDbQuery := fmt.Sprintf("SELECT 1 FROM pg_database WHERE datname = '%s';", dbName)
    if err := db.Raw(checkDbQuery).Scan(&dbExists).Error; err != nil {
        return nil, fmt.Errorf("error checking database existence: %v", err)
    }

    if !dbExists {
        createDbCmd := fmt.Sprintf("CREATE DATABASE \"%s\";", dbName)
        if res := db.Exec(createDbCmd); res.Error != nil {
            return nil, fmt.Errorf("error creating database: %v", res.Error)
        }
        fmt.Printf("Database \"%s\" created successfully.\n", dbName)
    } else {
        fmt.Printf("Database \"%s\" already exists.\n", dbName)
    }

    sqlDB, err := db.DB()
    if err != nil {
        return nil, err
    }
    sqlDB.Close()

    dsnWithDB := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable",
        dbHost, dbUsn, dbPass, dbPort, dbName)
    dbWithName, err := gorm.Open(postgres.Open(dsnWithDB), &gorm.Config{})
    if err != nil {
        return nil, fmt.Errorf("error connecting to database \"%s\": %v", dbName, err)
    }

    return &Database{db: dbWithName}, nil
}


func (d *Database) GetDB() *gorm.DB {
	return d.db
}
