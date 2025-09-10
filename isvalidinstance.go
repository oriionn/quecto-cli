package main

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

func IsValidInstance(ctx context.Context, cmd *cli.Command) error {
	fmt.Println(cmd.StringArg("domain"))
	return nil
}
