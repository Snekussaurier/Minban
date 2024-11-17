package controller

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/snekussaurier/minban-backend/database"
	"github.com/snekussaurier/minban-backend/mod"
	"gorm.io/gorm"
)

var jwtSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func Login(context *gin.Context) {

	var request = mod.LoginRequest{}

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, mod.ErrorResponse{Error: "Invalid input"})
		return
	}

	var user database.User
	result := database.DB.Where("name = ?", request.Username).First(&user)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, mod.ErrorResponse{Error: "Invalid username or password"})
		return
	}

	if hashPassword(request.Password) != user.Password {
		context.JSON(http.StatusBadRequest, mod.ErrorResponse{Error: "Invalid username or password"})
		return
	}

	token, err := generateJWTToken(user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, mod.ErrorResponse{Error: "Failed to generate token"})
		return
	}

	context.JSON(http.StatusOK, mod.LoginResponse{
		Token:  token,
		UserId: user.ID,
	})
}

func CreateDefaultUser() {
	var existingUser database.User

	result := database.DB.Where("name = ?", os.Getenv("USER_NAME")).First(&existingUser)

	if result.Error == gorm.ErrRecordNotFound {
		newUser := database.User{
			ID:       uuid.New().String(),
			Name:     os.Getenv("USER_NAME"),
			Password: hashPassword(os.Getenv("USER_PASSWORD")),
		}

		if err := database.DB.Create(&newUser).Error; err != nil {
			log.Printf("Error creating default user: %v", err)
		}
	} else if result.Error != nil {
		log.Printf("Error checking for existing user: %v", result.Error)
	}
}

func hashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

func generateJWTToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(), // Token expiration: 24 hours
		"iat":     time.Now().Unix(),                     // Issued at
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecretKey)
}
