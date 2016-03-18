package commands

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/bryanl/doit"
	"github.com/bryanl/doit/do"
	"github.com/digitalocean/godo"
	"github.com/spf13/cobra"
)

// Domain creates the domain commands heirarchy.
func Domain() *Command {
	cmd := &Command{
		Command: &cobra.Command{
			Use:   "domain",
			Short: "domain commands",
			Long:  "domain is used to access domain commands",
		},
		DocCategories: []string{"domain"},
		IsIndex:       true,
	}

	cmdDomainCreate := CmdBuilder(cmd, RunDomainCreate, "create <domain>", "create domain", Writer,
		aliasOpt("c"), displayerType(&domain{}), docCategories("domain"))
	AddStringFlag(cmdDomainCreate, doit.ArgIPAddress, "", "IP address", requiredOpt())

	CmdBuilder(cmd, RunDomainList, "list", "list domains", Writer,
		aliasOpt("ls"), displayerType(&domain{}), docCategories("domain"))

	CmdBuilder(cmd, RunDomainGet, "get <domain>", "get domain", Writer,
		aliasOpt("g"), displayerType(&domain{}), docCategories("domain"))

	CmdBuilder(cmd, RunDomainDelete, "delete <domain>", "delete droplet", Writer, aliasOpt("g"))

	cmdRecord := &Command{
		Command: &cobra.Command{
			Use:   "records",
			Short: "domain record commands",
			Long:  "commands for interacting with an individual domain",
		},
	}
	cmd.AddCommand(cmdRecord)

	cmdRecordList := CmdBuilder(cmdRecord, RunRecordList, "list <domain>", "list records", Writer,
		aliasOpt("ls"), displayerType(&domainRecord{}), docCategories("domain"))
	AddStringFlag(cmdRecordList, doit.ArgDomainName, "", "Domain name")

	cmdRecordCreate := CmdBuilder(cmdRecord, RunRecordCreate, "create <domain>", "create record", Writer,
		aliasOpt("c"), displayerType(&domainRecord{}), docCategories("domain"))
	AddStringFlag(cmdRecordCreate, doit.ArgRecordType, "", "Record type")
	AddStringFlag(cmdRecordCreate, doit.ArgRecordName, "", "Record name")
	AddStringFlag(cmdRecordCreate, doit.ArgRecordData, "", "Record data")
	AddIntFlag(cmdRecordCreate, doit.ArgRecordPriority, 0, "Record priority")
	AddIntFlag(cmdRecordCreate, doit.ArgRecordPort, 0, "Record port")
	AddIntFlag(cmdRecordCreate, doit.ArgRecordWeight, 0, "Record weight")

	CmdBuilder(cmdRecord, RunRecordDelete, "delete <domain> <record id...>", "delete record", Writer,
		aliasOpt("d"), docCategories("domain"))

	cmdRecordUpdate := CmdBuilder(cmdRecord, RunRecordUpdate, "update <domain>", "update record", Writer,
		aliasOpt("u"), displayerType(&domainRecord{}), docCategories("domain"))
	AddIntFlag(cmdRecordUpdate, doit.ArgRecordID, 0, "Record ID")
	AddStringFlag(cmdRecordUpdate, doit.ArgRecordType, "", "Record type")
	AddStringFlag(cmdRecordUpdate, doit.ArgRecordName, "", "Record name")
	AddStringFlag(cmdRecordUpdate, doit.ArgRecordData, "", "Record data")
	AddIntFlag(cmdRecordUpdate, doit.ArgRecordPriority, 0, "Record priority")
	AddIntFlag(cmdRecordUpdate, doit.ArgRecordPort, 0, "Record port")
	AddIntFlag(cmdRecordUpdate, doit.ArgRecordWeight, 0, "Record weight")

	return cmd
}

// RunDomainCreate runs domain create.
func RunDomainCreate(c *CmdConfig) error {
	if len(c.Args) != 1 {
		return doit.NewMissingArgsErr(c.NS)
	}
	domainName := c.Args[0]

	ds := c.Domains()

	ipAddress, err := c.Doit.GetString(c.NS, "ip-address")
	if err != nil {
		return err
	}

	req := &godo.DomainCreateRequest{
		Name:      domainName,
		IPAddress: ipAddress,
	}

	d, err := ds.Create(req)
	if err != nil {
		return err
	}

	return c.Display(&domain{domains: do.Domains{*d}})
}

// RunDomainList runs domain create.
func RunDomainList(c *CmdConfig) error {

	ds := c.Domains()

	domains, err := ds.List()
	if err != nil {
		return err
	}

	item := &domain{domains: domains}
	return c.Display(item)
}

// RunDomainGet retrieves a domain by name.
func RunDomainGet(c *CmdConfig) error {
	if len(c.Args) != 1 {
		return doit.NewMissingArgsErr(c.NS)
	}
	id := c.Args[0]

	ds := c.Domains()

	if len(id) < 1 {
		return errors.New("invalid domain name")
	}

	d, err := ds.Get(id)
	if err != nil {
		return err
	}

	item := &domain{domains: do.Domains{*d}}
	return c.Display(item)
}

// RunDomainDelete deletes a domain by name.
func RunDomainDelete(c *CmdConfig) error {
	if len(c.Args) != 1 {
		return doit.NewMissingArgsErr(c.NS)
	}
	name := c.Args[0]

	ds := c.Domains()

	if len(name) < 1 {
		return errors.New("invalid domain name")
	}

	err := ds.Delete(name)
	return err
}

// RunRecordList list records for a domain.
func RunRecordList(c *CmdConfig) error {
	if len(c.Args) != 1 {
		return doit.NewMissingArgsErr(c.NS)
	}
	name := c.Args[0]

	ds := c.Domains()

	if len(name) < 1 {
		return errors.New("domain name is missing")
	}

	list, err := ds.Records(name)
	if err != nil {
		return err
	}

	items := &domainRecord{domainRecords: list}
	return c.Display(items)

}

// RunRecordCreate creates a domain record.
func RunRecordCreate(c *CmdConfig) error {
	if len(c.Args) != 1 {
		return doit.NewMissingArgsErr(c.NS)
	}
	name := c.Args[0]

	ds := c.Domains()

	rType, err := c.Doit.GetString(c.NS, doit.ArgRecordType)
	if err != nil {
		return err
	}

	rName, err := c.Doit.GetString(c.NS, doit.ArgRecordName)
	if err != nil {
		return err
	}

	rData, err := c.Doit.GetString(c.NS, doit.ArgRecordData)
	if err != nil {
		return err
	}

	rPriority, err := c.Doit.GetInt(c.NS, doit.ArgRecordPriority)
	if err != nil {
		return err
	}

	rPort, err := c.Doit.GetInt(c.NS, doit.ArgRecordPort)
	if err != nil {
		return err
	}

	rWeight, err := c.Doit.GetInt(c.NS, doit.ArgRecordWeight)
	if err != nil {
		return err
	}

	drcr := &godo.DomainRecordEditRequest{
		Type:     rType,
		Name:     rName,
		Data:     rData,
		Priority: rPriority,
		Port:     rPort,
		Weight:   rWeight,
	}

	if len(drcr.Type) == 0 {
		return errors.New("record request is missing type")
	}

	r, err := ds.CreateRecord(name, drcr)
	if err != nil {
		return err
	}

	item := &domainRecord{domainRecords: do.DomainRecords{*r}}
	return c.Display(item)

}

// RunRecordDelete deletes a domain record.
func RunRecordDelete(c *CmdConfig) error {
	if len(c.Args) < 2 {
		return doit.NewMissingArgsErr(c.NS)
	}

	domainName, ids := c.Args[0], c.Args[1:]
	if len(ids) < 1 {
		return doit.NewMissingArgsErr(c.NS)
	}

	ds := c.Domains()

	for _, i := range ids {
		id, err := strconv.Atoi(i)
		if err != nil {
			return fmt.Errorf("invalid record id %q", i)
		}

		err = ds.DeleteRecord(domainName, id)
		if err != nil {
			return err
		}
	}

	return nil
}

// RunRecordUpdate updates a domain record.
func RunRecordUpdate(c *CmdConfig) error {
	if len(c.Args) != 1 {
		return doit.NewMissingArgsErr(c.NS)
	}
	domainName := c.Args[0]

	ds := c.Domains()

	recordID, err := c.Doit.GetInt(c.NS, doit.ArgRecordID)
	if err != nil {
		return err
	}

	rType, err := c.Doit.GetString(c.NS, doit.ArgRecordType)
	if err != nil {
		return err
	}

	rName, err := c.Doit.GetString(c.NS, doit.ArgRecordName)
	if err != nil {
		return err
	}

	rData, err := c.Doit.GetString(c.NS, doit.ArgRecordData)
	if err != nil {
		return err
	}

	rPriority, err := c.Doit.GetInt(c.NS, doit.ArgRecordPriority)
	if err != nil {
		return err
	}

	rPort, err := c.Doit.GetInt(c.NS, doit.ArgRecordPort)
	if err != nil {
		return err
	}

	rWeight, err := c.Doit.GetInt(c.NS, doit.ArgRecordWeight)
	if err != nil {
		return err
	}

	drcr := &godo.DomainRecordEditRequest{
		Type:     rType,
		Name:     rName,
		Data:     rData,
		Priority: rPriority,
		Port:     rPort,
		Weight:   rWeight,
	}

	r, err := ds.EditRecord(domainName, recordID, drcr)
	if err != nil {
		return err
	}

	item := &domainRecord{domainRecords: do.DomainRecords{*r}}
	return c.Display(item)
}
