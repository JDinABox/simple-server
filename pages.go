package simpleserver

import (
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/JDinABox/simple-server/app"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/svg"
)

func (s *Server) Pages() {
	m := minify.New()
	m.Add("text/html", &html.Minifier{
		KeepComments: true,
	})
	m.AddFunc("image/svg+xml", svg.Minify)

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

		if s.Config.Dev {
			s.Fiber.Get(
				publicPath,
				app.PageDev(path),
			)
		} else {
			// Read html file
			f := app.ReadPage(path)
			// minify
			f, err = m.Bytes("text/html", f)
			if err != nil {
				logger.Fatal(err)
			}
			// Add route to fiber
			s.Fiber.Get(
				publicPath,
				app.Page(f),
			)
		}

		return nil
	}))
}
