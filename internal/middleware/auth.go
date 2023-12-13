package middleware

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/HtetLinMaung/todo/internal/setting"
	"github.com/HtetLinMaung/todo/internal/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Authorization header is missing!",
			})
			return
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Invalid Authorization header format",
			})
			return
		}

		tokenString := bearerToken[1]
		token, err := utils.VerifyToken(tokenString, setting.GetJwtSecret())
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": err.Error(),
			})
			return
		}
		subject, err := token.Claims.GetSubject()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": err.Error(),
			})
			return
		}
		subjects := strings.Split(subject, ",")
		user_id, err := strconv.ParseInt(subjects[0], 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "Something went wrong!",
			})
			return
		}
		c.Set("user_id", user_id)
		c.Set("role", subjects[1])
		c.Next()
	}
}
