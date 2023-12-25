package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luizpbraga/microMP3/services/auth/src/database/connection"
	"github.com/luizpbraga/microMP3/services/auth/src/router"
	"log"
)

func main() {
	// DATABASE
	if db, err := connection.InitDataBase(); err != nil {
		log.Fatal(err)
		// BUG???
		defer db.Close()
	}

	if err := connection.Db.Ping(); err != nil {
		log.Fatal(err)
	}

	// SERVER
	app := fiber.New()
	router.LoadRoutes(app)
	log.Fatal(app.Listen("0.0.0.0:3000"))
}
