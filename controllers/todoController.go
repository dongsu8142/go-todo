package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hands8142/go-todo/services"
)

func Setup(app *fiber.App) {
	todoRouter := app.Group("/todo")
	ApiRoute(todoRouter)
}

func ApiRoute(todoRouter fiber.Router) {
	todoRouter.Get("/", services.GetTodos)
	todoRouter.Post("/", services.InsertTodo)
	todoRouter.Put("/:id", services.UpdateTodo)
	todoRouter.Delete("/:id", services.DeleteTodo)
}
