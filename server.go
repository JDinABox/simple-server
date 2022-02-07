//go:generate qtc -dir=app/template

package simpleserver

import (
	"strconv"

	"github.com/JDinABox/simple-server/app"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

type Server struct {
	Fiber  *fiber.App
	Config *app.Config
	addOns []func(*fiber.Ctx) error
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func New(c app.Config) *Server {
	s := new(Server)
	s.Fiber = fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	s.Config = &c
	s.Config.Default()

	return s
}

func (s *Server) Start() error {
	s.Fiber.Group("/", s.addOns...)
	// TODO SSL
	return s.Fiber.Listen(":" + strconv.Itoa(s.Config.Port))
}

func (s *Server) AddOn(m func(*fiber.Ctx) error) {
	s.addOns = append(s.addOns, m)
}