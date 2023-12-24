package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Heathchecker",
		Usage: "A small microservice that checks the current status of a provided website.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "domain",
				Aliases: []string{
					"d",
				},
				Usage:    "Domain name of the website for which you would like to view the current status.",
				Required: true,
			},
			&cli.StringFlag{
				Name: "port",
				Aliases: []string{
					"p",
				},
				Usage:    "Port of the given domain for which you would like to check the current status.",
				Required: false,
			},
		},
		Action: func(c *cli.Context) error {
			port := c.String("port")
			if c.String("port") == "" {
				port = "80"
			}
			status := Check(c.String("domain"), port)
			fmt.Println(status)
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
