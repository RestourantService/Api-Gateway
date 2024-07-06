package api

import (
	"api-gateway/api/handler"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("secret-key")

func New(server *handler.Server) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	return router
}

func VerifyJWTMiddleware(ctx *gin.Context) {
	tokenStr := ctx.GetHeader("Authorization")

	if !strings.HasPrefix(tokenStr, "Bearer ") {
		ctx.IndentedJSON(401, gin.H{"error": "unauthorized"})
		ctx.Abort()
		return
	}

	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		ctx.IndentedJSON(401, gin.H{"error": "invalid token"})
		ctx.Abort()
		return
	}

	ctx.Next()
}
