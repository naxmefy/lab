package cli

import (
	"fmt"

	cli "github.com/jawher/mow.cli"
)

func cmdRemove(cmd *cli.Cmd) {
	fmt.Println("this feature is still in progress")
}

const cmdRemoveAlias = "remove"
const cmdRemoveDescription = "removes the given folder index (will ask fo decrementing all following and confirmation)"
