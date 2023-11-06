package user

import (
	"github.com/Patt2012/authen/data"
	"github.com/Patt2012/authen/middlewere"
	"github.com/Patt2012/authen/modules/users/handlers"
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
	userHandler := handlers.NewUserHandler(userService)

	api := router.Group("/api")
	v1 := api.Group("/v1")

	v1.Use(middlewere.Authorization())
	v1.Post("/register", userHandler.Register)
	v1.Patch("/users/:id", userHandler.Update)
	v1.Delete("/users/:id", userHandler.Delete)
	v1.Get("/users", userHandler.FindAll)
	v1.Get("/me/:id", userHandler.FindById)
	v1.Get("/users/:email", userHandler.FindByEmail)
	//v1.Post("/login", userHandler.Login)
	
}