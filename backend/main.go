package main

import (
	"log"
	"os"

	"todo-list-backend/routes" // Ajuste pro seu módulo

	"github.com/gofiber/fiber/v2"
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
		host = "localhost"
	}
	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	log.Printf("Rodando em %s:%s no ambiente %s", host, port, env)

	app := fiber.New()

	// Configura todas as rotas
	routes.SetupRoutes(app)

	app.Listen(host + ":" + port)
}
