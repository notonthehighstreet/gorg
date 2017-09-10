package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/fatih/color"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/pkg/browser"
	"github.com/urfave/cli"

	"github.com/notonthehighstreet/gorg/pkg"
)

const (
	errMsgArg   = "missing args"
	errMsgTag   = "tag not supported"
	errExitCode = 86
)

var (
	cfg    *pkg.Config
	consul *pkg.Consul
)

var (
	initCmd = cli.Command{
		Name:  "init",
		Usage: "Initialise gorg configuration file",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "environment-name"},
			cli.StringFlag{Name: "domain"},
		},
		Action: func(c *cli.Context) error {
			domain := c.String("domain")
			if domain == "" {
				return cli.NewExitError(errMsgArg, errExitCode)
			}
			ename := c.String("environment-name")
			if ename == "" {
				return cli.NewExitError(errMsgArg, errExitCode)
			}
			env := pkg.NewEnvironment(ename, domain)
			path, err := filepath()
			if err != nil {
				return err
			}
			cfg = pkg.NewConfig(env, path, domain)
			return updateConfig()
		},
	}
	configCmd = cli.Command{
		Name:  "config",
		Usage: "Modify current configuration of your gorg",
		Subcommands: []cli.Command{
			showConfigCmd,
			addConfigCmd,
			rmConfigCmd,
		},
	}
	showConfigCmd = cli.Command{
		Name:   "show",
		Usage:  "Show current configuration of your gorg",
		Before: func(c *cli.Context) error { return loadConfig() },
		Action: func(c *cli.Context) error {
			env, err := cfg.LoadEnvironment(cfg.Default)
			if err != nil {
				return cli.NewExitError(err, errExitCode)
			}

			fmt.Printf("Using %s environment by default: \n", color.YellowString(cfg.Default))
			fmt.Printf("%s", env.Services)
			fmt.Print("\nEnvironments available: \n")
			for _, env := range cfg.Environments {
				fmt.Printf("  - %s \n", env.Name)
			}
			return nil
		},
	}
	addConfigCmd = cli.Command{
		Name:      "add",
		Usage:     "Add an environment to current configuration",
		ArgsUsage: "[environment-name]",
		Before:    func(c *cli.Context) error { return loadConfig() },
		After:     func(c *cli.Context) error { return updateConfig() },
		Action: func(c *cli.Context) error {
			if c.NArg() < 1 {
				return cli.NewExitError(errMsgArg, errExitCode)
			}
			env := pkg.NewEnvironment(c.Args().Get(0), cfg.Domain)
			return cfg.AddEnvironment(env)
		},
	}
	rmConfigCmd = cli.Command{
		Name:      "remove",
		Usage:     "Remove an environment from your current configuration",
		ArgsUsage: "[environment-name]",
		Before:    func(c *cli.Context) error { return loadConfig() },
		After:     func(c *cli.Context) error { return updateConfig() },
		Action: func(c *cli.Context) error {
			if c.NArg() < 1 {
				return cli.NewExitError(errMsgArg, errExitCode)
			}
			return cfg.RemoveEnvironment(c.Args().Get(0))
		},
	}
	sshUserCmd = cli.Command{
		Name:      "ssh",
		Usage:     "Set provided user for ssh",
		ArgsUsage: "[username]",
		Before:    func(c *cli.Context) error { return loadConfig() },
		Action: func(c *cli.Context) error {
			if c.NArg() < 1 {
				return cli.NewExitError(errMsgArg, errExitCode)
			}
			cfg.ChangeUser(c.Args().Get(0))
			return updateConfig()
		},
	}
	serviceCmd = cli.Command{
		Name:  "service",
		Usage: "Interact with running services according to consul",
		Subcommands: []cli.Command{
			listServiceCmd,
			openServiceCmd,
			showServiceCmd,
		},
	}
	listServiceCmd = cli.Command{
		Name:   "ls",
		Usage:  "List running services according to consul",
		Before: func(c *cli.Context) error { return loadConfig() },
		Action: func(c *cli.Context) error {
			w := new(tabwriter.Writer)
			w.Init(os.Stdout, 50, 10, 4, ' ', 0)
			ctl, err := consul.Services()
			if err != nil {
				return cli.NewExitError(err, errExitCode)
			}
			keys := []string{}
			for k := range ctl {
				keys = append(keys, k)
			}
			sort.Strings(keys)
			for _, key := range keys {
				line := fmt.Sprintf("%s\t%s", color.YellowString(key), strings.Join(ctl[key], ", "))
				fmt.Fprintln(w, line)
			}
			fmt.Fprintln(w)
			w.Flush()
			return nil
		},
	}
	openServiceCmd = cli.Command{
		Name:      "open",
		Usage:     "Open selected pkg listed on Consul in a browser",
		ArgsUsage: "[service-name]",
		Before:    func(c *cli.Context) error { return loadConfig() },
		Action: func(c *cli.Context) error {
			if c.NArg() < 1 {
				return cli.NewExitError(errMsgArg, errExitCode)
			}
			url, err := consul.ServiceURL(c.Args().Get(0), "")
			if err != nil {
				return cli.NewExitError(errMsgTag, errExitCode)
			}
			return browser.OpenURL(url)
		},
	}
	showServiceCmd = cli.Command{
		Name:      "show",
		Usage:     "Show a particular pkg from consul",
		ArgsUsage: "[service-name]",
		Before:    func(c *cli.Context) error { return loadConfig() },
		Action: func(c *cli.Context) error {
			if c.NArg() < 1 {
				return cli.NewExitError(errMsgArg, errExitCode)
			}
			name := c.Args().Get(0)
			services, err := consul.Discover(name)
			if err != nil {
				return err
			}
			for _, srv := range services {
				output, err := json.MarshalIndent(srv, "", "\t")
				if err != nil {
					return err
				}
				formatted := append(output, '\n')
				os.Stdout.Write(formatted)
			}
			return nil
		},
	}
	switchCmd = cli.Command{
		Name:      "use",
		Usage:     "Switch default environment",
		ArgsUsage: "[service-name]",
		Before:    func(c *cli.Context) error { return loadConfig() },
		After:     func(c *cli.Context) error { return updateConfig() },
		Action: func(c *cli.Context) error {
			if c.NArg() < 1 {
				return cli.NewExitError(errMsgArg, errExitCode)
			}
			return cfg.SwitchEnvironment(c.Args().Get(0))
		},
	}
)

func loadConfig() error {
	path, err := filepath()
	if err != nil {
		return err
	}
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(file, &cfg)
	if err != nil {
		return err
	}
	env, err := cfg.LoadEnvironment(cfg.Default)
	if err != nil {
		return err
	}
	consul, err = pkg.NewConsul(env.Services.ConsulUI)
	if err != nil {
		return err
	}
	return nil
}

func updateConfig() error {
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	file, err := os.Create(cfg.Filepath)
	defer file.Close()
	if err != nil {
		return err
	}
	w := bufio.NewWriter(file)
	w.Write(data)
	w.Flush()
	return nil
}

func filepath() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return home + string(os.PathSeparator) + "gorg.json", nil
}
