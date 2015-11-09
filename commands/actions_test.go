package commands

import (
	"io/ioutil"
	"testing"

	"github.com/bryanl/doit"
	"github.com/bryanl/doit/Godeps/_workspace/src/github.com/digitalocean/godo"
	"github.com/bryanl/doit/Godeps/_workspace/src/github.com/stretchr/testify/assert"
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
		Actions: &doit.ActionsServiceMock{
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

	withTestClient(client, func(c *TestConfig) {
		ns := "test"
		err := RunCmdActionList(ns, ioutil.Discard)
		assert.NoError(t, err)

		if !actionDidList {
			t.Errorf("Action() did not run")
		}
	})
}

func TestActionGet(t *testing.T) {
	client := &godo.Client{
		Actions: &doit.ActionsServiceMock{
			GetFn: func(id int) (*godo.Action, *godo.Response, error) {
				if got, expected := id, testAction.ID; got != expected {
					t.Errorf("GetFn() called with %d; expected %d", got, expected)
				}
				return &testAction, nil, nil
			},
		},
	}

	withTestClient(client, func(c *TestConfig) {
		c.Set("test", "action-id", testAction.ID)

		err := RunCmdActionGet("test", ioutil.Discard)
		assert.NoError(t, err)
	})
}
