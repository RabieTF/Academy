package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var (
	secretKey = []byte(os.Getenv("SECRET_TOKEN"))
)

// Middleware that checks if user is connected by validating his JWT token.
// Returns 500 if something went wrong.
// Returns 401 if user isn't authentified.
// Moves on to the next handler is user is authentified.
func VerifyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		prefix := "Bearer "
		header := c.GetHeader("Authorization")

		if header == "" {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "You are not authentified and therefore cannot perform this operation."})
			return
		}

		tokenString := strings.TrimPrefix(header, prefix)

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
			}

			return secretKey, nil
		})

		if err != nil {
			fmt.Println(err)
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Something went wrong, please contact your admin."})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println("claims: ", claims["userID"])
			c.Set("userID", claims["userID"])
			c.Next()
		} else {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "You are not authentified and therefore cannot perform this operation."})
			return
		}
		return
	}
}
