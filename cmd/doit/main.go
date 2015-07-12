package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/bryanl/doit"
	"github.com/codegangsta/cli"
	"golang.org/x/oauth2"
)

type tokenSource struct {
	AccessToken string
}

func (t *tokenSource) Token() (*oauth2.Token, error) {
	return &oauth2.Token{
		AccessToken: t.AccessToken,
	}, nil
}

func init() {
	logrus.SetOutput(os.Stderr)
	logrus.SetLevel(logrus.InfoLevel)

	doit.Bail = func(err error, msg string) {
		logrus.WithField("err", err).Fatal(msg)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "doit"
	app.Usage = "DigitalOcean Interactive Tool"
	app.Version = "0.4.0"
	app.Flags = []cli.Flag{
		tokenFlag(),
		debugFlag(),
	}

	app.Commands = []cli.Command{
		accountCommands(),
		actionCommands(),
		domainCommands(),
		dropletCommands(),
		dropletActionCommands(),
		imageActionCommands(),
		imageCommands(),
		sshKeyCommands(),
		regionCommands(),
		sizeCommands(),
		sshCommands(),
	}

	app.RunAndExitOnError()
}

func tokenFlag() cli.Flag {
	return cli.StringFlag{
		Name:   "token",
		Usage:  "DigitalOcean API V2 Token",
		EnvVar: "DIGITAL_OCEAN_TOKEN",
	}
}

func debugFlag() cli.Flag {
	return cli.BoolFlag{
		Name:  "debug",
		Usage: "Debug",
	}
}

func jsonFlag() cli.Flag {
	return cli.BoolFlag{
		Name:  doit.ArgDisplayJSON,
		Usage: "display JSON output",
	}
}

func textFlag() cli.Flag {
	return cli.BoolFlag{
		Name:  doit.ArgDisplayText,
		Usage: "display text output",
	}
}
