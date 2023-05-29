package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var debugMode bool = false

func main() {
	app := &cli.App{
		Name: "drop-tools",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "debug",
				Usage:       "Enable debug mode",
				Required:    false,
				Value:       false,
				Destination: &debugMode,
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "raffle",
				Aliases: []string{"r"},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "file",
						Usage:    "File to read",
						Required: true,
						Value:    "entries.csv",
					},
					&cli.IntFlag{
						Name:     "seed",
						Usage:    "Random number seed",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					return doRaffle(c.String("file"), c.Int("seed"))
				},
			},
			{
				Name:  "gen-list",
				Usage: "Generate the list of entries to raffle",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "input",
						Usage:    "List of all registrations for collection",
						Required: true,
						Value:    "registrations.csv",
					},
					&cli.StringFlag{
						Name:     "output",
						Usage:    "File to write verified registrations to",
						Required: false,
						Value:    "verified-registrations.csv",
					},
				},
				Action: func(c *cli.Context) error {
					return optInToRaffleList(c.String("input"), c.String("output"))
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func debugPrintf(s string, args ...interface{}) {
	if debugMode {
		fmt.Printf(s, args...)
	}
}

func debugPrintln(args ...interface{}) {
	if debugMode {
		fmt.Println(args...)
	}
}
