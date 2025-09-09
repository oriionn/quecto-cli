package main

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func Unshorten(ctx context.Context, cmd *cli.Command) error {
	fmt.Println(cmd.Args().First())
    return nil
}
