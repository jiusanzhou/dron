package cmd

import (
	"go.zoe.im/x/cli"
)

var (
	// global cmd contains all sub command
	cmd = cli.New(
		cli.Name("dron"),
		cli.Short("dron is a distributed cron service."),
		cli.Run(func(c *cli.Command, args ...string) {
			c.Help()
		}),
	)
)

// Register create a sub command
func Register(cmds ...*cli.Command) error {
	return cmd.Register(cmds...)
}

// Run execute command
func Run(opts ...cli.Option) error {
	return cmd.Run(opts...)
}
