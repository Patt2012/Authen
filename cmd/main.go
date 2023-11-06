package main

import (
	"log"

	"github.com/Patt2012/authen/configs"
	"github.com/Patt2012/authen/modules/routers"

	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/fiber/v2"
)

func main() {
	configs.InitConfig("../.")

	app := fiber.New()

	app.Use(logger.New(logger.Config{
		Format:     "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
        TimeFormat: "02-Jan-2006",
        TimeZone:   "Asia/Bangkok",
	}))

	routers.Route(app)

	log.Fatal(app.Listen(":8000"))
}