package routes

import (
	"todo-list-backend/routes/todo" // Import direto, sem alias, ajustado pro seu módulo

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes configura todas as rotas da API
func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")  // Prefixo /api pra todas as rotas
	todo.SetupTodoRoutes(api) // Chama as rotas do todo com o grupo
}
