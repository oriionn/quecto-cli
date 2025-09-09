package main

import (
	"context"
	"fmt"
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
				Action: func(ctx context.Context, cmd *cli.Command) error {
					fmt.Println(cmd.Args().First())
                    return nil
                },
			},
			{
				Name: "unshorten",
				Aliases: []string{"u"},
				Usage: "Unshorten a link",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					fmt.Println(cmd.Args().First())
                    return nil
                },
			},
			{
				Name: "ivi",
				Aliases: []string{"i"},
				Usage: "Check if a domain is a instance of Quecto",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					fmt.Println(cmd.Args().First())
                    return nil
                },
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
        log.Fatal(err)
    }
}
