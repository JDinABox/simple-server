package main

import (
	"log"

	simpleserver "github.com/JDinABox/simple-server"
	"github.com/JDinABox/simple-server/app"
)

func main() {
	ss := simpleserver.New(app.Config{})
	ss.AddOn(app.Page())
	log.Fatal(ss.Start())
}
