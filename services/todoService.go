package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hands8142/go-todo/database"
	"github.com/hands8142/go-todo/models"
)

func GetTodos(c *fiber.Ctx) error {
	var todos []models.Todo
	database.DB.Find(&todos)
	return c.JSON(todos)
}

func InsertTodo(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	todo := models.Todo{
		Completed: false,
		TodoName:  data["todoName"],
	}

	if result := database.DB.Create(&todo); result.Error != nil {
		return result.Error
	}

	return c.JSON(todo)
}

func UpdateTodo(c *fiber.Ctx) error {
	var todo models.Todo
	if result := database.DB.First(&todo, c.Params("id")); result.Error != nil {
		return result.Error
	}
	todo.Completed = !todo.Completed
	database.DB.Save(&todo)
	return c.JSON(todo)
}

func DeleteTodo(c *fiber.Ctx) error {
	var todo models.Todo
	if result := database.DB.Delete(&todo, c.Params("id")); result.Error != nil {
		return result.Error
	}
	return c.JSON(fiber.Map{"message": "success"})
}
