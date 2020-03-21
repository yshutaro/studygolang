package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "samplecli"
	app.Usage = "github.com/yshutaro/studygolang"
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name:    "all",
			Aliases: []string{"a"},
			Usage:   "View saved memo.",
			Action: func(c *cli.Context) error {
				fmt.Println("Hello!!")
				return nil
			},
		},
	}

	app.Run(os.Args)

}
