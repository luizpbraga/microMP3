package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/luizpbraga/microMP3/src/database"
	"github.com/luizpbraga/microMP3/src/services/auth/routes"
	"log"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// DATABASE
	if db, err := database.InitDataBase(); err != nil {
		log.Fatal(err)
		// BUG???
		defer db.Close()
	}

	if err := database.Db.Ping(); err != nil {
		log.Fatal(err)
	}

	// SERVER
	app := fiber.New()
	routes.LoadRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
