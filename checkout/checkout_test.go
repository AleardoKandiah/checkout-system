package checkout

import (
    "testing"
)

// TestCheckout_ScanAndTotalPrice tests the Scan and GetTotalPrice methods for various scenarios.
func TestCheckout_ScanAndTotalPrice(t *testing.T) {
    tests := []struct {
        name           string
        items          []string
        expectedTotal  int
    }{
        {
            name:          "Single item no special price",
            items:         []string{"C"},
            expectedTotal: 20,
        },
        {
            name:          "Multiple items no special price",
            items:         []string{"C", "D"},
            expectedTotal: 35,
        },
        {
            name:          "Single item with special price not triggered",
            items:         []string{"A", "A"},
            expectedTotal: 100,
        },
        {
            name:          "Single item with special price triggered",
            items:         []string{"A", "A", "A"},
            expectedTotal: 130,
        },
        {
            name:          "Complex scenario with special prices",
            items:         []string{"A", "A", "A", "B", "B", "D"},
            expectedTotal: 190,
        },
        {
            name:          "Items added in non-sequential order",
            items:         []string{"B", "A", "B", "A", "A"},
            expectedTotal: 175,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            co := NewCheckout()
            for _, item := range tt.items {
                co.Scan(item)
            }
            if gotTotal := co.GetTotalPrice(); gotTotal != tt.expectedTotal {
                t.Errorf("%s: expected total %d, got %d", tt.name, tt.expectedTotal, gotTotal)
            }
        })
    }
}
