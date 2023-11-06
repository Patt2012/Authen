package handler

import "github.com/gofiber/fiber/v2"

type RisterHandler interface{
	Login(*fiber.Ctx) error
}
