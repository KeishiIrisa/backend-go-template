package services

import (
	"fmt"

	"github.com/KeishiIrisa/backend-go-template/internal/database"
	"github.com/KeishiIrisa/backend-go-template/internal/schemas"
	"github.com/KeishiIrisa/backend-go-template/internal/utils"

	"github.com/KeishiIrisa/backend-go-template/internal/models"

	"golang.org/x/crypto/bcrypt"
)

// CreateUser creates a new user in the database using the UserSignupSchemaIn struct
func CreateUser(input schemas.UserSignupSchemaIn) (*models.User, error) {
	// Check if the email is already registered
	var existingUser models.User
	if err := database.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		return nil, fmt.Errorf("email already exists")
	}
	fmt.Println("I am here")

	// Hash the password before saving
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error hashing password: %v", err)
	}

	// Create a new User model instance and set fields
	newUser := models.User{
		Email:        input.Email,
		FirstName:    input.FirstName,
		LastName:     input.LastName,
		PasswordHash: string(hashedPassword),
	}

	// Save the user to the database
	if err := database.DB.Create(&newUser).Error; err != nil {
		return nil, fmt.Errorf("error creating user: %v", err)
	}

	return &newUser, nil
}

func GetUserById(userId uint) (models.User, error) {
	var user models.User
	if err := database.DB.First(&user, userId).Error; err != nil {
		return user, err
	}
	return user, nil
}

func UpdateUserById(userId uint, input schemas.UserUpdateSchemaIn) (models.User, error) {
	var user models.User
	if err := database.DB.First(&user, userId).Error; err != nil {
		return user, err
	}
	if input.FirstName != "" {
		user.FirstName = input.FirstName
	}
	if input.LastName != "" {
		user.LastName = input.LastName
	}

	return user, database.DB.Save(&user).Error
}

func DeleteUserById(userId uint) error {
	return database.DB.Delete(&models.User{}, userId).Error
}

// AuthenticateUser checks user credentials and returns a JWT if successful
func AuthenticateUser(input schemas.UserLoginSchemaIn) (string, error) {
	var user models.User

	// Find the user in the database by email
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return "", fmt.Errorf("user not found")
	}

	// Compare the provided password with the stored hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		return "", fmt.Errorf("incorrect password")
	}

	// Generate JWT token if authentication is successful
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return "", fmt.Errorf("failed to generate token")
	}

	return token, nil
}
