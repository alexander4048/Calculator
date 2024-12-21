package main

import (
	"github.com/alexander4048/calculator/internal/application"
)

func main() {
	app := application.New()
	app.RunServer()
}
