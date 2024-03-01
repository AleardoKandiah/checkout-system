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
	
	// Scenario 2: Testing with 2 A's and 3 B's
	caseTwoAsThreeBs := checkout.NewCheckout()
	twoAsThreeBsItems := []string{"A", "A", "B", "B", "B"}
	for _, item := range twoAsThreeBsItems {
		caseTwoAsThreeBs.Scan(item)
	}
	fmt.Printf("Total price for 2 A's and 3 B's: %d\n", caseTwoAsThreeBs.GetTotalPrice())

	// Test with 4 A's (which should trigger the special price for 3 A's and add the unit price for the extra A)
	caseFourAs := checkout.NewCheckout()
	fourAsItems := []string{"A", "A", "A", "A"}
	for _, item := range fourAsItems {
		caseFourAs.Scan(item)
	}
	fmt.Printf("Total price for 4 A's: %d\n", caseFourAs.GetTotalPrice())

	// Scenario for 3 C's, which have no special pricing but should total to their unit price multiplied by the quantity
	caseThreeCs := checkout.NewCheckout()
	threeCsItems := []string{"C", "C", "C"}
	for _, item := range threeCsItems {
		caseThreeCs.Scan(item)
	}
	fmt.Printf("Total price for 3 C's: %d\n", caseThreeCs.GetTotalPrice())
}