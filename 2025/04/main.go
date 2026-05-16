package main

import (
	"io"
	"log"

	"github.com/kmullin/advent-of-code/internal/cli"
)

func ReadInput(r io.Reader) int {
	return 0
}

func main() {
	ctx, err := cli.Setup(nil)
	if err != nil {
		log.Fatal(err)
	}

	thing := ReadInput(ctx.Reader())
	log.Print(thing)
}
