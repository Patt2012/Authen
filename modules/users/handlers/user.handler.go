package handlers

import (
	"github.com/Patt2012/authen/modules/users/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userHandler struct {
	service services.UserSevice
}

func NewUserHandler(service services.UserSevice) UserHandler {
	return userHandler{service}
}

func (h userHandler) Update(c *fiber.Ctx) error {

	user := services.UserUpdate{}

	id := c.Params("id")
	log.Info(id)
	log.Info(uuid.MustParse(id).String())
	err := c.BodyParser(&user)
	if err != nil {
		return err
	}
	
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return  fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	user = services.UserUpdate{
		Password: string(password),
		FullName: user.FullName,
		Active: user.Active,
	}

	err = h.service.Update(c.Context(), user, id)
	log.Info(err)
	log.Info("update : ", user)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound,"id not found")
	}

	return nil
}

func (h userHandler) Register(c *fiber.Ctx) error {

	var err error

	var data map[string]string

	err = c.BodyParser(&data)
	if err != nil {
		return err
	}	

	if data["email"] == "" || data["password"] == "" || data["full_name"] == "" {
		return fiber.NewError(fiber.StatusBadRequest, "should email or password or fullname not be null")
	}

	if data["password"] != data["password_confirm"] {
		return fiber.NewError(fiber.StatusBadRequest, "password not macth")
	}

	user, err := h.service.FindByEmail(c.Context(), data["email"])
	if user.Email == data["email"] {
		return fiber.NewError(fiber.StatusFound,"user exist")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(data["password"]), 10)
	if err != nil {
		return  fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	register := services.RegisterRequest{
		Email: data["email"],
		Password: string(password),
		FullName: data["full_name"],
	}

	err = h.service.Insert(c.Context(), register)
	if err != nil {
		return  fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (h userHandler) FindByEmail(c *fiber.Ctx) error {

	email := c.Params("email")

	user, err := h.service.FindByEmail(c.Context(), email)
	log.Info(user)
	log.Info(err)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound,"email not found")
	}
	if user.Email == "" {
		return fiber.NewError(fiber.StatusNotFound,"email not found")
	}

	return c.Status(fiber.StatusFound).JSON(user)
}

func (h userHandler) FindById(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := h.service.FindById(c.Context(), id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound,"id not found")
	}

	return c.Status(fiber.StatusFound).JSON(user)
}

func (h userHandler) FindAll(c *fiber.Ctx) error {
	users, err := h.service.FindAll(c.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	return c.Status(fiber.StatusFound).JSON(&users)
}

func (h userHandler) Delete(c *fiber.Ctx) error {

	id := c.Params("id")
	err := h.service.Delete(c.Context(), id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound,"id not found")
	}

	return c.SendStatus(fiber.StatusOK)
}

