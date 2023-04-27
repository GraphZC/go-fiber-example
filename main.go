package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"go-fiber-example/initializers"
	"go-fiber-example/router"
)

func init() {
	config, err := initializers.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load environment variables! \n", err.Error())
	}
	log.Println(config)
}

func main() {
	app := fiber.New(fiber.Config{
		AppName: "FiberApp",
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE",
		AllowCredentials: true,
	}))

	router.ConfigRoute(app)

	log.Fatal(app.Listen(":8080"))
}