package controllers

import (
	"todo-list-backend/models" // Ajuste o caminho pro seu módulo

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Todos é a "base de dados" em memória
var Todos []models.Todo

// GetTodos lista todas as tarefas
func GetTodos(c *fiber.Ctx) error {
	return c.JSON(Todos)
}

// GetTodo pega uma tarefa por ID
func GetTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	for _, todo := range Todos {
		if todo.ID == id {
			return c.JSON(todo)
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Tarefa não encontrada"})
}

// CreateTodo adiciona uma nova tarefa
func CreateTodo(c *fiber.Ctx) error {
	todo := new(models.Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	todo.ID = uuid.New().String()
	Todos = append(Todos, *todo)
	return c.Status(fiber.StatusCreated).JSON(todo)
}

// UpdateTodo atualiza uma tarefa
func UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	for i, todo := range Todos {
		if todo.ID == id {
			updatedTodo := new(models.Todo)
			if err := c.BodyParser(updatedTodo); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
			}
			Todos[i].Title = updatedTodo.Title
			Todos[i].Done = updatedTodo.Done
			return c.JSON(Todos[i])
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Tarefa não encontrada"})
}

// DeleteTodo remove uma tarefa
func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	for i, todo := range Todos {
		if todo.ID == id {
			Todos = append(Todos[:i], Todos[i+1:]...)
			return c.SendStatus(fiber.StatusNoContent)
		}
	}
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Tarefa não encontrada"})
}
