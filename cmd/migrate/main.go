package main

import (
	"github.com/iota-uz/iota-sdk/pkg/commands"
)

func main() {
	err := commands.Migrate()
	if err != nil {
		panic(err)
	}
}
