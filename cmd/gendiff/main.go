package main

import (
	"code/internal/parser"
	"fmt"
	"log"
	"os"
	"strings"

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

			data1, err := parser.Parse(filepath1)
			if err != nil {
				return err
			}
			data2, err := parser.Parse(filepath2)
			if err != nil {
				return err
			}

			fmt.Println(data1, data2, format)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
