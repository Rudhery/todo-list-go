package todo

import (
	"todo-list-backend/controllers" // Ajuste pro seu módulo

	"github.com/gofiber/fiber/v2"
)

// SetupTodoRoutes configura as rotas do CRUD de todos
func SetupTodoRoutes(router fiber.Router) {
	router.Get("/todos", controllers.GetTodos)
	router.Get("/todos/:id", controllers.GetTodo)
	router.Post("/todos", controllers.CreateTodo)
	router.Put("/todos/:id", controllers.UpdateTodo)
	router.Delete("/todos/:id", controllers.DeleteTodo)
}
