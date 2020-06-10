package main

import (
	"log"
	"os"

	"github.com/naxmefy/lab/pkg/cli"
)

func main() {
	err := cli.App().Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
