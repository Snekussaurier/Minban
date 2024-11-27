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

	context.SetCookie("auth_token", token, 60*60*24, "/", "", false, true)

	context.Status(http.StatusNoContent)
}

func CreateDefaultUser() {
	var existingUser database.User
	var userId = uuid.New().String()

	result := database.DB.Where("name = ?", os.Getenv("USER_NAME")).First(&existingUser)

	if result.Error == gorm.ErrRecordNotFound {
		newUser := database.User{
			ID:       userId,
			Name:     os.Getenv("USER_NAME"),
			Password: hashPassword(os.Getenv("USER_PASSWORD")),
		}

		if err := database.DB.Create(&newUser).Error; err != nil {
			log.Printf("Error creating default user: %v", err)
		}
	} else if result.Error != nil {
		log.Printf("Error checking for existing user: %v", result.Error)
	}

	initializeStates(userId)
	initializeTags(userId)
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

func initializeStates(userId string) {
	states := []database.State{
		{ID: 1, Name: "Todo", Position: 1, Color: "FF0000", UserID: userId},
		{ID: 2, Name: "In Progress", Position: 2, Color: "00FF00", UserID: userId},
		{ID: 3, Name: "Done", Position: 3, Color: "0000FF", UserID: userId},
	}

	for _, state := range states {
		if err := database.DB.Where("id = ?", state.ID).FirstOrCreate(&state).Error; err != nil {
			log.Fatalf("Failed to create state: %v", err)
		}
	}
}

func initializeTags(userId string) {
	tags := []database.Tag{
		{Name: "Feature", Color: "FF0000", UserID: userId},
		{Name: "Bug", Color: "00FF00", UserID: userId},
	}

	for _, tag := range tags {
		if err := database.DB.Where("name = ?", tag.Name).FirstOrCreate(&tag).Error; err != nil {
			log.Fatalf("Failed to create tag: %v", err)
		}
	}
}
