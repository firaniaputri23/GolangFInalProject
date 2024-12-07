package middleware

import (
    "github.com/labstack/echo/v4"
    "net/http"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        token := c.Request().Header.Get("Authorization")
        if token == "" {
            return c.JSON(http.StatusUnauthorized, "Unauthorized")
        }
        // Verifikasi token di sini
        return next(c)
    }
}
