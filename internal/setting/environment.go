package setting

import (
	"log"
	"os"
)

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8080"
	}
	return port
}

func GetConnectionString() string {
	connString := os.Getenv("DB_CONNECTION")
	if connString == "" {
		log.Fatal("DB_CONNECTION must be set")
	}
	return connString
}

func GetJwtSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET must be set")
	}
	return secret
}
