package awsses


// The policy for the receipt filter.
// Experimental.
type ReceiptFilterPolicy string

const (
	// Allow the ip address or range.
	// Experimental.
	ReceiptFilterPolicy_ALLOW ReceiptFilterPolicy = "ALLOW"
	// Block the ip address or range.
	// Experimental.
	ReceiptFilterPolicy_BLOCK ReceiptFilterPolicy = "BLOCK"
)

