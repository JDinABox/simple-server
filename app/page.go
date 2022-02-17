package app

import (
	"os"

	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

func Page(html []byte) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		c.Type("html", "utf-8")
		return c.SendString(string(html))
	}
}

func PageDev(path string) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		page := ReadPage(path)
		if page == nil {
			return c.Next()
		}

		return Page(page)(c)
	}
}

// ReadPage
//
// returns file | nil
func ReadPage(path string) []byte {
	f, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return nil
	} else {
		logger.Fatal(err)
	}
	return f
}
