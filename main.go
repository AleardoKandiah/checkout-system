package main

import (
    "fmt"
    "checkout-system/checkout" 
)

func main() {
    // Scenario 1: The original case
    originalCase := checkout.NewCheckout()
    originalItems := []string{"A", "B", "A", "A", "B"}
    for _, item := range originalItems {
        originalCase.Scan(item)
    }
    fmt.Printf("Total price for original case (3 A's, 2 B's): %d\n", originalCase.GetTotalPrice())
}