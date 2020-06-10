package cli

import (
	"fmt"

	cli "github.com/jawher/mow.cli"
)

func cmdCreate(cmd *cli.Cmd) {
	fmt.Println("this feature is still in progress")
}

const cmdCreateAlias = "create"
const cmdCreateDescription = "creates a new folder for the lab with an example for the given language"
