package commands

import (
	"errors"
	"fmt"
	"io"
	"os/user"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/bryanl/doit"
	"github.com/digitalocean/godo"
	"github.com/spf13/cobra"
)

const (
	sshNoAddress = "could not find droplet address"
)

var (
	errSSHInvalidOptions = fmt.Errorf("neither id or name were supplied")
	sshHostRE            = regexp.MustCompile("^((?P<m1>\\w+)@)?(?P<m2>.*?)(:(?P<m3>\\d+))?$")
)

// SSH creates the ssh commands heirarchy
func SSH() *cobra.Command {
	usr, err := user.Current()
	checkErr(err)

	path := filepath.Join(usr.HomeDir, ".ssh", "id_rsa")

	cmdSSH := cmdBuilder(nil, RunSSH, "ssh <droplet-id | host>", "ssh to droplet", writer)
	addStringFlag(cmdSSH, doit.ArgSSHUser, "root", "ssh user")
	addStringFlag(cmdSSH, doit.ArgsSSHKeyPath, path, "path to private ssh key")
	addIntFlag(cmdSSH, doit.ArgsSSHPort, 22, "port sshd is running on")

	return cmdSSH.Command
}

// RunSSH finds a droplet to ssh to given input parameters (name or id).
func RunSSH(ns string, config doit.Config, out io.Writer, args []string) error {
	client := config.GetGodoClient()

	if len(args) == 0 {
		return doit.NewMissingArgsErr(ns)
	}

	dropletID := args[0]

	if dropletID == "" {
		return doit.NewMissingArgsErr(ns)
	}

	user, err := config.GetString(ns, doit.ArgSSHUser)
	if err != nil {
		return err
	}

	keyPath, err := config.GetString(ns, doit.ArgsSSHKeyPath)
	if err != nil {
		return err
	}

	port, err := config.GetInt(ns, doit.ArgsSSHPort)
	if err != nil {
		return err
	}

	var droplet *godo.Droplet

	if id, err := strconv.Atoi(dropletID); err == nil {
		// dropletID is an integer
		droplet, err = getDropletByID(client, id)
	} else {
		// dropletID is a string
		var droplets []godo.Droplet
		droplets, err := listDroplets(client)
		if err != nil {
			return err
		}

		shi := extractHostInfo(dropletID)

		user = shi.user
		if i, err := strconv.Atoi(shi.port); shi.port != "" && err != nil {
			port = i
		}

		for _, d := range droplets {
			if d.Name == shi.host {
				droplet = &d
				break
			}
			if strconv.Itoa(d.ID) == shi.host {
				droplet = &d
				break
			}
		}

		if droplet == nil {
			return errors.New("could not find droplet")
		}

	}

	if user == "" {
		user = defaultSSHUser(droplet)
	}

	ips := extractDropletIPs(droplet)

	if ips["public"] == "" {
		return errors.New(sshNoAddress)
	}

	runner := config.SSH(user, ips["public"], keyPath, port)
	return runner.Run()
}

func defaultSSHUser(droplet *godo.Droplet) string {
	slug := strings.ToLower(droplet.Image.Slug)
	if strings.Contains(slug, "coreos") {
		return "core"
	}

	return "root"
}

type sshHostInfo struct {
	user string
	host string
	port string
}

func extractHostInfo(in string) sshHostInfo {
	m := sshHostRE.FindStringSubmatch(in)
	r := map[string]string{}
	for i, n := range sshHostRE.SubexpNames() {
		r[n] = m[i]
	}

	return sshHostInfo{
		user: r["m1"],
		host: r["m2"],
		port: r["m3"],
	}
}
