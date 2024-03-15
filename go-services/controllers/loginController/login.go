package loginController

import (
	"time"

	"example.com/web-service-gin/models"
	"example.com/web-service-gin/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("your_secret_key") // Replace with your secret key

func Login(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	var existingUser models.User
	if err := models.DB.Where("email = ?", user.Email).First(&existingUser).Error; err != nil {
		c.JSON(400, gin.H{"error": "User not found"})
		return
	}

	if !utils.CompareHashPassword(user.Password, existingUser.Password) {
		c.JSON(400, gin.H{"error": "Invalid password"})
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour) // Set token expiration to 24 hours

	claims := &models.Claims{
		Role: existingUser.Role,
		StandardClaims: jwt.StandardClaims{
			Subject:   existingUser.Email,
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not generate token"})
		return
	}

	// Set token as HTTPOnly cookie
	c.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "", false, true)

	c.JSON(200, gin.H{"message": "User logged in successfully"})
}
