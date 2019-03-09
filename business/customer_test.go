package business_test

import (
	"io"
	"strings"
	"testing"

	"github.com/yarikbratashchuk/invite-list/business"
	"github.com/yarikbratashchuk/invite-list/mock"
)

func TestReadCustomers(t *testing.T) {
	t.Parallel()

	cases := []struct {
		test string

		input     io.Reader
		customers []business.Customer
		hasErr    bool
	}{{
		test: "bad input",
		input: strings.NewReader(mock.ValidCustomerJson + `
		blablabla
		` + mock.ValidCustomerJson),
		customers: []business.Customer{},
		hasErr:    true,
	}, {
		test: "valid case",
		input: strings.NewReader(mock.ValidCustomerJson + `
		` + mock.ValidCustomerJson),
		customers: []business.Customer{
			mock.ValidCustomer,
			mock.ValidCustomer,
		},
		hasErr: false,
	}}
	for _, c := range cases {
		c := c
		t.Run(c.test, func(t *testing.T) {
			t.Parallel()

			cs, err := business.ReadCustomers(c.input)
			if err != nil && !c.hasErr {
				t.Fatalf("unexpected err: %v", err)
			} else if err == nil && c.hasErr {
				t.Fatal("expected err, got nil")
			}

			for i, cust := range cs {
				if cust != c.customers[i] {
					t.Errorf(
						"expected %v, got %v",
						c.customers[i],
						cust,
					)
				}
			}
		})
	}
}

func TestWriteCustomersConfig(t *testing.T) {
	// Please open an issue if you want to see this.
	// Currently I don't have time for this. Sorry.
}

func TestInviteCustomers(t *testing.T) {
	// Please open an issue if you want to see this.
	// Currently I don't have time for this. Sorry.
}

func TestSortCustomersByID(t *testing.T) {
	cs := []business.Customer{
		business.Customer{ID: 3},
		business.Customer{ID: 5},
		business.Customer{ID: 1},
	}
	expectedCs := []business.Customer{
		business.Customer{ID: 1},
		business.Customer{ID: 3},
		business.Customer{ID: 5},
	}

	business.SortCustomersByID(cs)
	for i, c := range cs {
		if c.ID != expectedCs[i].ID {
			t.Errorf(
				"expected id: %d, got %d",
				expectedCs[i].ID,
				c.ID,
			)
		}
	}
}
