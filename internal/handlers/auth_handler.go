package handlers

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/<username>/<project>/internal/models"
    "github.com/<username>/<project>/internal/repository"
)

type AuthHandler struct {
    userRepo repository.UserRepository
}

func NewAuthHandler(userRepo repository.UserRepository) *AuthHandler {
    return &AuthHandler{userRepo}
}

func (h *AuthHandler) Register(c echo.Context) error {
    var user models.User
    if err := c.Bind(&user); err != nil {
        return c.JSON(http.StatusBadRequest, err)
    }
    if err := h.userRepo.CreateUser(user); err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }
    return c.JSON(http.StatusCreated, user)
}
