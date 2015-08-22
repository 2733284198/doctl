package commands

import (
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/bryanl/doit"
	"github.com/digitalocean/godo"
)

var (
	testDomain     = godo.Domain{Name: "example.com"}
	testDomainList = []godo.Domain{
		testDomain,
	}
)

func TestDomainsCreate(t *testing.T) {
	client := &godo.Client{
		Domains: &doit.DomainsServiceMock{
			CreateFn: func(req *godo.DomainCreateRequest) (*godo.Domain, *godo.Response, error) {
				expected := &godo.DomainCreateRequest{
					Name:      testDomain.Name,
					IPAddress: "127.0.0.1",
				}
				if got := req; !reflect.DeepEqual(got, expected) {
					t.Errorf("CreateFn() called with %#v; expected %#v", got, expected)
				}
				return &testDomain, nil, nil
			},
		},
	}

	withTestClient(client, func(c doit.ViperConfig) {
		c.Set(doit.ArgDomainName, testDomain.Name)
		c.Set(doit.ArgIPAddress, "127.0.0.1")
		RunDomainCreate(ioutil.Discard)
	})
}
