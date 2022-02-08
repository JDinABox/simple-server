package app

import (
	"bytes"

	"github.com/JDinABox/simple-server/app/template"
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
