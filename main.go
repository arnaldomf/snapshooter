package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Snap Shooter"
	app.Usage = "Schedule snapshoots for instances"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Value: "config.json",
			Usage: "Set config file (.json)",
		},
	}

	app.Action = func(c *cli.Context) {
		file, err := ioutil.ReadFile(c.String("config"))

		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

		config, err := simplejson.NewJson(file)

		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}

		snapshooter := &Snapshooter{config: config}
		snapshooter.Start()
	}

	app.Run(os.Args)
}
