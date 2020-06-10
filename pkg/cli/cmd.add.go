package cli

import (
	"fmt"

	cli "github.com/jawher/mow.cli"
)

func cmdAdd(cmd *cli.Cmd) {
	fmt.Println("this feature is still in progress")
}

const cmdAddAlias = "add"
const cmdAddDescription = "add the next lab subfolder (incremeting - e.g. if 00015_... exists it will create 00016...)	"
