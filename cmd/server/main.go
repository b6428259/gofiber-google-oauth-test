package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2/middleware/logger"

	"gofiber-hex-google-oauth/internal/adapters/googleoauth"
	httpadapter "gofiber-hex-google-oauth/internal/adapters/http"
	"gofiber-hex-google-oauth/internal/app"
)

func main() {
	// Load .env file
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	oauthClient, err := googleoauth.New()
	if err != nil {
		log.Fatal(err)
	}

	svc := app.NewAuthService(oauthClient)

	srv := httpadapter.New(svc)
	app := srv.Router()
	app.Use(logger.New())

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Printf("listening on :%s\n", port)
	log.Fatal(app.Listen(":" + port))
}
