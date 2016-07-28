package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/notonthehighstreet/gorg/service"
	"github.com/pkg/browser"
	"github.com/urfave/cli"
)

var (
	ConfigFile         = "config.json"  // ConfigFile stores default config filename
	DefaultEnvironment = "integration"  // DefaultEnvironment stores default/initial environment to be initialised
	EnvironmentDomain  = "qa.noths.com" // EnvironmentDomain stores default environment domain for your services
)

func initCommand() cli.Command {
	return cli.Command{
		Name:  "init",
		Usage: "Initialise gorg configuration file",
		Action: func(c *cli.Context) error {
			env := service.NewEnvironment(DefaultEnvironment, EnvironmentDomain)
			err := service.NewConfig(env, ConfigFile)
			if err != nil {
				fmt.Fprintf(c.App.Writer, "Error: %s", err)
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
				err := loadConfig().ChangeUser(args[0])
				if err != nil {
					fmt.Fprintf(c.App.Writer, "Error: %s", err)
					return err
				}
				return nil
			}

			msg := "you need to supply username"
			fmt.Fprintf(c.App.Writer, "Error: %s", msg)
			return errors.New(msg)
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
				fmt.Fprintf(c.App.Writer, "Error: %s", err)
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
				env := service.NewEnvironment(args[0], EnvironmentDomain)
				err := loadConfig().AddConfigEnvironment(env)
				if err != nil {
					fmt.Fprintf(c.App.Writer, "Error: %s", err)
				}
				return nil
			}
			msg := "you need to supply environment name"
			fmt.Fprintf(c.App.Writer, "Error: %s", msg)
			return errors.New(msg)
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
				err := loadConfig().RemoveConfigEnvironment(args[0])
				if err != nil {
					fmt.Fprintf(c.App.Writer, "Error: %s", err)
				}
				return nil
			}
			msg := "you need to supply environment name"
			fmt.Fprintf(c.App.Writer, "Error: %s", msg)
			return errors.New(msg)
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
				err := loadConfig().UseEnvironment(args[0])
				if err != nil {
					fmt.Fprintf(c.App.Writer, "Error: %s", err)
				}
				return nil
			}
			msg := "you need to supply existing environment name"
			fmt.Fprintf(c.App.Writer, "Error: %s", msg)
			return errors.New(msg)
		},
	}
}

func serviceCommand() cli.Command {
	config := loadConfig()
	env, _ := config.GetEnvironment(config.Default)
	consul := service.NewConsul(env.Services.ConsulUI)

	return cli.Command{
		Name:  "service",
		Usage: "Interact with running services according to consul",
		Subcommands: []cli.Command{
			listServiceCommand(consul),
			openServiceCommand(consul),
			showServiceCommand(consul),
		},
	}
}

func listServiceCommand(consul *service.Consul) cli.Command {
	return cli.Command{
		Name:  "ls",
		Usage: "List running services according to consul",
		Action: func(c *cli.Context) error {
			consul.ListServices()
			return nil
		},
	}
}

func openServiceCommand(consul *service.Consul) cli.Command {
	return cli.Command{
		Name:  "open",
		Usage: "Open selected service listed on Consul in a browser",
		Action: func(c *cli.Context) error {
			args := c.Args()
			if len(args) == 1 {
				url := consul.OpenService(args[0])
				browser.OpenURL(url)
				return nil
			}
			msg := "you need to supply service name"
			fmt.Fprintf(c.App.Writer, "Error: %s", msg)
			return errors.New(msg)
		},
	}
}

func showServiceCommand(consul *service.Consul) cli.Command {
	return cli.Command{
		Name:  "show",
		Usage: "Show a particular service from consul",
		Action: func(c *cli.Context) error {
			args := c.Args()
			if len(args) == 1 {
				consul.ShowService(args[0])
				return nil
			}
			msg := "you need to supply service name"
			fmt.Fprintf(c.App.Writer, "Error: %s", msg)
			return errors.New(msg)
		},
	}
}

func loadConfig() service.Config {
	config, err := service.LoadConfig(ConfigFile)
	if err != nil {
		log.Fatal(err)
	}
	return config
}
