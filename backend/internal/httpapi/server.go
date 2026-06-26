package httpapi

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Server interface {
	Start()
}

type fiberServer struct {
	app  *fiber.App
	port string
}

func NewFiberServer(port string) Server {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
		AllowMethods: "*",
	}))

	return &fiberServer{
		app:  app,
		port: port,
	}
}

func (s *fiberServer) Start() {
	RegisterRoutes(s)

	log.Fatal(s.app.Listen(":" + s.port))
}
