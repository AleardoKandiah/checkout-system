package checkout

// ICheckout interface defines the contract for a checkout system.
// It specifies that any type that implements it must have Scan and GetTotalPrice methods.
type ICheckout interface {
    Scan(item string)
    GetTotalPrice() int
}
