package main

import (
	"context"
	"log"
	"os"

	"github.com/oriionn/quecto-cli/v2/config"
	"github.com/oriionn/quecto-cli/v2/utils"
	"github.com/urfave/cli/v3"
)

func main() {
	defaultDomain, err := config.Get()
	if err != nil {
		utils.PrintError(err)
		return
	}

	cmd := &cli.Command{
		Commands: []*cli.Command{
			{
				Name:    "shorten",
				Aliases: []string{"s"},
				Usage:   "Shorten a link",
				Arguments: []cli.Argument{
					&cli.StringArg{
						Name: "link",
					},
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "password",
						Aliases: []string{"p"},
					},
					&cli.StringFlag{
						Name:    "code",
						Aliases: []string{"c"},
					},
					&cli.IntFlag{
						Name:        "expires",
						DefaultText: "-1",
						Aliases:     []string{"e"},
					},
					&cli.StringFlag{
						Name: "domain",
						DefaultText: defaultDomain,
						Aliases: []string{"d"},
					},
				},
				Action: Shorten,
			},
			{
				Name:    "unshorten",
				Aliases: []string{"u"},
				Usage:   "Unshorten a link",
				Arguments: []cli.Argument{
					&cli.StringArg{
						Name: "link",
					},
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "password",
						Aliases: []string{"p"},
					},
				},
				Action: Unshorten,
			},
			{
				Name:    "ivi",
				Aliases: []string{"i", "is-valid-instance"},
				Usage:   "Check if a domain is a instance of Quecto",
				Arguments: []cli.Argument{
					&cli.StringArg{
						Name: "domain",
					},
				},
				Action: IsValidInstance,
			},
			{
				Name: "config",
				Aliases: []string{"c"},
				Commands: []*cli.Command{
					{
						Name: "get",
						Usage: "Get your default Quecto instance domain",
						Aliases: []string{"g"},
						Action: ConfigGet,
					},
					{
						Name: "set",
						Usage: "Set your default Quecto instance domain",
						Aliases: []string{"s"},
						Arguments: []cli.Argument{
							&cli.StringArg{
								Name: "domain",
							},
						},
						Action: ConfigSet,
					},
				},
				Hidden: true,
			},
			{
				Name: "config get",
				Usage: "Get your default Quecto instance domain",
				Aliases: []string{"c get"},
			},
			{
				Name: "config set",
				Usage: "Set your default Quecto instance domain",
				Arguments: []cli.Argument{
					&cli.StringArg{
						Name: "domain",
					},
				},
				Aliases: []string{"c s"},
			},
			{
				Name: "self-rename",
				Usage: "Rename the CLI executable to `quecto`",
				Action: SelfRename,
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
