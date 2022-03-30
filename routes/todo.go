package routes

import (
	"api-todo/controller"

	"github.com/gofiber/fiber/v2"
)

func TodoRoute(route fiber.Router) {
	route.Get("/", controller.GetTodos)
	route.Post("/", controller.CreateTodo)
	route.Put("/:id", controller.UpdateTodo)
	route.Delete("/:id", controller.DeleteTodo)
	route.Get("/:id", controller.GetTodo)
}
