package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/notonthehighstreet/gorg/service"
	"github.com/urfave/cli"
)

func loadConfig() service.Config {
	config, err := service.LoadConfig(ConfigFile)
	if err != nil {
		log.Fatal(err)
	}
	return config
}

func initCommand() cli.Command {
	return cli.Command{
		Name:  "init",
		Usage: "Initialise gorg configuration file",
		Action: func(c *cli.Context) error {
			return service.NewConfigFile(DefaultEnvironment, EnvironmentDomain)
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
				log.Fatal(err)
			}

			fmt.Printf("Using %s environment by default: \n", color.YellowString(loadConfig().Default))
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
				return loadConfig().AddConfigEnvironment(args[0], EnvironmentDomain)
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
				return loadConfig().RemoveConfigEnvironment(args[0])
			}

			return errors.New("you need to supply environment name")
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
				return loadConfig().ChangeUser(args[0])
			}

			return errors.New("you need to supply username")
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
				err := loadConfig().SwitchEnvironment(args[0])
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
