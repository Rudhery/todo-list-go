package main

import (
	"log"
	"os"

	"todo-list-backend/routes" // Ajuste pro seu módulo

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors" // Add this import
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Erro ao carregar .env, usando padrões:", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	host := os.Getenv("HOST")
	if host == "" {
		host = "0.0.0.0"
	}
	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	log.Printf("Rodando em %s:%s no ambiente %s", host, port, env)

	app := fiber.New()

	// Configure CORS with dynamic origin checking
	app.Use(cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool {
			return true // Allow all origins during development
		},
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
		ExposeHeaders:    "Content-Length, Content-Type",
	}))

	// Configura todas as rotas
	routes.SetupRoutes(app)

	app.Listen(host + ":" + port)
}
