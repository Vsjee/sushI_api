package main

import (
	"log"
	"os"
	"sushi-api/configs"
	routes "sushi-api/routes/sushi"

	"github.com/gofiber/fiber/v2"
)

func main() {
	//fiber instance
	app := fiber.New()

	//run database
	configs.ConnectDB()

	//routes
	//base api
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("API for study purposes, gitrepo: https://github.com/Vsjee/sushi_api")
	})
	//questions api
	routes.QuestionsRoute(app)

	//get port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	//start server
	log.Fatal(app.Listen("0.0.0.0:" + port))
}
