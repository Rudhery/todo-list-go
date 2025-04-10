package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

// Estrutura da tarefa
type Todo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// Slice pra armazenar as tarefas em memória
var todos []Todo

func main() {
	// Carrega o arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Erro ao carregar .env, usando valores padrão:", err)
	}

	// Pega as variáveis de ambiente ou usa valores padrão
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Valor padrão se não tiver no .env
	}

	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost" // Valor padrão
	}

	env := os.Getenv("ENV")
	if env == "" {
		env = "development" // Valor padrão
	}

	// Log pra debug
	log.Printf("Rodando em %s:%s no ambiente %s", host, port, env)

	// Configura o Fiber
	app := fiber.New()

	// Rota pra listar todas as tarefas (READ - List)
	app.Get("/todos", func(c *fiber.Ctx) error {
		return c.JSON(todos)
	})

	// Rota pra pegar uma tarefa por ID (READ - Single)
	app.Get("/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		for _, todo := range todos {
			if todo.ID == id {
				return c.JSON(todo)
			}
		}
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Tarefa não encontrada"})
	})

	// Rota pra criar uma nova tarefa (CREATE)
	app.Post("/todos", func(c *fiber.Ctx) error {
		todo := new(Todo)
		if err := c.BodyParser(todo); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		todo.ID = uuid.New().String()
		todos = append(todos, *todo)
		return c.Status(fiber.StatusCreated).JSON(todo)
	})

	// Rota pra atualizar uma tarefa (UPDATE)
	app.Put("/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		for i, todo := range todos {
			if todo.ID == id {
				updatedTodo := new(Todo)
				if err := c.BodyParser(updatedTodo); err != nil {
					return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
				}
				todos[i].Title = updatedTodo.Title
				todos[i].Done = updatedTodo.Done
				return c.JSON(todos[i])
			}
		}
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Tarefa não encontrada"})
	})

	// Rota pra deletar uma tarefa (DELETE)
	app.Delete("/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		for i, todo := range todos {
			if todo.ID == id {
				todos = append(todos[:i], todos[i+1:]...)
				return c.SendStatus(fiber.StatusNoContent)
			}
		}
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Tarefa não encontrada"})
	})

	// Inicia o servidor com host:port
	app.Listen(host + ":" + port)
}
