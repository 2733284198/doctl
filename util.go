package doit

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/digitalocean/godo"
	"golang.org/x/oauth2"
)

// DefaultConfig is the current configuration for the commands.
var DefaultConfig Config = &LiveConfig{}

type TokenSource struct {
	AccessToken string
}

type Runner interface {
	Run() error
}

// TestConfig is an implemenation of Config that can be inspected during tests.
type TestConfig struct {
	Client *godo.Client
	SSHFn  func(user, host string, options []string) Runner
}

type mockRunner struct {
	err error
}

func (tr *mockRunner) Run() error {
	return tr.err
}

// NewTestConfig creates a TestConfig.
func NewTestConfig(client *godo.Client) *TestConfig {
	return &TestConfig{
		Client: client,
		SSHFn: func(u, h string, o []string) Runner {
			logrus.WithFields(logrus.Fields{
				"user":    u,
				"host":    h,
				"options": o,
			}).Info("ssh")
			return &mockRunner{}
		},
	}
}

// NewClient returns the specified godo.Client.
func (cs *TestConfig) NewClient(_ string) *godo.Client {
	return cs.Client
}

// SSH allows the developer to inspect the status of the ssh connection during tests.
func (cs *TestConfig) SSH(user, host string, options []string) Runner {
	return cs.SSHFn(user, host, options)
}

func (t *TokenSource) Token() (*oauth2.Token, error) {
	return &oauth2.Token{
		AccessToken: t.AccessToken,
	}, nil
}

// Config holds configuration values for commands. It currently contains a godo Client
// and a method for running SSH.
type Config interface {
	NewClient(token string) *godo.Client
	SSH(user, host string, options []string) Runner
}

// LiveConfig
type LiveConfig struct{}

// NewClient creates creates a godo.Client givent a token.
func (cs *LiveConfig) NewClient(token string) *godo.Client {
	tokenSource := &TokenSource{
		AccessToken: token,
	}

	oauthClient := oauth2.NewClient(oauth2.NoContext, tokenSource)
	return godo.NewClient(oauthClient)
}

// SSH runs the ssh binary given a user and a host. It preserves stdin, stdout, and stderr.
func (cs *LiveConfig) SSH(user, host string, options []string) Runner {
	logrus.WithFields(logrus.Fields{
		"user": user,
		"host": host,
	}).Info("ssh")

	sshHost := fmt.Sprintf("%s@%s", user, host)

	args := []string{sshHost}
	for _, o := range options {
		args = append(args, "-o", o)
	}

	cmd := exec.Command("ssh", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd
}

// NewClient creates a Client.
func NewClient(c *cli.Context, cs Config) *godo.Client {
	if cs == nil {
		cs = &LiveConfig{}
	}

	pat := c.GlobalString("token")
	return cs.NewClient(pat)
}

func withinTest(cs Config, fs *flag.FlagSet, fn func(*cli.Context)) {
	ogSource := DefaultConfig
	DefaultConfig = cs

	defer func() {
		DefaultConfig = ogSource
	}()

	var b bytes.Buffer
	app := cli.NewApp()
	app.Writer = bufio.NewWriter(&b)

	globalSet := flag.NewFlagSet("global test", 0)
	globalSet.String("token", "token", "token")

	globalCtx := cli.NewContext(app, globalSet, nil)

	if fs == nil {
		fs = flag.NewFlagSet("local test", 0)
	}

	c := cli.NewContext(app, fs, globalCtx)

	fn(c)
}

func ErrWithUsage(c *cli.Context, msg string) {
	logrus.Error(msg)
	cli.ShowCommandHelp(c, c.Command.Name)
}

func bailFatal(err error, msg string) {
	logrus.WithField("err", err).Fatal(msg)
}

func extractDropletPublicIP(droplet *godo.Droplet) string {
	for _, in := range droplet.Networks.V4 {
		if in.Type == "public" {
			return in.IPAddress
		}
	}

	return ""

}
