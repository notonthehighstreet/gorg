package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/notonthehighstreet/gorg/service"
	"github.com/urfave/cli"
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

func loadConfig() service.Config {
	config, err := service.LoadConfig(ConfigFile)
	if err != nil {
		log.Fatal(err)
	}
	return config
}
