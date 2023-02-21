package awsses


// The policy for the receipt filter.
type ReceiptFilterPolicy string

const (
	// Allow the ip address or range.
	ReceiptFilterPolicy_ALLOW ReceiptFilterPolicy = "ALLOW"
	// Block the ip address or range.
	ReceiptFilterPolicy_BLOCK ReceiptFilterPolicy = "BLOCK"
)

