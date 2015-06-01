package docli

import (
	"net/url"
	"strconv"

	log "github.com/Sirupsen/logrus"
	"github.com/digitalocean/godo"
)

// Generator is a function that generates the list to be paginated.
type Generator func(*godo.ListOptions) ([]interface{}, *godo.Response, error)

// PaginateResp paginates a Response.
func PaginateResp(gen Generator, opts *Opts) ([]interface{}, error) {
	opt := &godo.ListOptions{}
	list := []interface{}{}

	for {
		items, resp, err := gen(opt)
		if err != nil {
			return nil, err
		}

		for _, i := range items {
			list = append(list, i)
		}

		if uStr := resp.Links.Pages.Next; len(uStr) > 0 {
			u, err := url.Parse(uStr)
			if err != nil {
				return nil, err
			}

			if opts.Debug {
				log.WithFields(log.Fields{
					"page.current": opt.Page,
					"page.per":     opt.PerPage,
				}).Debug("retrieving page")
			}
			pageStr := u.Query().Get("page")
			page, err := strconv.Atoi(pageStr)
			if err != nil {
				return nil, err
			}

			opt.Page = page
			continue
		}

		break
	}

	return list, nil

}
