package commands

import (
	"testing"

	"github.com/bryanl/doit"
	"github.com/digitalocean/godo"
	"github.com/stretchr/testify/assert"
)

var (
	testRegion     = godo.Region{Slug: "dev0"}
	testRegionList = []godo.Region{testRegion}
)

func TestRegionCommand(t *testing.T) {
	cmd := Region()
	assert.NotNil(t, cmd)
	assertCommandNames(t, cmd, "list")
}

func TestRegionsList(t *testing.T) {
	didList := false

	client := &godo.Client{
		Regions: &doit.RegionsServiceMock{
			ListFn: func(opt *godo.ListOptions) ([]godo.Region, *godo.Response, error) {
				didList = true

				resp := &godo.Response{
					Links: &godo.Links{
						Pages: &godo.Pages{},
					},
				}
				return testRegionList, resp, nil
			},
		},
	}

	withTestClient(client, func(config *cmdConfig) {
		err := RunRegionList(config)
		assert.True(t, didList)
		assert.NoError(t, err)
	})
}
