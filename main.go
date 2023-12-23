package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luizpbraga/microMP3/src/database"
	"github.com/luizpbraga/microMP3/src/server/routes"
	"log"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	// DATABASE
	db, err := database.InitDataBase()

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	// SERVER
	app := fiber.New()
	routes.LoadRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
