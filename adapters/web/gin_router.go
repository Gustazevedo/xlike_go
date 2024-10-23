package web

import (
    "github.com/gin-gonic/gin"
    "github.com/Gustazevedo/xlike_go/tree/develop/domain/service"
    "net/http"
)

type UserHandler struct {
    userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
    return &UserHandler{
        userService: userService,
    }
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
    var input struct {
        Username string `json:"username"`
        Email    string `json:"email"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    user, err := h.userService.RegisterUser(input.Username, input.Email)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User registered successfully", "user": user})
}
