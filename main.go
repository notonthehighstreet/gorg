package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/notonthehighstreet/gorg/service"
	"github.com/urfave/cli"
)

var (
	CurrentEnvironment string // CurrentEnvironment the current environment in use

	ConfigFile         = "config.json"  // ConfigFile stores default config filename
	DefaultEnvironment = "integration"  // DefaultEnvironment stores default/initial environment to be initialised
	EnvironmentDomain  = "qa.noths.com" // EnvironmentDomain stores default environment domain for your services
)

func main() {
	app := cli.NewApp()
	app.Name = "gorg"
	app.Version = "v0.0.1"
	app.Compiled = time.Now()
	app.Usage = "CLI to interact with mesos, marathon, chronos and others"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Eggya Chiquita",
			Email: "eggyachiquita@notonthehighstreet.com",
		},
	}

	app.Commands = []cli.Command{
		initCommand(),
		configCommand(),
		sshuserCommand(),
		useCommand(),
	}

	app.Before = func(c *cli.Context) error {
		setCurrentEnvironment()
		return nil
	}

	app.After = func(c *cli.Context) error {
		return nil
	}

	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Fprintf(c.App.Writer, "Unknown command: %q\n", command)
	}

	app.Run(os.Args)
}

func setCurrentEnvironment() {
	CurrentEnvironment = DefaultEnvironment
}

func loadConfig() service.Config {
	config, err := service.LoadConfig(ConfigFile)
	if err != nil {
		log.Fatal(err)
	}
	return config
}
