//go:generate qtc -dir=app/template

package simpleserver

import (
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/JDinABox/simple-server/app"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
	"k8s.io/klog/v2"
)

type Server struct {
	Fiber  *fiber.App
	Config *app.Config
	addOns []func(*fiber.Ctx) error
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

const Version = "0.0.0-alpha"

func New(confPath string) *Server {
	s := new(Server)
	s.Fiber = fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	s.Config = app.NewConfig(confPath)
	logger.InitKlog(s.Config.Logs.Level, s.Config.Logs.Dir, s.Config.Logs.Path)

	return s
}

var (
	cacheDuration = 24 * time.Hour
	maxAge        = 60 * 60 * 24 * 356 // 1 year
)

func (s *Server) Start() error {
	// Load addons
	s.Fiber.Group("/*", s.addOns...)

	// Setup static file serving
	static := fiber.Static{}
	if !s.Config.Dev {
		static.CacheDuration = cacheDuration
		static.MaxAge = maxAge
	}
	s.Fiber.Static("/assets", s.Config.Paths.Assets, static)

	// Load page paths
	s.Pages()
	// TODO SSL
	return s.Fiber.Listen(":" + strconv.Itoa(s.Config.Port))
}

func (s *Server) AddOn(m func(*fiber.Ctx) error) {
	s.addOns = append(s.addOns, m)
}

func (s *Server) AwaitAndClose(serverClosed chan struct{}) {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint

	log.Println("Shutting down Fiber")
	logger.Error(s.Fiber.Shutdown())
	log.Println("Flushing klog")
	klog.Flush()

	// Done
	close(serverClosed)
}
