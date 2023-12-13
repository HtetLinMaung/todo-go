package api

import (
	"fmt"
	"net/http"

	"github.com/HtetLinMaung/todo/internal/model"
	"github.com/HtetLinMaung/todo/internal/service"
	"github.com/HtetLinMaung/todo/internal/setting"
	"github.com/HtetLinMaung/todo/internal/utils"
	"github.com/gin-gonic/gin"
)

type AuthRoute struct {
	userServicee *service.UserService
}

func NewAuthRoute(userServicee *service.UserService) *AuthRoute {
	return &AuthRoute{userServicee: userServicee}
}

func (ar *AuthRoute) AuthRoutes(r *gin.Engine) {
	authGroup := r.Group("/api/auth")
	authGroup.POST("/register", ar.Register)
	authGroup.POST("/login", ar.Login)
}

func (ar *AuthRoute) Register(c *gin.Context) {
	var userRequest model.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	exists, err := ar.userServicee.IsUserExists(userRequest.Username)

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Error checking user existence!",
		})
		return
	}

	if exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "User already exists!",
		})
		return
	}

	userRequest.AccountStatus = "active"
	userRequest.Role = "user"

	_, err = ar.userServicee.AddUser(&userRequest)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Error adding user to database!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "User registered successfully.",
	})
}

func (ar *AuthRoute) Login(c *gin.Context) {
	var loginRequest model.LoginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	user, err := ar.userServicee.GetUserByUsername(loginRequest.Username)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Error getting user from username!",
		})
		return
	}
	if err == nil && user == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid username!",
		})
		return
	}

	if user.AccountStatus != "active" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "Your account has not been activated yet. Please wait for an admin to approve your account or contact support for further assistance!",
		})
		return
	}

	if !utils.VerifyPassword(user.Password, loginRequest.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "Invalid password!",
		})
		return
	}

	tokenString, err := utils.SignToken(fmt.Sprintf("%d,%s", user.UserID, user.Role), setting.GetJwtSecret())
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Error signing jwt token!",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Login successful.",
		"data": gin.H{
			"token":         tokenString,
			"name":          user.Name,
			"profile_image": user.ProfileImage,
			"role":          user.Role,
		},
	})
}
