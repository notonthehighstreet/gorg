package main

import (
	"os"
	"time"

	"github.com/urfave/cli"

	"github.com/notonthehighstreet/gorg/cmd/gorgcli/command"
)

const exitCode int = 86

var (
	initCmd       = command.InitCommand{}
	switchEnvCmd  = command.SwitchEnvironmentCommand{}
	switchUserCmd = command.SwitchUserCommand{}
	cfgShowCmd    = command.ConfigShowCommand{}
	cfgAddCmd     = command.ConfigAddCommand{}
	cfgRmCmd      = command.ConfigRemoveCommand{}
	srvListCmd    = command.ServiceListCommand{}
	srvShowCmd    = command.ServiceShowCommand{}
	srvOpenCmd    = command.ServiceOpenCommand{}
)

type Command interface {
	Load() error
	Validate(c *cli.Context) error
	Run() error
	String()
}

func do(cmd Command) func(*cli.Context) error {
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
				cli.StringFlag{Name: "environment-name"},
				cli.StringFlag{Name: "domain"},
			},
			Action: do(&initCmd),
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
