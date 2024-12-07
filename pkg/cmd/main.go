package main

import (
    "github.com/labstack/echo/v4"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/<username>/<project>/internal/handlers"
    "github.com/<username>/<project>/internal/repository"
)

func main() {
    db, err := gorm.Open(postgres.Open("host=localhost user=youruser dbname=yourdb sslmode=disable"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    userRepo := repository.NewUserRepository(db)
    authHandler := handlers.NewAuthHandler(userRepo)

    e := echo.New()

    e.POST("/register", authHandler.Register)
    // Tambahkan rute lainnya

    e.Logger.Fatal(e.Start(":8080"))
}
