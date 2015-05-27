package main

import (
	"encoding/json"

	"code.google.com/p/goauth2/oauth"

	"github.com/codegangsta/cli"
	"github.com/digitalocean/godo"
)

func main() {
	app := cli.NewApp()
	app.Name = "docli"
	app.Usage = "DigitalOcean API CLI"
	app.Version = "0.1.0"
	app.Flags = []cli.Flag{
		tokenFlag(),
	}

	app.Commands = []cli.Command{
		accountCommands(),
		actionCommands(),
		domainCommands(),
		dropletCommands(),
		dropletActionCommands(),
		imageActionCommands(),
		imageCommands(),
		regionCommands(),
		sizeCommands(),
		sshKeyCommands(),
	}

	app.RunAndExitOnError()
}

func tokenFlag() cli.Flag {
	return cli.StringFlag{
		Name:   "token",
		Usage:  "DigitalOcean API V2 Token",
		EnvVar: "DO_TOKEN",
	}
}

func toJSON(item interface{}) (string, error) {
	b, err := json.MarshalIndent(item, "", "  ")
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func newClient(c *cli.Context) *godo.Client {
	token := c.GlobalString("token")
	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: token},
	}

	return godo.NewClient(t.Client())
}
