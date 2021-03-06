package main

import (
	"fmt"
	"os"
	"regexp"
	"text/tabwriter"

	"github.com/urfave/cli"
	"github.com/zalando-incubator/pazuzu"
)

const (
	PazuzufileName   = "Pazuzufile"
	DockerfileName   = "Dockerfile"
	TestSpecFileName = "test.bats"
	directoryOption  = "directory"
)

var cnfGetCmd = cli.Command{
	Name:   "get",
	Usage:  "Get pazuzu configuration",
	Action: getConfig,
}

var cnfSetCmd = cli.Command{
	Name:   "set",
	Usage:  "Set pazuzu configuration",
	Action: setConfig,
}

var cnfHelpCmd = cli.Command{
	Name:   "help",
	Usage:  "Print help on configuration",
	Action: helpConfigs,
}

var cnfListCmd = cli.Command{
	Name:   "list",
	Usage:  "List current effective configuration",
	Action: listConfigs,
}

func setConfig(c *cli.Context) error {
	a := c.Args()
	if len(a) != 2 {
		return pazuzu.ErrTooFewOrManyParameters
	}
	//
	givenPath := a.Get(0)
	givenValRepr := a.Get(1)
	cfg := pazuzu.GetConfig()
	cfgMirror := pazuzu.GetConfigMirror()
	errSet := cfgMirror.SetConfig(givenPath, givenValRepr)
	if errSet == nil {
		// Oh, it's nice.
		_ = cfg.Save()
		return nil
	}
	fmt.Printf("FAIL [%v]\n", errSet)
	return pazuzu.ErrNotFound
}

func getConfig(c *cli.Context) error {
	a := c.Args()
	if len(a) != 1 {
		return pazuzu.ErrTooFewOrManyParameters
	}
	//
	givenPath := a.Get(0)
	cfgMirror := pazuzu.GetConfigMirror()
	repr, err := cfgMirror.GetRepr(givenPath)
	if err == nil {
		fmt.Println(repr)
		return nil
	}
	return pazuzu.ErrNotFound
}

func helpConfigs(c *cli.Context) error {
	cfgMirror := pazuzu.GetConfigMirror()
	fmt.Println("Pazuzu CLI Config related commands:")
	fmt.Println("\tpazuzu config list\t -- Listing of configuration.")
	fmt.Println("\tpazuzu config help\t-- This help documentation.")
	fmt.Println("\tpazuzu config get KEY\t-- Get specific configuration value.")
	fmt.Println("\tpazuzu config set KEY VALUE\t-- Set configuration.")
	fmt.Printf("\nConfiguration keys and its descriptions:\n")
	for _, k := range cfgMirror.GetKeys() {
		help, errHelp := cfgMirror.GetHelp(k)
		if errHelp == nil {
			fmt.Printf("\t%s\t\t%s\n", k, help)
		}
	}
	return nil
}

func listConfigs(c *cli.Context) error {
	cfgMirror := pazuzu.GetConfigMirror()
	for _, k := range cfgMirror.GetKeys() {
		repr, errRepr := cfgMirror.GetRepr(k)
		if errRepr == nil {
			fmt.Printf("%s=%s\n", k, repr)
		}
	}
	return nil
}

var configCmd = cli.Command{
	Name:  "config",
	Usage: "Configure pazuzu",
	// Action: configure,
	Subcommands: []cli.Command{
		cnfGetCmd,
		cnfSetCmd,
		cnfHelpCmd,
		cnfListCmd,
	},
}

var searchCmd = cli.Command{
	Name:      "search",
	Usage:     "search for features in registry",
	ArgsUsage: "[regexp] - Regexp to be used for feature lookup",
	Action: func(c *cli.Context) error {
		sc, err := pazuzu.GetStorageReader(*pazuzu.GetConfig())
		if err != nil {
			return err // TODO: process properly into human-readable message
		}

		arg := c.Args().Get(0)
		searchRegexp, err := regexp.Compile(arg)

		if err != nil {
			return fmt.Errorf("could not process search regexp '%s': %s", arg, err.Error())
		}
		features, err := sc.SearchMeta(searchRegexp)
		if err != nil {
			return fmt.Errorf("could not search for features: %s", err.Error())
		}

		if len(features) == 0 {
			fmt.Println("no features found")
			return nil
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight)
		fmt.Fprintf(w, "Name \tAuthor \tDescription\n")
		for _, f := range features {
			fmt.Fprintf(w, "%s \t%s \t%s\n", f.Name, f.Author, f.Description)
		}

		w.Flush()

		return nil
	},
}

var composeFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "a, add",
		Usage: "Add features from comma-separated list of `FEATURES`",
	},
	cli.StringFlag{
		Name:  "i, init",
		Usage: "Init set of features from comma-separated list of `FEATURES`",
	},
	cli.StringFlag{
		Name:  "d, directory",
		Usage: "Sets destination directory for Docketfile and Pazuzufile to `DESTINATION`",
	},
}

var composeCmd = cli.Command{
	Name:      "compose",
	Usage:     "Compose Pazuzufile and Dockerfile out of the selected features",
	ArgsUsage: " ", // Do not show arguments
	Description: "Compose step takes list of features as input, validates feature dependencies" +
		" and creates both Pazuzufile and Dockerfile.",
	Flags:  composeFlags,
	Action: composeAction,
}

var buildFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "d, directory",
		Usage: "Sets source path where Docketfile are located.",
	},
	cli.StringFlag{
		Name:  "n, name",
		Usage: "Sets a name for docker image",
	},
}

var buildCmd = cli.Command{
	Name:      "build",
	Usage:     "Builds and tests Docker image from Dockerfile",
	ArgsUsage: " ",
	Flags:     buildFlags,
	Action:    buildFeatures,
}
