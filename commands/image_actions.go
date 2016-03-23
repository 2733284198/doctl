/*
Copyright 2016 The Doctl Authors All rights reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package commands

import (
	"fmt"
	"strconv"

	"github.com/digitalocean/doctl"
	"github.com/digitalocean/doctl/do"
	"github.com/digitalocean/godo"
	"github.com/spf13/cobra"
)

// ImageAction creates the image action commmand.
func ImageAction() *Command {
	cmd := &Command{
		Command: &cobra.Command{
			Use:   "image-action",
			Short: "image-action commands",
			Long:  "image-action commands",
		},
	}

	cmdImageActionsGet := CmdBuilder(cmd, RunImageActionsGet,
		"get <image-id>", "get image action", Writer,
		displayerType(&action{}), docCategories("image"))
	AddIntFlag(cmdImageActionsGet, doit.ArgActionID, 0, "action id", requiredOpt())

	cmdImageActionsTransfer := CmdBuilder(cmd, RunImageActionsTransfer,
		"transfer <image-id>", "transfer image", Writer,
		displayerType(&action{}), docCategories("image"))
	AddStringFlag(cmdImageActionsTransfer, doit.ArgRegionSlug, "", "region", requiredOpt())
	AddBoolFlag(cmdImageActionsTransfer, doit.ArgCommandWait, false, "Wait for action to complete")

	return cmd
}

// RunImageActionsGet retrieves an action for an image.
func RunImageActionsGet(c *CmdConfig) error {
	ias := c.ImageActions()

	if len(c.Args) != 1 {
		return doit.NewMissingArgsErr(c.NS)
	}

	imageID, err := strconv.Atoi(c.Args[0])
	if err != nil {
		return err
	}

	actionID, err := c.Doit.GetInt(c.NS, doit.ArgActionID)
	if err != nil {
		return err
	}

	a, err := ias.Get(imageID, actionID)
	if err != nil {
		return err
	}

	item := &action{actions: do.Actions{*a}}
	return c.Display(item)
}

// RunImageActionsTransfer an image.
func RunImageActionsTransfer(c *CmdConfig) error {
	ias := c.ImageActions()

	if len(c.Args) != 1 {
		return doit.NewMissingArgsErr(c.NS)
	}

	id, err := strconv.Atoi(c.Args[0])
	if err != nil {
		return err
	}

	region, err := c.Doit.GetString(c.NS, doit.ArgRegionSlug)
	if err != nil {
		return err
	}

	req := &godo.ActionRequest{
		"region": region,
	}

	a, err := ias.Transfer(id, req)
	if err != nil {
		checkErr(fmt.Errorf("could not transfer image: %v", err))
	}

	wait, err := c.Doit.GetBool(c.NS, doit.ArgCommandWait)
	if err != nil {
		return err
	}

	if wait {
		a, err = actionWait(c, a.ID, 5)
		if err != nil {
			return err
		}

	}

	item := &action{actions: do.Actions{*a}}
	return c.Display(item)
}
