package routes

import (
    "keycloakwithgo/api/handlers"

    "github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(app fiber.Router, h *handlers.UserHandler) {
    app.Post("/users", h.Create)
    app.Get("/users", h.GetAll)
    app.Get("/users/:id", h.GetByID)
    app.Put("/users/:id", h.Update)
    app.Delete("/users/:id", h.Delete)
}
