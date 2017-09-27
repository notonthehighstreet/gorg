package main

import (
	"os"
	"time"

	"github.com/urfave/cli"

	"github.com/notonthehighstreet/gorg/cmd/gorgcli/command"
)

const exitCode int = 86

var (
	initCmd       = command.NewInitCmd()
	consoleCmd    = command.NewConsoleCmd()
	switchEnvCmd  = command.NewSwitchEnvironmentCmd()
	switchUserCmd = command.NewSwitchUserCmd()
	cfgAddCmd     = command.NewConfigAddCmd()
	cfgShowCmd    = command.NewConfigShowCmd()
	cfgRmCmd      = command.NewConfigRemoveCmd()
	kvGetCmd      = command.NewKVGetCmd()
	kvListCmd     = command.NewKVListCmd()
	srvListCmd    = command.NewServiceListCmd()
	srvShowCmd    = command.NewServiceShowCmd()
	srvOpenCmd    = command.NewServiceOpenCmd()
)

func do(cmd command.Command) func(*cli.Context) error {
	return func(c *cli.Context) (err error) {
		err = cmd.Load()
		if err != nil {
			return cli.NewExitError(err, exitCode)
		}
		err = cmd.Validate(c)
		if err != nil {
			return cli.NewExitError(err, exitCode)
		}
		err = cmd.Run()
		if err != nil {
			return cli.NewExitError(err, exitCode)
		}

		cmd.String()
		return nil
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "gorg"
	app.Version = "0.1.0"
	app.Compiled = time.Now()
	app.Usage = "CLI to interact with mesos, marathon, chronos and others"
	app.Authors = []cli.Author{
		{
			Name:  "Eggya Chiquita",
			Email: "eggyachiquita@notonthehighstreet.com",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:  "init",
			Usage: "initialise gorg configuration file",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "domain"},
			},
			Action: do(&initCmd),
		},
		{
			Name:      "console",
			Usage:     "ssh into a service box",
			ArgsUsage: "[service-name]",
			Action:    do(&consoleCmd),
		},
		{
			Name:      "use",
			Usage:     "Switch default environment",
			ArgsUsage: "[environment-name]",
			Action:    do(&switchEnvCmd),
		},
		{
			Name:      "ssh",
			Usage:     "Set provided user for ssh",
			ArgsUsage: "[username]",
			Action:    do(&switchUserCmd),
		},
		{
			Name:  "config",
			Usage: "Modify current configuration of your gorg",
			Subcommands: []cli.Command{
				{
					Name:   "show",
					Usage:  "Show current configuration of your gorg",
					Action: do(&cfgShowCmd),
				},
				{
					Name:      "add",
					Usage:     "Add an environment to current configuration",
					ArgsUsage: "[environment-name]",
					Action:    do(&cfgAddCmd),
				},
				{
					Name:      "remove",
					Usage:     "Remove an environment from your current configuration",
					ArgsUsage: "[environment-name]",
					Action:    do(&cfgRmCmd),
				},
			},
		},
		{
			Name:  "kv",
			Usage: "Interact with KVs stored in consul",
			Subcommands: []cli.Command{
				{
					Name:      "ls",
					Usage:     "List kvs for certain service",
					ArgsUsage: "[service-name]",
					Action:    do(&kvListCmd),
				},
				{
					Name:      "get",
					Usage:     "Show a particular kv from consul",
					ArgsUsage: "[key]",
					Action:    do(&kvGetCmd),
				},
			},
		},
		{
			Name:  "service",
			Usage: "Interact with running services according to consul",
			Subcommands: []cli.Command{
				{
					Name:   "ls",
					Usage:  "List running services according to consul",
					Action: do(&srvListCmd),
				},
				{
					Name:      "show",
					Usage:     "Show a particular pkg from consul",
					ArgsUsage: "[service-name]",
					Action:    do(&srvShowCmd),
				},
				{
					Name:      "open",
					Usage:     "Open selected pkg listed on Consul in a browser",
					ArgsUsage: "[service-name]",
					Action:    do(&srvOpenCmd),
				},
			},
		},
	}
	app.Run(os.Args)
}
