package commands

import (
	"io"

	"github.com/bryanl/doit"
	"github.com/bryanl/doit/Godeps/_workspace/src/github.com/Sirupsen/logrus"
	"github.com/bryanl/doit/Godeps/_workspace/src/github.com/digitalocean/godo"
	"github.com/bryanl/doit/Godeps/_workspace/src/github.com/spf13/cobra"
)

// ImageAction creates the image action commmand.
func ImageAction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "image-action",
		Short: "image-action commands",
		Long:  "image-action commands",
	}

	cmdImageActionsGet := cmdBuilder(RunImageActionsGet,
		"get", "get image action", writer)
	cmd.AddCommand(cmdImageActionsGet)
	addIntFlag(cmdImageActionsGet, doit.ArgImageID, 0, "image id")
	addIntFlag(cmdImageActionsGet, doit.ArgActionID, 0, "action id")

	cmdImageActionsTransfer := cmdBuilder(RunImageActionsTransfer,
		"transfer", "transfer imagr", writer)
	cmd.AddCommand(cmdImageActionsTransfer)
	addIntFlag(cmdImageActionsTransfer, doit.ArgImageID, 0, "image id")
	addStringFlag(cmdImageActionsTransfer, doit.ArgRegionSlug, "", "region")

	return cmd
}

// RunImageActionsGet retrieves an action for an image.
func RunImageActionsGet(ns string, out io.Writer) error {
	client := doit.DoitConfig.GetGodoClient()
	imageID := doit.DoitConfig.GetInt(ns, doit.ArgImageID)
	actionID := doit.DoitConfig.GetInt(ns, doit.ArgActionID)

	action, _, err := client.ImageActions.Get(imageID, actionID)
	if err != nil {
		return err
	}

	return doit.DisplayOutput(action, out)
}

// RunImageActionsTransfer an image.
func RunImageActionsTransfer(ns string, out io.Writer) error {
	client := doit.DoitConfig.GetGodoClient()
	id := doit.DoitConfig.GetInt(ns, doit.ArgImageID)
	req := &godo.ActionRequest{
		"region": doit.DoitConfig.GetString(ns, doit.ArgRegionSlug),
	}

	action, _, err := client.ImageActions.Transfer(id, req)
	if err != nil {
		logrus.WithField("err", err).Fatal("could not transfer image")
	}

	return doit.DisplayOutput(action, out)
}
