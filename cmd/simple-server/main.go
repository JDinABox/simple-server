package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	simpleserver "github.com/JDinABox/simple-server"
	"github.com/JDinABox/simple-server/app"
	jsoniter "github.com/json-iterator/go"
)

var configPath string
var json = jsoniter.ConfigCompatibleWithStandardLibrary

func init() {
	const (
		defaultConfig = "./config.json"
		configUsage   = "Path to config.json"
	)

	flag.StringVar(&configPath, "config", defaultConfig, configUsage)
	flag.StringVar(&configPath, "c", defaultConfig, configUsage+" (shorthand)")
	v := flag.Bool("version", false, "Returns version")
	genConf := flag.Bool("genConf", false, "Returns default Config")
	flag.Parse()
	if *v {
		fmt.Println(simpleserver.Version)
		os.Exit(0)
	}
	if *genConf {
		j, _ := json.MarshalIndent(app.DefaultConfig(), "", "  ")
		fmt.Println(string(j))
		os.Exit(0)
	}
}

func main() {
	ss := simpleserver.New(configPath)
	log.Fatal(ss.Start())
}
