package cli

import (
	"fmt"
	"os"

	cli "github.com/urfave/cli"
)

func Main() {
	app := cli.NewApp()
	app.Name = "welcome to S3xyFS"

	app.Commands = []cli.Command{
		{
			Name:        "init",
			HelpName:    "init",
			Action:      cliInitializeFS,
			ArgsUsage:   ` `,
			Usage:       `Initialize folder as an S3 service`,
			Description: `Exports folder as S3 `,
			Flags:       []cli.Flag{},
		},
		{
			Name:        "nc",
			HelpName:    "nc",
			Action:      cliNewCredentials,
			ArgsUsage:   ` `,
			Usage:       `Create new credentials for your folder`,
			Description: `Creates new credentials`,
			Flags:       []cli.Flag{},
		},
	}

	app.Run(os.Args)
}

func cliInitializeFS(c *cli.Context) {
	fmt.Println("Initializing FS")
}

func cliNewCredentials(c *cli.Context) {
	fmt.Println("Creating new credentials")
}
