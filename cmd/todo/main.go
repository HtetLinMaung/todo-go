package main

import (
	"fmt"
	"log"
	"os"

	"github.com/HtetLinMaung/todo/internal/api"
	"github.com/HtetLinMaung/todo/internal/db"
	"github.com/HtetLinMaung/todo/internal/service"
	"github.com/HtetLinMaung/todo/internal/setting"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	imageDir := "./images"
	if _, err := os.Stat(imageDir); os.IsNotExist(err) {
		err = os.MkdirAll(imageDir, os.ModePerm)
		if err != nil {
			log.Fatalf("Error creating directory: %v", err)
		}
	}

	connString := setting.GetConnectionString()
	port := setting.GetPort()

	fmt.Println("Connection string: ", connString)
	database, err := db.NewConnection(connString)
	if err != nil {
		log.Fatal("Could not connect to database:", err)
	}
	defer database.Close()

	messageService := service.NewMessageService(database)
	messageRoute := api.NewMessageRoute(messageService)

	userService := service.NewUserService(database)
	authRoute := api.NewAuthRoute(userService)

	todoService := service.NewTodoService(database)
	todoRoute := api.NewTodoRoute(todoService)

	imageRoute := api.NewImageRoute()

	router := api.NewRouter(messageRoute, authRoute, todoRoute, imageRoute)
	gin.SetMode(gin.ReleaseMode)
	r := router.SetupRouter()
	r.Run(fmt.Sprintf("0.0.0.0:%s", port))
}
