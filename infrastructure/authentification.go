package infrastructure

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/RomaneCAVEY/FeatureFlag-Manager/domain/entities"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type SignedDetails struct {
	User entities.User `json:"user"`
	jwt.RegisteredClaims
}

func GetHMACSecret() string {
	MINIMUM_LENGTH := 12
	privateKey := os.Getenv("CONFIG_JWT_PRIVATE_KEY")
	if len(privateKey) < MINIMUM_LENGTH {
		panic(fmt.Sprintf("Invalid CONFIG_JWT_PRIVATE_KEY environment variable. Min length should be %v characters", MINIMUM_LENGTH))
	}
	return privateKey
}

func ValidateToken(tokenString string, privateKey string) (entities.User, error) {
	hmacSampleSecret := []byte(privateKey)
	token, err := jwt.ParseWithClaims(tokenString, &SignedDetails{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSampleSecret, nil
	})
	if err != nil {
		return entities.User{}, err
	}

	claims, tokenValid := token.Claims.(*SignedDetails)
	if tokenValid {
		return claims.User, nil
	}

	return entities.User{}, err
}

func Authenticate(privateKey string) gin.HandlerFunc {

	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("Authorization")
		if clientToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("No Authorization Header Provided")})
			c.Abort()
			return
		}
		bearerToken, found := strings.CutPrefix(clientToken, "Bearer ")
		if !found {
			c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Bearer token not found")})
			c.Abort()
			return
		}
		User, err := ValidateToken(bearerToken, privateKey)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Set("user", User)
		c.Next()
	}
}
