package register

import (
	"github.com/Patt2012/authen/data"
	"github.com/Patt2012/authen/modules/register/handler"
	"github.com/Patt2012/authen/modules/users/repositories"
	"github.com/Patt2012/authen/modules/users/services"
	"github.com/gofiber/fiber/v2"
)

func Route(router *fiber.App) {

	db, err := data.Database()
	if err != nil {
		panic(err)
	}

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	regisHandler := handler.NewRegisterHandler(userService)

	api := router.Group("/api")

	api.Post("/login", regisHandler.Login)
	
}