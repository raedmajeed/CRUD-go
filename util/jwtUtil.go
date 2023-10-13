package util

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JwtUtil struct {}

type Claims struct {
	Username string
	Role string
	*jwt.StandardClaims
}

func (*JwtUtil) CreateToken(username string, role string) (error, string) {
	claims := &Claims{
		Username: username,
		Role: role,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	strToken, err := token.SignedString([]byte("secret"))

	if err != nil {
		log.Println("ERROR CREATING JWT TOKEN")
		return errors.New("ERROR CREATING JWT TOKEN"), ""
	}

	return nil, strToken
}

func (j *JwtUtil) ValidateToken(role string) gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString := c.GetHeader("Authorization")
		tokenString = string([]byte(tokenString)[7:])
		if tokenString == "" {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
						"error": "Unauthorized",
				})
				return
		}
		claims := &Claims{}
		parsedToken, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if err != nil {
			log.Println("ERROR PARSING JWT TOKEN")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
		})
			return
		}

		fmt.Println(claims.Role)

		if claims.Role != role || !parsedToken.Valid {
			log.Println("USER NOT AUTHORIZED TO ACCESS THE URL")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			return
		}
		c.Next()
	}
}

func NewJwtUtil() *JwtUtil {
	return &JwtUtil{}
}