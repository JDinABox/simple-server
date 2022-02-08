package simpleserver

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/JDinABox/simple-server/app"
	"github.com/JDinABox/simple-server/app/template"
	"github.com/allocamelus/allocamelus/pkg/logger"
)

func (s *Server) Pages() {
	logger.Fatal(filepath.Walk(s.Config.Paths.Pages, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		publicPath := strings.TrimLeft(path, filepath.Clean(s.Config.Paths.Pages))
		// Trim .html
		publicPath = strings.TrimRight(publicPath, ".html")
		// Trim index
		if strings.Contains(info.Name(), "index") {
			publicPath = strings.TrimRight(publicPath, "index")
		}

		// Read html file
		f, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		// Add route to fiber
		s.Fiber.Get(
			publicPath,
			app.Page(&template.Index{Header: s.Config.GenHeader, BodyHtml: string(f)}),
		)
		return nil
	}))
}
