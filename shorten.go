package main

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func Shorten(ctx context.Context, cmd *cli.Command) error {
	fmt.Println(cmd.Args().First())
    return nil
}
