package api

import (
	"github.com/HtetLinMaung/todo/internal/service"
	"github.com/gin-gonic/gin"
)

type MessageRoute struct {
	messageService *service.MessageService
}

func NewMessageRoute(messageService *service.MessageService) *MessageRoute {
	return &MessageRoute{messageService: messageService}
}

func (mr *MessageRoute) MessageRoutes(r *gin.Engine) {
	messageGroup := r.Group("/messages")
	messageGroup.GET("/", mr.GetMessage)
}

func (mr *MessageRoute) GetMessage(c *gin.Context) {
	msg := mr.messageService.GetMessage()
	c.JSON(200, gin.H{"message": msg})
}
