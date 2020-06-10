package cli

import (
	cli "github.com/jawher/mow.cli"
	"github.com/naxmefy/lab/pkg"
)

// App is the primary cli application of this project
func App() *cli.Cli {
	app := cli.App("github", "github cli tool")
	app.Version("v version", pkg.VERSION)

	app.Command(cmdCreateAlias, cmdCreateDescription, cmdCreate)
	app.Command(cmdAddAlias, cmdAddDescription, cmdAdd)
	app.Command(cmdRemoveAlias, cmdRemoveDescription, cmdRemove)
	app.Command(cmdExecAlias, cmdExecDescription, cmdExec)

	app.Action = func() {
		app.PrintHelp()
	}

	return app
}
