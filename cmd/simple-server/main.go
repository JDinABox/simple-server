package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"

	simpleserver "github.com/JDinABox/simple-server"
	"github.com/allocamelus/allocamelus/pkg/logger"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/etag"
)

var (
	configPath string
	dev        bool
	build      bool
)

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
		devUsage      = "Run package.json script dev & Start Server in Development mode"
		buildUsage    = "Run package.json script build & Start Server"
	)

	if len(os.Args) > 1 {
		switch strings.ToLower(os.Args[1]) {
		case "version", "v":
			v = true
		case "init":
			initialize = true
		case "help", "h":
			help = true
		case "dev":
			dev = true
		case "build":
			build = true
		}
	}

	flag.StringVar(&configPath, "config", defaultConfig, configUsage)
	flag.StringVar(&configPath, "c", defaultConfig, configUsage+" (shorthand)")
	flag.BoolVar(&v, "version", v, versionUsage)
	flag.BoolVar(&initialize, "init", initialize, initUsage)
	flag.BoolVar(&dev, "dev", dev, devUsage)
	flag.BoolVar(&dev, "D", dev, devUsage)
	flag.BoolVar(&build, "build", build, buildUsage)
	flag.BoolVar(&build, "b", build, buildUsage)

	// Custom help message
	flag.Usage = func() {
		fmtVer()
		fmt.Printf("Usage: simple-server [flags] [command] %s\n\n", simpleserver.Version)
		// Config
		usageLines("-c, -config string", configUsage)
		// Dev
		usageLines("-D, -dev, dev", devUsage)
		// Build
		usageLines("-B, -build, build", buildUsage)
		// Version
		usageLines("-version, version, v", versionUsage)
		// Init
		usageLines("-init, init", initUsage)
		// Help
		usageLines("-h, -help, help, h", "Output usage (this message)")
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
	ss.Config.Dev = dev

	serverClosed := make(chan struct{})
	go ss.AwaitAndClose(serverClosed)

	if ss.Config.Dev {
		cmd := runDev()
		defer func() {
			log.Printf("Killing %s\n", cmd.Path)
			cmd.Process.Signal(os.Kill)
		}()
	} else {
		ss.AddOn(compress.New(compress.ConfigDefault))
		ss.AddOn(etag.New(etag.ConfigDefault))
	}

	if build {
		runBuild()
	}

	logger.Error(ss.Start())
	<-serverClosed
}

// fmtVer print version
func fmtVer() {
	fmt.Printf("Simple Server v%s %s/%s\n", simpleserver.Version, runtime.GOOS, runtime.GOARCH)
}

func usageLines(flags, usageText string) {
	fmt.Printf("  %s\n", flags)
	fmt.Printf("  	%s\n", usageText)
}
