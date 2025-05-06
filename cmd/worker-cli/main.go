package main

import (
	"log"
	"os"

	"github.com/LesterCerioli/worker-generator/pkg/generator"
	"github.com/urfave/cli/v2" // Instead of "github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name:  "worker-cli",
		Usage: "Generate Go worker templates",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "name",
				Aliases:  []string{"n"}, // Now valid in v2
				Usage:    "Project name",
				Required: true,
			},
			&cli.StringFlag{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "Go version (e.g., 1.24)",
				Value:   "1.24",
			},
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "Output directory",
				Value:   "./",
			},
		},
		Action: func(c *cli.Context) error {
			config := generator.WorkerConfig{
				ProjectName: c.String("name"),
				GoVersion:   c.String("version"),
			}
			return generator.GenerateWorker(config, c.String("output"))
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
