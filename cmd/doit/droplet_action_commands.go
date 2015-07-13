package main

import (
	"github.com/bryanl/doit"
	"github.com/codegangsta/cli"
)

func dropletActionCommands() cli.Command {
	return cli.Command{
		Name:  "droplet-action",
		Usage: "droplet action commands",
		Subcommands: []cli.Command{
			dropletChangeKernel(),
			dropletDisableBackups(),
			dropletEnableIPv6(),
			dropletEnablePrivateNetworking(),
			dropletPasswordReset(),
			dropletPowerCycle(),
			dropletPowerOff(),
			dropletPowerOn(),
			dropletReboot(),
			dropletRebuild(),
			dropletRename(),
			dropletRestore(),
			dropletShutdown(),
			dropletUpgrade(),
			dropletActionGet(),
		},
	}
}

func dropletDisableBackups() cli.Command {
	fn := doit.DropletActionDisableBackups
	return noArgDropletCommand("disable-backups", "disables backups for droplet", fn)
}

func dropletReboot() cli.Command {
	fn := doit.DropletActionReboot
	return noArgDropletCommand("reboot", "reboot droplet", fn)
}

func dropletPowerCycle() cli.Command {
	fn := doit.DropletActionPowerCycle
	return noArgDropletCommand("power-cycle", "power cyle droplet", fn)
}

func dropletShutdown() cli.Command {
	fn := doit.DropletActionShutdown
	return noArgDropletCommand("shutdown", "shutdown droplet", fn)
}

func dropletPowerOff() cli.Command {
	fn := doit.DropletActionPowerOff
	return noArgDropletCommand("power-off", "power off droplet", fn)
}

func dropletPowerOn() cli.Command {
	fn := doit.DropletActionPowerOn
	return noArgDropletCommand("power-on", "power on droplet", fn)
}

func dropletPasswordReset() cli.Command {
	fn := doit.DropletActionPasswordReset
	return noArgDropletCommand("password-reset", "reset password for droplet", fn)
}

func dropletEnableIPv6() cli.Command {
	fn := doit.DropletActionEnableIPv6
	return noArgDropletCommand("power-on", "enable ipv6 for droplet", fn)
}

func dropletEnablePrivateNetworking() cli.Command {
	fn := doit.DropletActionPasswordReset
	return noArgDropletCommand("private-networking", "enable private networking for droplet", fn)
}

func dropletUpgrade() cli.Command {
	fn := doit.DropletActionUpgrade
	return noArgDropletCommand("upgrade", "upgrade droplet", fn)
}

func dropletRestore() cli.Command {
	return cli.Command{
		Name:  "restore",
		Usage: "restore droplet",
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "id",
				Usage: "droplet id (required)",
			},
			// TODO make this a string, so it can handle slugs
			cli.IntFlag{
				Name:  "image",
				Usage: "image slug or id (required)",
			},
		},
		Action: doit.DropletActionRestore,
	}
}

func dropletResize() cli.Command {
	return cli.Command{
		Name:  "resize",
		Usage: "resize droplet",
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "id",
				Usage: "droplet id (required)",
			},
			cli.StringFlag{
				Name:  "size",
				Usage: "size slug to resize to (required)",
			},
			cli.BoolFlag{
				Name:  "disk",
				Usage: "increase disk size",
			},
		},
		Action: doit.DropletActionResize,
	}
}

func dropletRebuild() cli.Command {
	return cli.Command{
		Name:  "rebuild",
		Usage: "rebuild droplet",
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "id",
				Usage: "droplet id (required)",
			},
			cli.StringFlag{
				Name:  "image",
				Usage: "image slug or image id (required)",
			},
		},
		Action: doit.DropletActionRebuild,
	}
}

func dropletRename() cli.Command {
	return cli.Command{
		Name:  "rename",
		Usage: "rename droplet",
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "id",
				Usage: "droplet id (required)",
			},
			cli.StringFlag{
				Name:  "name",
				Usage: "new name for droplet (required)",
			},
		},
		Action: doit.DropletActionRename,
	}
}

func dropletChangeKernel() cli.Command {
	return cli.Command{
		Name:  "change-kernel",
		Usage: "change kernel for droplet",
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "id",
				Usage: "droplet id (required)",
			},
			cli.IntFlag{
				Name:  "kernel",
				Usage: "new kernel for droplet (required)",
			},
		},
		Action: doit.DropletActionChangeKernel,
	}
}

func dropletSnapshot() cli.Command {
	return cli.Command{
		Name:  "snapshot",
		Usage: "snapshot droplet",
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "id",
				Usage: "droplet id (required)",
			},
			cli.StringFlag{
				Name:  "name",
				Usage: "name for snapshot",
			},
		},
		Action: doit.DropletActionSnapshot,
	}
}

type noArgDropletFn func(c *cli.Context)

func noArgDropletCommand(name, usage string, fn noArgDropletFn) cli.Command {
	return cli.Command{
		Name:  name,
		Usage: usage,
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "id",
				Usage: "droplet id (required)",
			},
		},
		Action: fn,
	}
}

func dropletActionGet() cli.Command {
	return cli.Command{
		Name:  "get",
		Usage: "get droplet action",
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:  "id",
				Usage: "droplet id",
			},
			cli.StringFlag{
				Name:  "action-id",
				Usage: "action id",
			},
		},
		Action: doit.DropletActionGet,
	}
}
