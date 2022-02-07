package app

import (
	"bytes"

	"github.com/JDinABox/simple-server/app/template"
	"github.com/gofiber/fiber/v2"
)

func Page() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		c.Type("html", "utf-8")
		i := template.Index{BodyHtml: "<p>Hello, World!</p>"}
		buf := new(bytes.Buffer)
		i.WriteIndexTPL(buf)
		return c.SendStream(buf)
	}
}
