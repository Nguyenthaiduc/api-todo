package main

import (
	"api-todo/routes"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func setupRoutes(app *fiber.App){
	app.Get("/",func(c *fiber.Ctx) error{
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"success" : true,
			"message" : "Bạn đang ở cuối trang",
		})
	})
	//set up route in Group

	api := app.Group("/api")

	//send response khi user at / api
	api.Get("/",func(c *fiber.Ctx) error{
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"success" : true,
			"message" : "Bạn đang ở cuối trang",
		})
	})
	//routes group
	routes.TodoRoute(api.Group("/todos"))

}

func main() {
	//khai bao app = fiber tuong tu nhu const app = express()
	app := fiber.New()
	app.Use(logger.New())
	app.Get("/home", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(fiber.Map{"message": "Hello World"})
	})

	setupRoutes(app)

	log.Fatal(app.Listen(":5000"))
}
