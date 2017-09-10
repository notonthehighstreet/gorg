package main

import (
	"os"
	"time"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "gorg"
	app.Version = "v0.0.1"
	app.Compiled = time.Now()
	app.Usage = "CLI to interact with mesos, marathon, chronos and others"
	app.Authors = []cli.Author{
		{
			Name:  "Eggya Chiquita",
			Email: "eggyachiquita@notonthehighstreet.com",
		},
	}
	app.Commands = []cli.Command{
		initCmd,
		configCmd,
		sshUserCmd,
		switchCmd,
		serviceCmd,
	}
	app.Run(os.Args)
}
