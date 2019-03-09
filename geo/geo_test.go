package geo_test

import (
	"fmt"
	"testing"

	"github.com/yarikbratashchuk/invite-list/geo"
	"github.com/yarikbratashchuk/invite-list/mock"
)

func TestDistance(t *testing.T) {
	t.Parallel()

	cases := []struct {
		test         string
		loc1, loc2   geo.Locator
		expectedDist float64
		hasErr       bool
	}{{
		test:         "valid case",
		loc1:         mock.Locator1,
		loc2:         mock.Locator2,
		expectedDist: 157249.598,
		hasErr:       false,
	}, {
		test:         "valid case",
		loc1:         mock.Locator2,
		loc2:         mock.Locator1,
		expectedDist: 157249.598,
		hasErr:       false,
	}, {
		test:         "valid case",
		loc1:         mock.Locator1,
		loc2:         mock.Locator3,
		expectedDist: 157249.598,
		hasErr:       false,
	}, {
		test:         "valid case",
		loc1:         mock.Locator2,
		loc2:         mock.Locator3,
		expectedDist: 314499.197,
		hasErr:       false,
	}, {
		test:         "invalid case",
		loc1:         mock.LocatorInv1,
		loc2:         mock.Locator2,
		expectedDist: 0,
		hasErr:       true,
	}, {
		test:         "invalid case",
		loc1:         mock.Locator1,
		loc2:         mock.LocatorInv2,
		expectedDist: 0,
		hasErr:       true,
	}}
	for _, c := range cases {
		c := c
		t.Run(c.test, func(t *testing.T) {
			t.Parallel()

			d, err := geo.Distance(c.loc1, c.loc2)
			if err != nil && !c.hasErr {
				t.Fatalf("unexpected error: %v", err)
			} else if err == nil && c.hasErr {
				t.Fatalf("expected error, got nil")
			}

			if d < 0 {
				t.Errorf("negative distance returned")
			}

			dStr := fmt.Sprintf("%.3f", d)
			expDStr := fmt.Sprintf("%.3f", c.expectedDist)
			if dStr != expDStr {
				t.Errorf("expected: %s, got: %s", expDStr, dStr)
			}
		})
	}
}
