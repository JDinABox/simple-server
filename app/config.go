package app

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
	"k8s.io/klog/v2"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Config struct {
	Paths struct {
		Assets string `json:"assets"`
		Pages  string `json:"pages"`
	} `json:"paths"`
	Headers     []string `json:"headers"`
	headerPaths *HeaderPaths
	GenHeader   string
	Logs        struct {
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

	config.fillHeaderPaths()
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
		c.Paths.Assets = "./assets"
	}
	if c.Paths.Pages == "" {
		c.Paths.Pages = "./pages"
	}
	if c.Headers == nil {
		c.Headers = []string{}
	}
}

func (c *Config) Validate() error {
	// TODO
	return nil
}

type HeaderPaths struct {
	JS  []string
	CSS []string
}

func (c *Config) fillHeaderPaths() {
	c.headerPaths = &HeaderPaths{}

	for _, v := range c.Headers {
		fileP := filepath.Join(c.Paths.Assets, v)
		switch filepath.Ext(v) {
		case ".css":
			c.headerPaths.CSS = append(c.headerPaths.CSS, fileP)
			break
		case ".js":
			c.headerPaths.JS = append(c.headerPaths.JS, fileP)
			break
		default:
			klog.Warningf("Warning: %s not css or js\n", v)
		}
	}
	c.GenHeader = c.headerPaths.genHeader()
}

func (h *HeaderPaths) genHeader() string {
	var header strings.Builder

	for _, v := range h.JS {
		header.WriteString(`<script src="/`)
		header.WriteString(v)
		header.WriteString(`" defer></script>`)
	}

	for _, v := range h.CSS {
		header.WriteString(`<link rel="stylesheet" href="/`)
		header.WriteString(v)
		header.WriteString(`">`)
	}

	return header.String()
}
