package handler

import (
	"time"

	"github.com/Patt2012/authen/configs"
	"github.com/Patt2012/authen/modules/users/services"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type registerHandler struct {
	service services.UserSevice
}

func NewRegisterHandler(service services.UserSevice) RisterHandler {
	return registerHandler{service}
}

func (h registerHandler) Login(c *fiber.Ctx) error {

	config, _ := configs.InitConfig(".")

	//email := c.FormValue("email")
	var data map[string]string
	
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	if data["email"] == "" || data["password"] == "" {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "should email or password not be null")
	}

	user, err := h.service.FindByEmail(c.Context(), data["email"])
	if user.Email == "" {		
		return fiber.NewError(fiber.StatusNotFound, "incorrect email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"]))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "incorrect password")
	}

	cliams := jwt.RegisteredClaims{
		Issuer: user.ID,
		Subject: user.Email,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)
	token, err := jwtToken.SignedString([]byte(config.JwtSecret))
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}

