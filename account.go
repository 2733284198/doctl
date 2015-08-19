package doit

import (
	"io"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/digitalocean/godo"
	"github.com/spf13/viper"
)

func AccountGet(c *cli.Context) {
	client := NewClient(c, DefaultConfig)
	err := accountGet(client, c.App.Writer)
	if err != nil {
		log.WithField("err", err).Fatal("could not display account")
	}
}

func NewAccountGet() {
	token := viper.GetString("token")
	client := DefaultConfig.NewClient(token)
	_ = accountGet(client, os.Stdout)
}

func accountGet(client *godo.Client, w io.Writer) error {
	a, _, err := client.Account.Get()
	if err != nil {
		return err
	}

	return writeJSON(a, w)
}
