package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

func main() {
	var environment string
	var platform string

	app := cli.NewApp()
	app.Name = "frdm"
	app.Usage = "your favourite OSS Content Management API"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "env",
			Value:       "local",
			Usage:       "the environment",
			Destination: &environment,
		},
		cli.StringFlag{
			Name:        "platform",
			Value:       "kubernetes",
			Usage:       "the deployment platform",
			Destination: &platform,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "deploy",
			Aliases: []string{"d"},
			Usage:   "deploy frdm",
			Action: func(c *cli.Context) error {
				if platform == "kubernetes" {
					configFileLocation, err := configFileLocation(environment, platform)
					if err != nil {
						fmt.Println(err)
						return nil
					}
					cmdStr := fmt.Sprintf("kubectl create -f %s", configFileLocation)
					out, err := exec.Command("/bin/bash", "-c", cmdStr).Output()
					if err != nil {
						fmt.Println(err)
						return nil
					}
					fmt.Println(out)
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func configFileLocation(environment string, platform string) (string, error) {
	return fmt.Sprintf("./deployments/%s.%s.yaml", environment, platform), nil
}
