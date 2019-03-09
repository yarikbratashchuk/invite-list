// Package business holds business domains such as customer and office
// and all related functionality
package business

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"sort"

	"github.com/yarikbratashchuk/invite-list/geo"
)

// Customer represents Intercom customer
type Customer struct {
	ID uint64 `json:"user_id"`

	Name string `json:"name"`

	geo.Coords
}

// Location implements geo.Locator
func (c Customer) Location() geo.Coords { return c.Coords }

// ReadCustomers reads customers from r. Each customer must be
// on a separate line and be in a JSON format
func ReadCustomers(r io.Reader) ([]Customer, error) {
	log.Info("reading customers")

	if r == nil {
		log.Error("reading customers: nil reader")
		return []Customer{}, errors.New("nil reader")
	}

	var err error
	customers := make([]Customer, 0, 20)

	br := bufio.NewScanner(r)
	for br.Scan() {
		var c Customer
		err = json.Unmarshal(br.Bytes(), &c)
		if err != nil {
			log.Errorf("parsing customer: %v", err)
			return []Customer{}, err
		}
		log.Debugf("customer: %v", c)
		customers = append(customers, c)
	}
	if err = br.Err(); err != nil {
		log.Errorf("reading customers: %v", err)
		return []Customer{}, err
	}

	return customers, nil
}

type WriteConfig byte

const (
	ID WriteConfig = 1 << iota
	Verbose
	Padding
)

const DefaultWriteConf = ID | Verbose | Padding

// WriteCustomers writes customers to w
func WriteCustomers(w io.Writer, cs []Customer) {
	WriteCustomersConfig(w, cs, DefaultWriteConf)
}

// WriteCustomersConfig writes customers to w. Use config to configure
// the output.
func WriteCustomersConfig(w io.Writer, cs []Customer, config WriteConfig) {
	log.Info("writing customers")

	if config == 0 {
		config = DefaultWriteConf
	}

	verbose := config&Verbose != 0
	padding := config&Padding != 0

	if padding {
		w.Write([]byte("\n\n"))
	}

	buf := bytes.NewBuffer(make([]byte, 0, 256))
	for _, c := range cs {
		if padding {
			buf.WriteRune('\t')
		}

		if config&ID != 0 {
			if verbose {
				buf.WriteString("ID: ")
			}
			buf.WriteString(fmt.Sprintf("%d", c.ID))
			buf.WriteString("\t\t")
		}
		if verbose {
			buf.WriteString("Name: ")
		}
		buf.WriteString(c.Name)
		buf.WriteString("\n")

		w.Write(buf.Bytes())

		buf.Reset()
	}

	if padding {
		w.Write([]byte("\n\n"))
	}
}

// InviteCustomers returns the list of customers who are closer
// than maxDist meters to the office
func InviteCustomers(cs []Customer, office Office, maxDist uint) []Customer {
	log.Info("inviting customers")

	inviteList := make([]Customer, 0, len(cs))

	for _, c := range cs {
		if geo.Distance(c, office) > float64(maxDist) {
			continue
		}
		inviteList = append(inviteList, c)
	}

	return inviteList
}

type customers []Customer

// Len is part of sort.Interface
func (cs customers) Len() int {
	return len(cs)
}

// Swap is part of sort.Interface.
func (cs customers) Swap(i, j int) {
	cs[i], cs[j] = cs[j], cs[i]
}

// Less is part of sort.Interface.
func (cs customers) Less(i, j int) bool {
	return cs[i].ID < cs[j].ID
}

// SortCustomersByID sorts customers by ID in ascending order
func SortCustomersByID(cs []Customer) {
	log.Info("sorting customers")

	sort.Sort(customers(cs))
}
