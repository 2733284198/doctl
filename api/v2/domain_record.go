package apiv2

import (
	"fmt"
	"strings"
)

// id           number  The unique id for the individual record.
// type         string  The DNS record type (A, MX, CNAME, etc).
// name         string  The host name, alias, or service being defined by the record. See the [domain record] object to find out more.
// data         string  Variable data depending on record type. See the [domain record] object for more detail on each record type.
// priority     nullable number The priority of the host (for SRV and MX records. null otherwise).
// port         nullable number The port that the service is accessible on (for SRV records only. null otherwise).
// weight       nullable number The weight of records with the same priority (for SRV records only. null otherwise).
type DomainRecord struct {
	ID       int    `json:"id,omitempty"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	Data     string `json:"data"`
	Priority int    `json:"priority,omitempty"`
	Port     int    `json:"port,omitempty"`
	Weight   int    `json:"weight,omitempty"`
	client   *Client
}

type DomainRecordListResponse struct {
	DomainRecords *DomainRecordList `json:"domain_records"`
	Meta          struct {
		Total int `json:"total"`
	} `json:"meta"`
}

type DomainRecordList []*DomainRecord

type DomainRecordResponse struct {
	DomainRecord *DomainRecord `json:"domain_record"`
}

func (c *Client) NewDomainRecord() *DomainRecord {
	return &DomainRecord{
		client: c,
	}
}

func (c *Client) LoadDomainRecord(name string) (*DomainRecord, error) {
	domain, findErr := c.FindDomainFromName(name)
	if findErr != nil {
		return nil, findErr
	}

	domainRecordList, listErr := c.ListAllDomainRecords(domain.Name)
	if listErr != nil {
		return nil, listErr
	}

	name = strings.Replace(name, fmt.Sprintf(".%s", domain.Name), "", 1)

	for _, domainRecord := range *domainRecordList {
		if domainRecord.Name == name {
			return domainRecord, nil
		}
	}

	return nil, fmt.Errorf("%s Not Found.", name)
}

func (c *Client) ListAllDomainRecords(domain string) (*DomainRecordList, error) {
	domainList := &DomainRecordListResponse{}

	err := c.Get(fmt.Sprintf("domains/%s/records", domain), nil, domainList, nil)
	if err != nil {
		return nil, fmt.Errorf("API Error: %s", err.Message)
	}

	return domainList.DomainRecords, nil
}

func (c *Client) FindDomainFromName(search string) (*Domain, error) {
	domains, err := c.ListAllDomains()
	if err != nil {
		return nil, err
	}

	for _, domain := range domains.Domains {
		if strings.Contains(search, domain.Name) {
			return domain, nil
		}
	}

	return nil, fmt.Errorf("%s Not Found", search)
}

func (c *Client) CreateDomainRecord(domainRecord *DomainRecord, domain *Domain) (*DomainRecord, error) {
	var response DomainRecordResponse

	domain, err := c.FindDomainFromName(domainRecord.Name)
	if err != nil {
		return nil, err
	}

	domainRecord.Name = strings.Replace(domainRecord.Name, fmt.Sprintf(".%s", domain.Name), "", 1)

	apiErr := c.Post(fmt.Sprintf("domains/%s/records", domain.Name), domainRecord, &response, nil)
	if apiErr != nil {
		return nil, fmt.Errorf("API Error: %s", apiErr.Message)
	}

	return response.DomainRecord, nil
}

func (c *Client) DestroyDomainRecord(name string) error {
	domain, err := c.FindDomainFromName(name)
	if err != nil {
		return err
	}

	record, loadErr := c.LoadDomainRecord(name)
	if loadErr != nil {
		return loadErr
	}

	apiErr := c.Delete(fmt.Sprintf("domains/%s/records/%d", domain.Name, record.ID), nil, nil)
	if apiErr != nil {
		return fmt.Errorf("API Error: %s", apiErr.Message)
	}

	return nil
}
