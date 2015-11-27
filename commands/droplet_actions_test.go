package commands

import (
	"io/ioutil"
	"testing"

	"github.com/bryanl/doit"
	"github.com/digitalocean/godo"
)

func TestDropletActionsChangeKernel(t *testing.T) {
	client := &godo.Client{
		DropletActions: &doit.DropletActionsServiceMock{
			ChangeKernelFn: func(id, kernelID int) (*godo.Action, *godo.Response, error) {
				if got, expected := id, 1; got != expected {
					t.Errorf("ChangeKernelFn() id = %d; expected %d", got, expected)
				}
				if got, expected := kernelID, 2; got != expected {
					t.Errorf("ChangeKernelFn() kernelID = %d; expected %d", got, expected)
				}
				return &testAction, nil, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"
		c.Set(ns, doit.ArgKernelID, 2)

		RunDropletActionChangeKernel(ns, c, ioutil.Discard, []string{"1"})
	})
}
func TestDropletActionsDisableBackups(t *testing.T) {
	client := &godo.Client{
		DropletActions: &doit.DropletActionsServiceMock{
			DisableBackupsFn: func(id int) (*godo.Action, *godo.Response, error) {
				if got, expected := id, 1; got != expected {
					t.Errorf("DisableBackupsFn() id = %d; expected %d", got, expected)
				}
				return &testAction, nil, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"

		RunDropletActionDisableBackups(ns, c, ioutil.Discard, []string{"1"})
	})

}
func TestDropletActionsEnableIPv6(t *testing.T) {
	client := &godo.Client{
		DropletActions: &doit.DropletActionsServiceMock{
			EnableIPv6Fn: func(id int) (*godo.Action, *godo.Response, error) {
				if got, expected := id, 1; got != expected {
					t.Errorf("EnableIPv6Fn() id = %d; expected %d", got, expected)
				}
				return &testAction, nil, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"

		RunDropletActionEnableIPv6(ns, c, ioutil.Discard, []string{"1"})
	})
}

func TestDropletActionsEnablePrivateNetworking(t *testing.T) {
	client := &godo.Client{
		DropletActions: &doit.DropletActionsServiceMock{
			EnablePrivateNetworkingFn: func(id int) (*godo.Action, *godo.Response, error) {
				if got, expected := id, 1; got != expected {
					t.Errorf("EnablePrivateNetworkingFn() id = %d; expected %d", got, expected)
				}
				return &testAction, nil, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"

		RunDropletActionEnablePrivateNetworking(ns, c, ioutil.Discard, []string{"1"})
	})
}
func TestDropletActionsGet(t *testing.T) {
	client := &godo.Client{
		DropletActions: &doit.DropletActionsServiceMock{
			GetFn: func(dropletID, actionID int) (*godo.Action, *godo.Response, error) {
				if got, expected := dropletID, 1; got != expected {
					t.Errorf("GetFn() droplet id = %d; expected %d", got, expected)
				}
				if got, expected := actionID, 2; got != expected {
					t.Errorf("GetFn() action id = %d; expected %d", got, expected)
				}
				return &testAction, nil, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"
		c.Set(ns, doit.ArgActionID, 2)

		RunDropletActionGet(ns, c, ioutil.Discard, []string{"1"})
	})
}

func TestDropletActionsPasswordReset(t *testing.T) {
	client := &godo.Client{
		DropletActions: &doit.DropletActionsServiceMock{
			PasswordResetFn: func(id int) (*godo.Action, *godo.Response, error) {
				if got, expected := id, 1; got != expected {
					t.Errorf("PasswordResetFn() id = %d; expected %d", got, expected)
				}
				return &testAction, nil, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"

		RunDropletActionPasswordReset(ns, c, ioutil.Discard, []string{"1"})
	})
}

func TestDropletActionsPowerCycle(t *testing.T) {
	client := &godo.Client{
		DropletActions: &doit.DropletActionsServiceMock{
			PowerCycleFn: func(id int) (*godo.Action, *godo.Response, error) {
				if got, expected := id, 1; got != expected {
					t.Errorf("PowerCycleFn() id = %d; expected %d", got, expected)
				}
				return &testAction, nil, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"

		RunDropletActionPowerCycle(ns, c, ioutil.Discard, []string{"1"})
	})

}
func TestDropletActionsPowerOff(t *testing.T) {
	client := &godo.Client{
		DropletActions: &doit.DropletActionsServiceMock{
			PowerOffFn: func(id int) (*godo.Action, *godo.Response, error) {
				if got, expected := id, 1; got != expected {
					t.Errorf("PowerOffFn() id = %d; expected %d", got, expected)
				}
				return &testAction, nil, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"

		RunDropletActionPowerOff(ns, c, ioutil.Discard, []string{"1"})
	})
}
func TestDropletActionsPowerOn(t *testing.T) {
	client := &godo.Client{
		DropletActions: &doit.DropletActionsServiceMock{
			PowerOnFn: func(id int) (*godo.Action, *godo.Response, error) {
				if got, expected := id, 1; got != expected {
					t.Errorf("PowerOnFn() id = %d; expected %d", got, expected)
				}
				return &testAction, nil, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"

		RunDropletActionPowerOn(ns, c, ioutil.Discard, []string{"1"})
	})

}
func TestDropletActionsReboot(t *testing.T) {
	client := &godo.Client{
		DropletActions: &doit.DropletActionsServiceMock{
			RebootFn: func(id int) (*godo.Action, *godo.Response, error) {
				if got, expected := id, 1; got != expected {
					t.Errorf("RebootFn() id = %d; expected %d", got, expected)
				}
				return &testAction, nil, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"

		RunDropletActionReboot(ns, c, ioutil.Discard, []string{"1"})
	})
}

func TestDropletActionsRebuildByImageID(t *testing.T) {
	client := &godo.Client{
		DropletActions: &doit.DropletActionsServiceMock{
			RebuildByImageIDFn: func(id, imageID int) (*godo.Action, *godo.Response, error) {
				if got, expected := id, 1; got != expected {
					t.Errorf("RebuildByImageIDFn() id = %d; expected %d", got, expected)
				}
				if got, expected := imageID, 2; got != expected {
					t.Errorf("RebuildByImageIDFn() image id = %d; expected %d", got, expected)
				}
				return &testAction, nil, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"
		c.Set(ns, doit.ArgImage, "2")

		RunDropletActionRebuild(ns, c, ioutil.Discard, []string{"1"})
	})
}

func TestDropletActionsRebuildByImageSlug(t *testing.T) {
	client := &godo.Client{
		DropletActions: &doit.DropletActionsServiceMock{
			RebuildByImageSlugFn: func(id int, slug string) (*godo.Action, *godo.Response, error) {
				if got, expected := id, 1; got != expected {
					t.Errorf("RebuildByImageSlugFn() id = %d; expected %d", got, expected)
				}
				if got, expected := slug, "slug"; got != expected {
					t.Errorf("RebuildByImageSlugFn() slug = %q; expected %q", got, expected)
				}
				return &testAction, nil, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"
		c.Set(ns, doit.ArgImage, "slug")

		RunDropletActionRebuild(ns, c, ioutil.Discard, []string{"1"})
	})

}
func TestDropletActionsRename(t *testing.T) {
	client := &godo.Client{
		DropletActions: &doit.DropletActionsServiceMock{
			RenameFn: func(id int, name string) (*godo.Action, *godo.Response, error) {
				if got, expected := id, 1; got != expected {
					t.Errorf("RenameFn() id = %d; expected %d", got, expected)
				}
				if got, expected := name, "name"; got != expected {
					t.Errorf("RenameFn() name = %q; expected %q", got, expected)
				}
				return &testAction, nil, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"
		c.Set(ns, doit.ArgDropletName, "name")

		RunDropletActionRename(ns, c, ioutil.Discard, []string{"1"})
	})
}

func TestDropletActionsResize(t *testing.T) {
	client := &godo.Client{
		DropletActions: &doit.DropletActionsServiceMock{
			ResizeFn: func(id int, slug string, resize bool) (*godo.Action, *godo.Response, error) {
				if got, expected := id, 1; got != expected {
					t.Errorf("ResizeFn() id = %d; expected %d", got, expected)
				}
				if got, expected := slug, "slug"; got != expected {
					t.Errorf("ResizeFn() name = %q; expected %q", got, expected)
				}
				if got, expected := resize, true; got != expected {
					t.Errorf("ResizeFn() resize = %t; expected %t", got, expected)
				}
				return &testAction, nil, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"
		c.Set(ns, doit.ArgImageSlug, "slug")
		c.Set(ns, doit.ArgResizeDisk, true)

		RunDropletActionResize(ns, c, ioutil.Discard, []string{"1"})
	})
}

func TestDropletActionsRestore(t *testing.T) {
	client := &godo.Client{
		DropletActions: &doit.DropletActionsServiceMock{
			RestoreFn: func(id, imageID int) (*godo.Action, *godo.Response, error) {
				if got, expected := id, 1; got != expected {
					t.Errorf("RestoreFn() id = %d; expected %d", got, expected)
				}
				if got, expected := imageID, 2; got != expected {
					t.Errorf("RestoreFn() imageID = %d; expected %d", got, expected)
				}
				return &testAction, nil, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"
		c.Set(ns, doit.ArgImageID, 2)

		RunDropletActionRestore(ns, c, ioutil.Discard, []string{"1"})
	})
}

func TestDropletActionsShutdown(t *testing.T) {
	client := &godo.Client{
		DropletActions: &doit.DropletActionsServiceMock{
			ShutdownFn: func(id int) (*godo.Action, *godo.Response, error) {
				if got, expected := id, 1; got != expected {
					t.Errorf("ShutdownFn() id = %d; expected %d", got, expected)
				}
				return &testAction, nil, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"

		RunDropletActionShutdown(ns, c, ioutil.Discard, []string{"1"})
	})
}

func TestDropletActionsSnapshot(t *testing.T) {
	client := &godo.Client{
		DropletActions: &doit.DropletActionsServiceMock{
			SnapshotFn: func(id int, name string) (*godo.Action, *godo.Response, error) {
				if got, expected := id, 1; got != expected {
					t.Errorf("ShutdownFn() id = %d; expected %d", got, expected)
				}
				if got, expected := name, "name"; got != expected {
					t.Errorf("ShutdownFn() name = %q; expected %q", got, expected)
				}
				return &testAction, nil, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"
		c.Set(ns, doit.ArgSnapshotName, "name")

		RunDropletActionSnapshot(ns, c, ioutil.Discard, []string{"1"})
	})
}

func TestDropletActionsUpgrade(t *testing.T) {
	client := &godo.Client{
		DropletActions: &doit.DropletActionsServiceMock{
			UpgradeFn: func(id int) (*godo.Action, *godo.Response, error) {
				if got, expected := id, 1; got != expected {
					t.Errorf("RenameFn() id = %d; expected %d", got, expected)
				}
				return &testAction, nil, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		ns := "test"

		RunDropletActionUpgrade(ns, c, ioutil.Discard, []string{"1"})
	})
}
