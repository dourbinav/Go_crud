package handlers

import (
    "keycloakwithgo/domain/user"
    "strconv"
	"fmt"
    "github.com/gofiber/fiber/v2"
)

type UserHandler struct {
    service *user.Service
}

func NewUserHandler(s *user.Service) *UserHandler {
    return &UserHandler{service: s}
}

func (h *UserHandler) Create(c *fiber.Ctx) error {
    var u user.User
	fmt.Println(c)
    if err := c.BodyParser(&u); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "invalid input"})
    }
    created, _ := h.service.CreateUser(&u)
    return c.Status(201).JSON(created)
}

func (h *UserHandler) GetAll(c *fiber.Ctx) error {
    users, _ := h.service.GetUsers()
    return c.JSON(users)
}

func (h *UserHandler) GetByID(c *fiber.Ctx) error {
    id, _ := strconv.Atoi(c.Params("id"))
    u, err := h.service.GetUser(id)
    if err != nil {
        return c.Status(404).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(u)
}

func (h *UserHandler) Update(c *fiber.Ctx) error {
    id, _ := strconv.Atoi(c.Params("id"))
    var u user.User
    if err := c.BodyParser(&u); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "invalid input"})
    }
    updated, err := h.service.UpdateUser(id, &u)
    if err != nil {
        return c.Status(404).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(updated)
}

func (h *UserHandler) Delete(c *fiber.Ctx) error {
    id, _ := strconv.Atoi(c.Params("id"))
    err := h.service.DeleteUser(id)
    if err != nil {
        return c.Status(404).JSON(fiber.Map{"error": err.Error()})
    }
    return c.SendStatus(204)
}
