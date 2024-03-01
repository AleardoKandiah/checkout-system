package checkout

// ICheckout interface defines the contract for a checkout system.
// It specifies that any type that implements it must have Scan and GetTotalPrice methods.
type ICheckout interface {
    Scan(item string)
    GetTotalPrice() int
}

// Checkout struct holds the state of a checkout session.
type Checkout struct {
    // itemCounts maps each item (by its SKU as a string) to the quantity of that item added to the cart.
    itemCounts map[string]int
    // totalPrice stores the running total price of all items in the cart, taking into account any special pricing rules.
    totalPrice int
}

// unitPrices maps each item SKU to its individual unit price.
var unitPrices = map[string]int{
    "A": 50,
    "B": 30, 
    "C": 20, 
    "D": 15, 
}

// specialPrices maps each item SKU to an array representing a special pricing rule.
// The first element is the quantity of items required for the special price, and the second element is the special price itself.
var specialPrices = map[string][2]int{
    "A": [2]int{3, 130}, // Buying 3 of item A costs 130
    "B": [2]int{2, 45},  // Buying 2 of item B costs 45
    //Everything else priced as normal
}	

// creates and returns a new instance of a checkout.
// Initializes the itemCounts map to keep track of the items scanned.
func NewCheckout() *Checkout {
    return &Checkout{
        itemCounts: make(map[string]int),
    }
}

// Scan adds an item to the checkout session.
// It increments the count of the scanned item and recalculates the total price.
func (c *Checkout) Scan(item string) {
    c.itemCounts[item]++ // increment 
    c.recalculateTotal() // recalculate
}

// returns the current total price of all items in the cart,
// includes any discounts from special pricing rules.
func (c *Checkout) GetTotalPrice() int {
    return c.totalPrice
}	