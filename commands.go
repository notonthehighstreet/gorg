package main

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
	"github.com/notonthehighstreet/gorg/service"
	"github.com/urfave/cli"
)

func initCommand() cli.Command {
	return cli.Command{
		Name:  "init",
		Usage: "Initialise gorg configuration file",
		Action: func(c *cli.Context) error {
			env := service.NewEnvironment(DefaultEnvironment, EnvironmentDomain)
			err := service.NewConfig(env)
			if err != nil {
				return err
			}

			return nil
		},
	}
}

func sshuserCommand() cli.Command {
	return cli.Command{
		Name:  "sshuser",
		Usage: "Set provided user for ssh",
		Action: func(c *cli.Context) error {
			args := c.Args()
			if len(args) == 1 {
				config := loadConfig()
				return config.ChangeUser(args[0])
			}
			return errors.New("you need to supply username")
		},
	}
}

func configCommand() cli.Command {
	return cli.Command{
		Name:  "config",
		Usage: "Modify current configuration of your gorg",
		Subcommands: []cli.Command{
			showConfigCommand(),
			addConfigCommand(),
			removeConfigCommand(),
		},
	}
}

func showConfigCommand() cli.Command {
	return cli.Command{
		Name:  "show",
		Usage: "Show current configuration of your gorg",
		Action: func(c *cli.Context) error {
			config := loadConfig()
			env, err := config.GetEnvironment(config.Default)
			if err != nil {
				return err
			}

			fmt.Printf("Using %s environment by default: \n", color.YellowString(config.Default))
			fmt.Printf("%s", env.Services)
			fmt.Print("\nEnvironments available: \n")
			for _, env := range config.Environments {
				fmt.Printf("  - %s \n", env.Name)
			}

			return nil
		},
	}
}

func addConfigCommand() cli.Command {
	return cli.Command{
		Name:  "add",
		Usage: "Add an environment to current configuration",
		Action: func(c *cli.Context) error {
			args := c.Args()
			if len(args) == 1 {
				config := loadConfig()
				env := service.NewEnvironment(args[0], EnvironmentDomain)
				return config.AddConfigEnvironment(env)
			}
			return errors.New("you need to supply environment name")
		},
	}
}

func removeConfigCommand() cli.Command {
	return cli.Command{
		Name:  "remove",
		Usage: "Remove an environment from your current configuration",
		Action: func(c *cli.Context) error {
			args := c.Args()
			if len(args) == 1 {
				config := loadConfig()
				return config.RemoveConfigEnvironment(args[0])
			}
			return errors.New("you need to supply environment name")
		},
	}
}

func useCommand() cli.Command {
	return cli.Command{
		Name:  "use",
		Usage: "Switch default environment",
		Action: func(c *cli.Context) error {
			args := c.Args()
			if len(args) == 1 {
				config := loadConfig()
				env, err := config.GetEnvironment(args[0])
				if err != nil {
					return err
				}

				err = config.SwitchEnvironment(env)
				if err != nil {
					return err
				}

				CurrentEnvironment = DefaultEnvironment
				return nil
			}
			return errors.New("you need to supply existing environment name")
		},
	}
}
