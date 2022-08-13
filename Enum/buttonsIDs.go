package Enum

const (
	// Sendor
	Vendor = iota + 1
	Supervisor
	Franchisee
	// Yes/No
	Yes
	No
	// Email request
	Email
)

const (
	// Vendors
	NewVendor = iota + 10
	ExistingVendor
)

const (
	// Current Vendors
	AdDocumentation = iota + 100
	Discount
	PriceChange
	PriceList
	VAT
)

const (
	// New venders expectations => 10
	ContractTerms = iota + 200
	VenderRegistration
	ItemRegistration
	Location
	CallArrangement
	VisitArrangement
	VendorComplain
)
