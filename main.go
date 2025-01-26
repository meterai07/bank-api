package main

import (
	"flag"
	"fmt"
	"os"

	"bank/src/database"
	"bank/src/handlers"
	"bank/src/repository"

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

	app := fiber.New()
	handler := handlers.NewNasabahHandler(nasabahRepo)

	app.Post("/daftar", handler.Daftar)
	app.Post("/tabung", handler.Tabung)
	app.Post("/tarik", handler.Tarik)
	app.Get("/saldo/:no_rekening", handler.Saldo)

	port := flag.String("port", fmt.Sprintf(":%s", apiPort), "Server port")
	flag.Parse()

	log.Info().Msgf("Server starting on port %s", *port)
	app.Listen(*port)
}
