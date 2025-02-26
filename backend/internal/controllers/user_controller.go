package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/KeishiIrisa/backend-go-template/internal/schemas"
	"github.com/KeishiIrisa/backend-go-template/internal/services"

	"github.com/gin-gonic/gin"
)

// RegisterUser godoc
// @Summary Register a new user
// @Description Create a new user by providing email and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body schemas.UserSignupSchemaIn true "User Data"
// @Success 201 {object} models.User
// @Failure 400 {string} string "Bad request"
// @Router /sign-up [post]
func RegisterUser(c *gin.Context) {
	var input schemas.UserSignupSchemaIn

	// Bind the JSON input to the UserSignupSchemaIn DTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Call the service to create the user, passing the entire input schema
	user, err := services.CreateUser(input)
	if err != nil {
		if err.Error() == "email already exists" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusCreated, user)
}

// LoginUser godoc
// @Summary Log in a user
// @Description Log in by providing email and password to receive a JWT
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body schemas.UserLoginSchemaIn true "User Login Data"
// @Success 200 {string} string "JWT Token"
// @Failure 401 {string} string "Unauthorized"
// @Router /login [post]
func LoginUser(c *gin.Context) {
	var input schemas.UserLoginSchemaIn

	// Bind the JSON input to the UserLoginSchemaIn DTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Authenticate the user and generate JWT token using the service
	token, err := services.AuthenticateUser(input)
	if err != nil {
		if err.Error() == "user not found" || err.Error() == "incorrect password" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// GetLoggedInUser godoc
// @Summary
// @Description
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Router /users/me [get]
func GetLoggedInUser(c *gin.Context) {
	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	user, err := services.GetUserById(userId.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetUserById godoc
// @Summary
// @Description
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Security BearerAuth
// @Router /users/{id} [get]
func GetUserById(c *gin.Context) {
	idStr := c.Param("id")
	userId, err := strconv.Atoi(idStr)
	fmt.Printf("userId: %v", userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	user, err := services.GetUserById(uint(userId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// UpdateUser godoc
// @Summary
// @Description
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param item body schemas.UserUpdateSchemaIn true "UserUpdateSchemaIn"
// @Security BearerAuth
// @Router /users/{id} [put]
func UpdateUser(c *gin.Context) {
	var input schemas.UserUpdateSchemaIn
	currentUserId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userIdStr := c.Param("id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	// ログインしたユーザーと変更対象のユーザーが同じかどうかを検証
	if currentUserId != uint(userId) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}
	user, err := services.UpdateUserById(uint(userId), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary
// @Description
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Security BearerAuth
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	currentUserId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userIdStr := c.Param("id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	if currentUserId != uint(userId) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		return
	}

	if err := services.DeleteUserById(uint(userId)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
