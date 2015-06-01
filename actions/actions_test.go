package actions

import (
	"flag"
	"testing"

	"github.com/bryanl/docli/docli"
	"github.com/codegangsta/cli"
	"github.com/digitalocean/godo"
)

var (
	testAction     = godo.Action{ID: 1}
	testActionList = []godo.Action{
		testAction,
	}
)

func TestActionList(t *testing.T) {
	actionDidList := false

	client := &godo.Client{
		Actions: &docli.ActionsServiceMock{
			ListFn: func(opts *godo.ListOptions) ([]godo.Action, *godo.Response, error) {
				actionDidList = true
				resp := &godo.Response{
					Links: &godo.Links{
						Pages: &godo.Pages{},
					},
				}
				return testActionList, resp, nil
			},
		},
	}

	cs := &docli.TestClientSource{client}

	docli.WithinTest(cs, nil, func(c *cli.Context) {
		Action(c)
		if !actionDidList {
			t.Errorf("Action() did not run")
		}
	})
}

func TestActionGet(t *testing.T) {
	client := &godo.Client{
		Actions: &docli.ActionsServiceMock{
			GetFn: func(id int) (*godo.Action, *godo.Response, error) {
				if got, expected := id, testAction.ID; got != expected {
					t.Errorf("GetFn() called with %d; expected %d", got, expected)
				}
				return &testAction, nil, nil
			},
		},
	}

	cs := &docli.TestClientSource{client}
	fs := flag.NewFlagSet("flag set", 0)
	fs.Int("action-id", testAction.ID, "action-id")

	docli.WithinTest(cs, fs, func(c *cli.Context) {
		Get(c)
	})
}
