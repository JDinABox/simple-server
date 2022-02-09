package app

import (
	"bytes"
	"os"

	"github.com/JDinABox/simple-server/app/template"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

func Page(i *template.Index) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		c.Type("html", "utf-8")
		buf := new(bytes.Buffer)
		i.WriteIndexTPL(buf)
		return c.SendStream(buf)
	}
}

func PageDev(header, path string) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		page := ReadPage(path)
		if page == nil {
			return c.Next()
		}

		return Page(&template.Index{Header: header, BodyHtml: string(page)})(c)
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
