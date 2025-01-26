package main

import (
	"flag"
	"fmt"
	"os"

	"bank/src/database"
	"bank/src/handlers"
	"bank/src/middleware"
	"bank/src/repository"
	"bank/src/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// for local development
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal().Err(err).Msg("Failed to load .env file")
	// }

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	apiPort := os.Getenv("PORT")
	if apiPort == "" {
		apiPort = "8080"
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	nasabahRepo := repository.NewNasabahRepository(db)
	if err := nasabahRepo.AutoMigrate(); err != nil {
		log.Fatal().Err(err).Msg("Failed to migrate database")
	}

	nasabahHandler := handlers.NewNasabahHandler(nasabahRepo)
	authHandler := handlers.NewAuthHandler(nasabahRepo)

	app := fiber.New()
	middleware.InitJWTSecret(os.Getenv("JWT_SECRET"))
	routes.SetupRoutes(app, nasabahHandler, authHandler)

	port := flag.String("port", fmt.Sprintf(":%s", apiPort), "Server port")
	flag.Parse()

	log.Info().Msgf("Server starting on port %s", *port)
	app.Listen(*port)
}
