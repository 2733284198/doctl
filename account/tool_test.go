package account

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/bryanl/docli/docli"
	"github.com/codegangsta/cli"
	"github.com/digitalocean/godo"
)

var testAccount = &godo.Account{
	DropletLimit:  10,
	Email:         "user@example.com",
	UUID:          "1234",
	EmailVerified: true,
}

func TestAccountAction(t *testing.T) {
	accountDidGet := false

	client := &godo.Client{
		Account: &docli.AccountServiceMock{
			GetFn: func() (*godo.Account, *godo.Response, error) {
				accountDidGet = true
				return testAccount, nil, nil
			},
		},
	}

	cs := &docli.TestClientSource{client}

	docli.WithinTest(cs, nil, func(c *cli.Context) {
		Action(c)
		if !accountDidGet {
			t.Errorf("Action() did not run")
		}
	})
}

func TestAccountGet(t *testing.T) {
	client := &godo.Client{
		Account: &docli.AccountServiceMock{
			GetFn: func() (*godo.Account, *godo.Response, error) {
				return testAccount, nil, nil
			},
		},
	}

	var b bytes.Buffer
	w := bufio.NewWriter(&b)

	AccountGet(client, w)
	w.Flush()

	var ar godo.Account
	err := json.Unmarshal(b.Bytes(), &ar)
	if err != nil {
		t.Fatalf("AccountGet() can't unmarshal: %v", err)
	}

	if got, expected := ar, *testAccount; !reflect.DeepEqual(got, expected) {
		t.Errorf("AccountGet() = %#v; expected %#v", got, expected)
	}
}

func TestAccountGet_APIError(t *testing.T) {
	client := &godo.Client{
		Account: &docli.AccountServiceMock{
			GetFn: func() (*godo.Account, *godo.Response, error) {
				return nil, nil, fmt.Errorf("an error")
			},
		},
	}

	var b bytes.Buffer
	w := bufio.NewWriter(&b)

	err := AccountGet(client, w)
	w.Flush()

	if err == nil {
		t.Errorf("AccountGet expected error")
	}

}
