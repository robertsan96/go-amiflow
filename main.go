package main

import (
	"github.com/robertsan96/go-amiflow/adapter/router"
)

func startGinEngine() {
	engine := router.NewGinEngine()
	engine.Run(":8080")
}

func main() {
	startGinEngine()
}
