package main

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
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
				Aliases: []string{"i", "isvalidinstance"},
				Usage:   "Check if a domain is a instance of Quecto",
				Arguments: []cli.Argument{
					&cli.StringArg{
						Name: "domain",
					},
				},
				Action: IsValidInstance,
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
