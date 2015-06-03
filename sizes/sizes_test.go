package sizes

import (
	"flag"
	"testing"

	"github.com/bryanl/docli"
	"github.com/codegangsta/cli"
	"github.com/digitalocean/godo"
	"github.com/stretchr/testify/assert"
)

var (
	testSize     = godo.Size{Slug: "small"}
	testSizeList = []godo.Size{testSize}
)

func TestSizesList(t *testing.T) {
	didList := false

	client := &godo.Client{
		Sizes: &docli.SizesServiceMock{
			ListFn: func(opt *godo.ListOptions) ([]godo.Size, *godo.Response, error) {
				didList = true

				resp := &godo.Response{
					Links: &godo.Links{
						Pages: &godo.Pages{},
					},
				}
				return testSizeList, resp, nil
			},
		},
	}

	cs := &docli.TestClientSource{client}
	fs := flag.NewFlagSet("flag set", 0)

	docli.WithinTest(cs, fs, func(c *cli.Context) {
		List(c)
		assert.True(t, didList)
	})
}
