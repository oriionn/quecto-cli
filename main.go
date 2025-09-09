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
				Name: "shorten",
				Aliases: []string{"s"},
				Usage: "Shorten a link",
				Action: Shorten,
			},
			{
				Name: "unshorten",
				Aliases: []string{"u"},
				Usage: "Unshorten a link",
				Action: Unshorten,
			},
			{
				Name: "ivi",
				Aliases: []string{"i"},
				Usage: "Check if a domain is a instance of Quecto",
				Action: IsValidInstance,
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
        log.Fatal(err)
    }
}
