package app

import (
	"errors"
	"os"

	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Config struct {
	Dev   bool `json:"-"`
	Paths struct {
		Assets string `json:"assets"`
		Pages  string `json:"pages"`
	} `json:"paths"`
	Logs struct {
		Level int8   `json:"level"`
		Dir   bool   `json:"dir"`
		Path  string `json:"path"`
	} `json:"logs"`
	Port int `json:"port"`
}

func NewConfig(path string) *Config {
	config, err := ReadConfig(path)
	logger.Fatal(err)

	config.Default()

	// Only check the config as master
	if !fiber.IsChild() {
		err = config.Validate()
		logger.Fatal(err)
	}

	return config
}

// ReadConfig initializes Config
func ReadConfig(path string) (*Config, error) {
	configration := new(Config)
	file, err := os.Open(path)
	if err != nil {
		return &Config{}, errors.New("Error reading config @ " + path)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&configration); err != nil {
		return &Config{}, err
	}
	return configration, nil
}

func DefaultConfig() *Config {
	c := new(Config)
	c.Default()
	return c
}

func (c *Config) Default() {
	if c.Port == 0 {
		c.Port = 8080
	}
	if c.Paths.Assets == "" {
		c.Paths.Assets = "./dist/assets"
	}
	if c.Paths.Pages == "" {
		c.Paths.Pages = "./dist/pages"
	}
}

func (c *Config) Validate() error {
	// TODO
	return nil
}
