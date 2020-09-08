package main

import (
	"github.com/gogearbox/gearbox"
)

func main() {
	// Setup gearbox
	g := gearbox.New()

	// Define your handlers
	g.Get("/", func(ctx gearbox.Context) {
		ctx.SendString("Hello Eddy!")
	})

	// Start service
	g.Start(":3001")
}
