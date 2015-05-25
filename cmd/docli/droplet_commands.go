package main

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/bryanl/docli/droplets"
	"github.com/codegangsta/cli"
)

func dropletCommands() cli.Command {
	return cli.Command{
		Name:  "droplet",
		Usage: "droplet commands",
		Subcommands: []cli.Command{
			dropletList(),
			dropletCreate(),
		},
	}
}

func dropletList() cli.Command {
	return cli.Command{
		Name:  "list",
		Usage: "list droplets",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "json",
				Usage: "return list of droplets as JSON array",
			},
		},
		Action: func(c *cli.Context) {
			token := c.GlobalString("token")
			client := newClient(token)

			list, err := droplets.List(client)
			if err != nil {
				panic(err)
			}
			if c.Bool("json") {
				j, err := toJSON(list)
				if err != nil {
					panic(err)
				}
				fmt.Println(j)
			} else {
				for _, d := range list {
					fmt.Printf("%s\n", droplets.ToText(&d))
				}
			}

		},
	}
}

func dropletCreate() cli.Command {
	return cli.Command{
		Name:  "create",
		Usage: "create droplet",
		Flags: []cli.Flag{

			cli.StringFlag{
				Name:  "name",
				Usage: "droplet name",
			},
			cli.StringFlag{
				Name:  "region",
				Usage: "droplet region",
			},
			cli.StringFlag{
				Name:  "size",
				Usage: "droplet size",
			},
			cli.StringFlag{
				Name:  "image",
				Usage: "droplet image",
			},
			cli.StringSliceFlag{
				Name:  "ssh-keys",
				Value: &cli.StringSlice{},
				Usage: "droplet public SSH keys",
			},
			cli.BoolFlag{
				Name:  "backups",
				Usage: "enable droplet backups",
			},
			cli.BoolFlag{
				Name:  "ipv6",
				Usage: "enable droplet IPv6",
			},
			cli.BoolFlag{
				Name:  "private-networking",
				Usage: "enable droplet private networking",
			},
			cli.StringFlag{
				Name:  "user-data",
				Usage: "droplet name",
			},
		},
		Action: func(c *cli.Context) {
			client := newClient(c.GlobalString("token"))
			cr := &droplets.CreateRequest{
				Name:              c.String("name"),
				Region:            c.String("region"),
				Size:              c.String("size"),
				Image:             c.String("image"),
				SSHKeys:           c.StringSlice("ssh-keys"),
				Backups:           c.Bool("backups"),
				IPv6:              c.Bool("ipv6"),
				PrivateNetworking: c.Bool("private-networking"),
				UserData:          c.String("user-data"),
			}

			root, err := droplets.Create(client, cr)
			if err != nil {
				log.WithField("err", err).Error("unable to create droplet")
				return
			}

			fmt.Printf("created droplet %d\n", root.Droplet.ID)
		},
	}
}
