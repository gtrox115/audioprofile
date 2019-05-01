package main

import (
	"log"

	"github.com/gtrox115/audioprofile/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
