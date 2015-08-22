package commands

import (
	"github.com/bryanl/doit"
	"github.com/digitalocean/godo"
)

type testFn func(c doit.ViperConfig)

func withTestClient(client *godo.Client, tFn testFn) {
	ogConfig := doit.VConfig
	defer func() {
		doit.VConfig = ogConfig
	}()

	doit.VConfig = doit.NewTestViperConfig(client)

	tFn(doit.VConfig)
}
