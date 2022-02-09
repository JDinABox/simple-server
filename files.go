package simpleserver

import (
	"embed"
)

var (
	//go:embed example/lib example/pages example/package.json
	//go:embed example/postcss.config.js example/tailwind.config.js
	//go:embed example/tsconfig.json example/vite.config.js
	Files embed.FS
)
