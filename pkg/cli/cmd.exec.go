package cli

import (
	"fmt"

	cli "github.com/jawher/mow.cli"
)

func cmdExec(cmd *cli.Cmd) {
	fmt.Println("this feature is still in progress")
}

const cmdExecAlias = "exec"
const cmdExecDescription = "executes the given command in all lab subfolder"
