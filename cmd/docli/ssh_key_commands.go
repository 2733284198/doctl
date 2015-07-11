package main

import (
	"github.com/bryanl/doit"
	"github.com/codegangsta/cli"
)

func sshKeyCommands() cli.Command {
	return cli.Command{
		Name:  "key",
		Usage: "ssh key commands",
		Subcommands: []cli.Command{
			sshKeyList(),
			sshKeyCreate(),
			sshKeyGet(),
			sshKeyUpdate(),
			sshKeyDelete(),
		},
	}
}

func sshKeyList() cli.Command {
	return cli.Command{
		Name:  "list",
		Usage: "list ssh keys",
		Flags: []cli.Flag{
			jsonFlag(),
			textFlag(),
		},
		Action: doit.KeyList,
	}
}

func sshKeyCreate() cli.Command {
	return cli.Command{
		Name:  "create",
		Usage: "create ssh key",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  doit.ArgKeyName,
				Usage: "ssh key name",
			},
			cli.StringFlag{
				Name:  doit.ArgKeyPublicKey,
				Usage: "ssh public key",
			},
			jsonFlag(),
			textFlag(),
		},
		Action: doit.KeyCreate,
	}
}

func sshKeyGet() cli.Command {
	return cli.Command{
		Name:  "get",
		Usage: "get ssh key",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  doit.ArgKey,
				Usage: "ssh key id or fingerprint",
			},
			jsonFlag(),
			textFlag(),
		},
		Action: doit.KeyGet,
	}
}

func sshKeyUpdate() cli.Command {
	return cli.Command{
		Name:  "update",
		Usage: "update ssh key",
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  doit.ArgKey,
				Usage: "ssh key id",
			},
			cli.StringFlag{
				Name:  doit.ArgKeyName,
				Usage: "ssh key name",
			},
			jsonFlag(),
			textFlag(),
		},
		Action: doit.KeyUpdate,
	}
}

func sshKeyDelete() cli.Command {
	return cli.Command{
		Name:  "delete",
		Usage: "delete ssh key",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  doit.ArgKey,
				Usage: "ssh key id or fingerprint",
			},
		},
		Action: doit.KeyDelete,
	}
}
