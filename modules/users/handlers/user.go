package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type UserHandler interface{
	Register(*fiber.Ctx) error
	Update(*fiber.Ctx) error
	Delete(*fiber.Ctx) error
	FindById(*fiber.Ctx) error
	FindByEmail(*fiber.Ctx) error
	FindAll(*fiber.Ctx) error
}