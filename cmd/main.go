package main

import (
    "keycloakwithgo/api/handlers"
    "keycloakwithgo/api/routes"
    "keycloakwithgo/domain/user"
    "keycloakwithgo/infra/database"

    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    // Repo + Service
    repo := database.NewInMemoryUserRepo()
    service := user.NewService(repo)

    // Handler
    userHandler := handlers.NewUserHandler(service)

    // Routes
    routes.RegisterUserRoutes(app, userHandler)

    app.Listen(":3000")
}
