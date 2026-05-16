package main

import (
	"log"

	"github.com/kmullin/advent-of-code/internal/cli"
)

func ReadInput(b []byte) int {
	return 0
}

func main() {
	ctx, err := cli.Setup(nil)
	if err != nil {
		log.Fatal(err)
	}

	thing := ReadInput(ctx.Bytes())
	log.Print(thing)
}
