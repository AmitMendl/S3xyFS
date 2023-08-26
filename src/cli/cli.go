package cli

import (
	"fmt"

	cli "github.com/urfave/cli"
)

func Main() {
	app := cli.NewApp()
	app.Name = "welcome to S3xyFS"

	app.Commands = []cli.Command{
		{
			Name:        "nc",
			HelpName:    "nc",
			Action:      newCredentials,
			ArgsUsage:   ` `,
			Usage:       `Create new credentials for your folder`,
			Description: `Creates new credentials`,
			Flags:       []cli.Flag{},
		},
	}
}

func newCredentials(c *cli.Context) {
	fmt.Println("Hello hi")
}
