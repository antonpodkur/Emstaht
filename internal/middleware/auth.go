package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/antonpodkur/Emstaht/config"
	"github.com/antonpodkur/Emstaht/internal/auth"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

//TODO: Change c.JSON to not return golang error but json error {"error": "error message"}

func (mw *MiddlewareManager) JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerHeader := c.Request.Header.Get("Authorization")

		if bearerHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is empty."})
			return
		}

		headerParts := strings.Split(bearerHeader, " ")
		if len(headerParts) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong authorization header."})
			return
		}

		tokenString := headerParts[1]

		if err := mw.validateJwtToken(c, tokenString, mw.authUsecase, mw.cfg); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token validation failed."})
			return
		}

		c.Next()
	}
}

func (mw *MiddlewareManager) validateJwtToken(c *gin.Context, tokenString string, authUsecase auth.Usecase, cfg *config.Config) error {
	if tokenString == "" {
		return errors.New("token is empty")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		}
		secret := []byte(cfg.Server.JwtSecretKey)
		return secret, nil
	})
	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("token is not valid")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, ok := claims["id"].(string)
		if !ok {
			return errors.New("invalid jwt claims")
		}

		userUUID, err := uuid.Parse(userID)
		if err != nil {
			return err
		}

		user, err := authUsecase.GetByID(userUUID)
		if err != nil {
			return err
		}

		c.Set("user", user)
		// ctx := context.WithValue(c.Request.Context(), utils.UserCtxKey{}, user)
		// c
	}
	return nil
}
