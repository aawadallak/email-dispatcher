package http

import (
	"fmt"
	"latest/app/http/routes"
	"latest/config"
	zapper "latest/pkg/logger"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	Port   string
	Server *fiber.App
}

func NewServer() Server {
	return Server{
		Port:   config.GetConfig().ServerPort,
		Server: fiber.New(fiber.Config{}),
	}
}

func (s *Server) Run() {

	s.Server.Use(logger.New(logger.ConfigDefault))

	zapper.Instance().Info("Starting API Service")

	router := routes.SetRoute(s.Server)
	log.Fatalln(router.Listen(fmt.Sprintf(":%s", s.Port)))
}
