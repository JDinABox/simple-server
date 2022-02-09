package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"

	simpleserver "github.com/JDinABox/simple-server"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

var configPath string

func init() {
	var (
		v          bool
		initialize bool
		help       bool
	)
	const (
		defaultConfig = "./config.json"
		configUsage   = "Path to config.json"
		versionUsage  = "Returns version"
		initUsage     = "Interactively create Simple Server project"
	)

	if len(os.Args) > 1 {
		switch strings.ToLower(os.Args[1]) {
		case "version", "v":
			v = true
		case "init":
			initialize = true
		case "help", "h":
			help = true
		}
	}

	flag.StringVar(&configPath, "config", defaultConfig, configUsage)
	flag.StringVar(&configPath, "c", defaultConfig, configUsage+" (shorthand)")
	flag.BoolVar(&v, "version", v, versionUsage)
	flag.BoolVar(&initialize, "init", initialize, initUsage)

	// Custom help message
	flag.Usage = func() {
		fmtVer()
		fmt.Printf("Usage: simple-server [flags] [command] %s\n\n", simpleserver.Version)
		// Config
		fmt.Printf("  %s\n", "-c, -config string")
		fmt.Printf("  	%s\n", configUsage)
		// Version
		fmt.Printf("  %s\n", "-version, version, v")
		fmt.Printf("  	%s\n", versionUsage)
		// Init
		fmt.Printf("  %s\n", "-init, init")
		fmt.Printf("  	%s\n", initUsage)
		// Help
		fmt.Printf("  %s\n", "-help, -h, help, h")
		fmt.Printf("  	%s\n", "Output usage (this message)")
	}

	flag.Parse()

	if v {
		fmtVer()
		os.Exit(0)
	}

	if initialize {
		setup()
		os.Exit(0)
	}

	if help {
		flag.Usage()
		os.Exit(0)
	}
}

func main() {
	ss := simpleserver.New(configPath)
	ss.AddOn(compress.New(compress.ConfigDefault))

	serverClosed := make(chan struct{})
	go ss.AwaitAndClose(serverClosed)

	logger.Error(ss.Start())

	<-serverClosed
}

// fmtVer print version
func fmtVer() {
	fmt.Printf("Simple Server v%s %s/%s\n", simpleserver.Version, runtime.GOOS, runtime.GOARCH)
}
