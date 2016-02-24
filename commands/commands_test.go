package commands

import (
	"fmt"
	"io/ioutil"
	"sort"
	"testing"

	"github.com/bryanl/doit"
	"github.com/bryanl/doit/do"
	domocks "github.com/bryanl/doit/do/mocks"
	"github.com/bryanl/doit/pkg/runner"
	"github.com/digitalocean/godo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

var (
	testDroplet = do.Droplet{
		Droplet: &godo.Droplet{
			ID: 1,
			Image: &godo.Image{
				ID:           1,
				Name:         "an-image",
				Distribution: "DOOS",
			},
			Name: "a-droplet",
			Networks: &godo.Networks{
				V4: []godo.NetworkV4{
					{IPAddress: "8.8.8.8", Type: "public"},
					{IPAddress: "172.16.1.2", Type: "private"},
				},
			},
			Region: &godo.Region{
				Slug: "test0",
				Name: "test 0",
			},
		},
	}

	anotherTestDroplet = do.Droplet{
		Droplet: &godo.Droplet{
			ID: 3,
			Image: &godo.Image{
				ID:           1,
				Name:         "an-image",
				Distribution: "DOOS",
			},
			Name: "another-droplet",
			Networks: &godo.Networks{
				V4: []godo.NetworkV4{
					{IPAddress: "8.8.8.9", Type: "public"},
					{IPAddress: "172.16.1.4", Type: "private"},
				},
			},
			Region: &godo.Region{
				Slug: "test0",
				Name: "test 0",
			},
		},
	}

	testPrivateDroplet = do.Droplet{
		Droplet: &godo.Droplet{
			ID: 1,
			Image: &godo.Image{
				ID:           1,
				Name:         "an-image",
				Distribution: "DOOS",
			},
			Name: "a-droplet",
			Networks: &godo.Networks{
				V4: []godo.NetworkV4{
					{IPAddress: "172.16.1.2", Type: "private"},
				},
			},
			Region: &godo.Region{
				Slug: "test0",
				Name: "test 0",
			},
		},
	}

	testDropletList        = do.Droplets{testDroplet, anotherTestDroplet}
	testPrivateDropletList = do.Droplets{testPrivateDroplet}
	testKernel             = do.Kernel{Kernel: &godo.Kernel{ID: 1}}
	testKernelList         = do.Kernels{testKernel}
	testFloatingIP         = do.FloatingIP{
		FloatingIP: &godo.FloatingIP{
			Droplet: testDroplet.Droplet,
			Region:  testDroplet.Region,
			IP:      "127.0.0.1",
		},
	}
	testFloatingIPList = do.FloatingIPs{testFloatingIP}
)

func assertCommandNames(t *testing.T, cmd *cobra.Command, expected ...string) {
	var names []string

	for _, c := range cmd.Commands() {
		names = append(names, c.Name())
	}

	sort.Strings(expected)
	sort.Strings(names)
	assert.Equal(t, expected, names)
}

type testFn func(c *cmdConfig)

type testCmdConfig struct {
	*cmdConfig

	doitConfig *TestConfig
}

func withTestClient(tFn testFn) {
	ogConfig := doit.DoitConfig
	defer func() {
		doit.DoitConfig = ogConfig
	}()

	cfg := NewTestConfig()
	doit.DoitConfig = cfg

	config := &cmdConfig{
		ns:         "test",
		doitConfig: cfg,
		out:        ioutil.Discard,

		ks:   &domocks.KeysService{},
		ss:   &domocks.SizesService{},
		rs:   &domocks.RegionsService{},
		is:   &domocks.ImagesService{},
		ias:  &domocks.ImageActionsService{},
		fis:  &domocks.FloatingIPsService{},
		fias: &domocks.FloatingIPActionsService{},
		ds:   &domocks.DropletsService{},
		das:  &domocks.DropletActionsService{},
		dos:  &domocks.DomainsService{},
		acts: &domocks.ActionsService{},
		as:   &domocks.AccountService{},
	}

	tFn(config)
}

type TestConfig struct {
	SSHFn func(user, host, keyPath string, port int) runner.Runner
	v     *viper.Viper
}

var _ doit.Config = &TestConfig{}

func NewTestConfig() *TestConfig {
	return &TestConfig{
		SSHFn: func(u, h, kp string, p int) runner.Runner {
			return &doit.MockRunner{}
		},
		v: viper.New(),
	}
}

var _ doit.Config = &TestConfig{}

func (c *TestConfig) GetGodoClient() *godo.Client {
	return &godo.Client{}
}

func (c *TestConfig) SSH(user, host, keyPath string, port int) runner.Runner {
	return c.SSHFn(user, host, keyPath, port)
}

func (c *TestConfig) Set(ns, key string, val interface{}) {
	nskey := fmt.Sprintf("%s-%s", ns, key)
	c.v.Set(nskey, val)
}

func (c *TestConfig) GetString(ns, key string) (string, error) {
	nskey := fmt.Sprintf("%s-%s", ns, key)
	return c.v.GetString(nskey), nil
}

func (c *TestConfig) GetInt(ns, key string) (int, error) {
	nskey := fmt.Sprintf("%s-%s", ns, key)
	return c.v.GetInt(nskey), nil
}

func (c *TestConfig) GetStringSlice(ns, key string) ([]string, error) {
	nskey := fmt.Sprintf("%s-%s", ns, key)
	return c.v.GetStringSlice(nskey), nil
}

func (c *TestConfig) GetBool(ns, key string) (bool, error) {
	nskey := fmt.Sprintf("%s-%s", ns, key)
	return c.v.GetBool(nskey), nil
}
