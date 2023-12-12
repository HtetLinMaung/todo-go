package api

import "github.com/gin-gonic/gin"

type Router struct {
	messageRoute *MessageRoute
	authRoute    *AuthRoute
	todoRoute    *TodoRoute
	imageRoute   *ImageRoute
}

func NewRouter(messageRoute *MessageRoute, authRoute *AuthRoute, todoRoute *TodoRoute, imageRoute *ImageRoute) *Router {
	return &Router{
		messageRoute: messageRoute,
		authRoute:    authRoute,
		todoRoute:    todoRoute,
		imageRoute:   imageRoute,
	}
}

func (router *Router) SetupRouter() *gin.Engine {
	r := gin.Default()
	router.messageRoute.MessageRoutes(r)
	router.authRoute.AuthRoutes(r)
	router.todoRoute.TodoRoutes(r)
	router.imageRoute.ImageRoutes(r)
	r.Static("/images", "./images")
	return r
}
