package commands

import (
	"errors"
	"io"

	"github.com/bryanl/doit"
	"github.com/bryanl/doit/Godeps/_workspace/src/github.com/digitalocean/godo"
	"github.com/bryanl/doit/Godeps/_workspace/src/github.com/spf13/cobra"
)

// Domain creates the domain commands heirarchy.
func Domain() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "domain",
		Short: "domain commands",
		Long:  "domain is used to access domain commands",
	}

	cmdDomainCreate := cmdBuilder(RunDomainCreate, "create", "create domain", writer, aliasOpt("c"))
	cmd.AddCommand(cmdDomainCreate)
	addStringFlag(cmdDomainCreate, doit.ArgDomainName, "", "Domain name")
	addStringFlag(cmdDomainCreate, doit.ArgIPAddress, "", "IP address")

	cmdDomainList := cmdBuilder(RunDomainList, "list", "list comains", writer, aliasOpt("ls"))
	cmd.AddCommand(cmdDomainList)

	cmdDomainGet := cmdBuilder(RunDomainGet, "get", "get domain", writer, aliasOpt("g"))
	cmd.AddCommand(cmdDomainGet)
	addStringFlag(cmdDomainGet, doit.ArgDomainName, "", "Domain name")

	cmdDomainDelete := cmdBuilder(RunDomainDelete, "delete", "delete droplet", writer, aliasOpt("g"))
	cmd.AddCommand(cmdDomainDelete)
	addStringFlag(cmdDomainDelete, doit.ArgDomainName, "", "Domain name")

	cmdRecord := &cobra.Command{
		Use:   "records",
		Short: "domain record commands",
		Long:  "commands for interacting with an individual domain",
	}
	cmd.AddCommand(cmdRecord)

	cmdRecordList := cmdBuilder(RunRecordList, "list", "list records", writer, aliasOpt("ls"))
	cmdRecord.AddCommand(cmdRecordList)
	addStringFlag(cmdRecordList, doit.ArgDomainName, "", "Domain name")

	cmdRecordCreate := cmdBuilder(RunRecordCreate, "create", "create record", writer, aliasOpt("c"))
	cmdRecord.AddCommand(cmdRecordCreate)
	addStringFlag(cmdRecordCreate, doit.ArgDomainName, "", "Domain name")
	addStringFlag(cmdRecordCreate, doit.ArgRecordType, "", "Record type")
	addStringFlag(cmdRecordCreate, doit.ArgRecordName, "", "Record name")
	addStringFlag(cmdRecordCreate, doit.ArgRecordData, "", "Record data")
	addIntFlag(cmdRecordCreate, doit.ArgRecordPriority, 0, "Record priority")
	addIntFlag(cmdRecordCreate, doit.ArgRecordPort, 0, "Record port")
	addIntFlag(cmdRecordCreate, doit.ArgRecordWeight, 0, "Record weight")

	cmdRecordDelete := cmdBuilder(RunRecordDelete, "delete", "delete record", writer, aliasOpt("d"))
	cmdRecord.AddCommand(cmdRecordDelete)
	addStringFlag(cmdRecordDelete, doit.ArgDomainName, "", "Domain name")
	addIntFlag(cmdRecordDelete, doit.ArgRecordID, 0, "Record ID")

	cmdRecordUpdate := cmdBuilder(RunRecordUpdate, "update", "update record", writer, aliasOpt("u"))
	cmdRecord.AddCommand(cmdRecordUpdate)
	addStringFlag(cmdRecordUpdate, doit.ArgDomainName, "", "Domain name")
	addIntFlag(cmdRecordUpdate, doit.ArgRecordID, 0, "Record ID")
	addStringFlag(cmdRecordUpdate, doit.ArgRecordType, "", "Record type")
	addStringFlag(cmdRecordUpdate, doit.ArgRecordName, "", "Record name")
	addStringFlag(cmdRecordUpdate, doit.ArgRecordData, "", "Record data")
	addIntFlag(cmdRecordUpdate, doit.ArgRecordPriority, 0, "Record priority")
	addIntFlag(cmdRecordUpdate, doit.ArgRecordPort, 0, "Record port")
	addIntFlag(cmdRecordUpdate, doit.ArgRecordWeight, 0, "Record weight")

	return cmd
}

// RunDomainCreate runs domain create.
func RunDomainCreate(ns string, config doit.Config, out io.Writer) error {
	client := config.GetGodoClient()

	domainName, err := config.GetString(ns, "domain-name")
	if err != nil {
		return err
	}

	ipAddress, err := config.GetString(ns, "ip-address")
	if err != nil {
		return err
	}

	req := &godo.DomainCreateRequest{
		Name:      domainName,
		IPAddress: ipAddress,
	}

	d, _, err := client.Domains.Create(req)
	if err != nil {
		return err
	}

	return doit.DisplayOutput(d, out)
}

// RunDomainList runs domain create.
func RunDomainList(ns string, config doit.Config, out io.Writer) error {
	client := config.GetGodoClient()

	f := func(opt *godo.ListOptions) ([]interface{}, *godo.Response, error) {
		list, resp, err := client.Domains.List(opt)
		if err != nil {
			return nil, nil, err
		}

		si := make([]interface{}, len(list))
		for i := range list {
			si[i] = list[i]
		}

		return si, resp, err
	}

	si, err := doit.PaginateResp(f)
	if err != nil {
		return err
	}

	list := make([]godo.Domain, len(si))
	for i := range si {
		list[i] = si[i].(godo.Domain)
	}

	return doit.DisplayOutput(list, out)
}

// RunDomainGet retrieves a domain by name.
func RunDomainGet(ns string, config doit.Config, out io.Writer) error {
	client := config.GetGodoClient()
	id, err := config.GetString(ns, doit.ArgDomainName)
	if err != nil {
		return err
	}

	if len(id) < 1 {
		return errors.New("invalid domain name")
	}

	d, _, err := client.Domains.Get(id)
	if err != nil {
		return err
	}

	return doit.DisplayOutput(d, out)
}

// RunDomainDelete deletes a domain by name.
func RunDomainDelete(ns string, config doit.Config, out io.Writer) error {
	client := config.GetGodoClient()
	name, err := config.GetString(ns, doit.ArgDomainName)
	if err != nil {
		return err
	}

	if len(name) < 1 {
		return errors.New("invalid domain name")
	}

	_, err = client.Domains.Delete(name)
	return err
}

// RunRecordList list records for a domain.
func RunRecordList(ns string, config doit.Config, out io.Writer) error {
	client := config.GetGodoClient()
	name, err := config.GetString(ns, doit.ArgDomainName)
	if err != nil {
		return err
	}

	if len(name) < 1 {
		return errors.New("domain name is missing")
	}

	f := func(opt *godo.ListOptions) ([]interface{}, *godo.Response, error) {
		list, resp, err := client.Domains.Records(name, opt)
		if err != nil {
			return nil, nil, err
		}

		si := make([]interface{}, len(list))
		for i := range list {
			si[i] = list[i]
		}

		return si, resp, err
	}

	si, err := doit.PaginateResp(f)
	if err != nil {
		return err
	}

	list := make([]godo.DomainRecord, len(si))
	for i := range si {
		list[i] = si[i].(godo.DomainRecord)
	}

	return doit.DisplayOutput(list, out)
}

// RunRecordCreate creates a domain record.
func RunRecordCreate(ns string, config doit.Config, out io.Writer) error {
	client := config.GetGodoClient()
	name, err := config.GetString(ns, doit.ArgDomainName)
	if err != nil {
		return err
	}

	rType, err := config.GetString(ns, doit.ArgRecordType)
	if err != nil {
		return err
	}

	rName, err := config.GetString(ns, doit.ArgRecordName)
	if err != nil {
		return err
	}

	rData, err := config.GetString(ns, doit.ArgRecordData)
	if err != nil {
		return err
	}

	rPriority, err := config.GetInt(ns, doit.ArgRecordPriority)
	if err != nil {
		return err
	}

	rPort, err := config.GetInt(ns, doit.ArgRecordPort)
	if err != nil {
		return err
	}

	rWeight, err := config.GetInt(ns, doit.ArgRecordWeight)
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

	r, _, err := client.Domains.CreateRecord(name, drcr)
	if err != nil {
		return err
	}

	return doit.DisplayOutput(r, out)
}

// RunRecordDelete deletes a domain record.
func RunRecordDelete(ns string, config doit.Config, out io.Writer) error {
	client := config.GetGodoClient()
	domainName, err := config.GetString(ns, doit.ArgDomainName)
	if err != nil {
		return err
	}

	recordID, err := config.GetInt(ns, doit.ArgRecordID)
	if err != nil {
		return err
	}

	_, err = client.Domains.DeleteRecord(domainName, recordID)
	return err
}

// RunRecordUpdate updates a domain record.
func RunRecordUpdate(ns string, config doit.Config, out io.Writer) error {
	client := config.GetGodoClient()
	domainName, err := config.GetString(ns, doit.ArgDomainName)
	if err != nil {
		return err
	}

	recordID, err := config.GetInt(ns, doit.ArgRecordID)
	if err != nil {
		return err
	}

	rType, err := config.GetString(ns, doit.ArgRecordType)
	if err != nil {
		return err
	}

	rName, err := config.GetString(ns, doit.ArgRecordName)
	if err != nil {
		return err
	}

	rData, err := config.GetString(ns, doit.ArgRecordData)
	if err != nil {
		return err
	}

	rPriority, err := config.GetInt(ns, doit.ArgRecordPriority)
	if err != nil {
		return err
	}

	rPort, err := config.GetInt(ns, doit.ArgRecordPort)
	if err != nil {
		return err
	}

	rWeight, err := config.GetInt(ns, doit.ArgRecordWeight)
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

	r, _, err := client.Domains.EditRecord(domainName, recordID, drcr)
	if err != nil {
		return err
	}

	return doit.DisplayOutput(r, out)
}
