package business_test

import (
	"testing"

	"github.com/yarikbratashchuk/invite-list/business"
)

func TestReadCustomers(t *testing.T) {
	// Please open an issue if you want to see this.
	// Currently I don't have time for this. Sorry.
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
