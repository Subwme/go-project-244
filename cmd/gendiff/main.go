package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"code"

	"github.com/urfave/cli/v2"
)

func init() {
	orig := cli.FlagStringer
	cli.FlagStringer = func(f cli.Flag) string {
		result := orig(f)
		if _, ok := f.(*cli.StringFlag); ok {
			result = strings.ReplaceAll(result, " value", " string")
		}
		return result
	}
}

func main() {
	app := &cli.App{
		Name:            "gendiff",
		Usage:           "Compares two configuration files and shows a difference.",
		UsageText:       "gendiff [global options]",
		HideHelpCommand: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "format",
				Aliases: []string{"f"},
				Value:   "stylish",
				Usage:   "output format",
			},
		},
		Action: func(c *cli.Context) error {
			filepath1 := c.Args().Get(0)
			filepath2 := c.Args().Get(1)
			format := c.String("format")

			result, err := code.GenDiff(filepath1, filepath2, format)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
