package routers

import (
	register "github.com/Patt2012/authen/modules/register/routes"
	user "github.com/Patt2012/authen/modules/users/routes"
	"github.com/gofiber/fiber/v2"
)

func Route(router *fiber.App) {
	user.Route(router)
	register.Route(router)	
}