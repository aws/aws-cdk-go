package awsdynamodb


// DynamoDB's Read/Write capacity modes.
type BillingMode string

const (
	// Pay only for what you use.
	//
	// You don't configure Read/Write capacity units.
	BillingMode_PAY_PER_REQUEST BillingMode = "PAY_PER_REQUEST"
	// Explicitly specified Read/Write capacity units.
	BillingMode_PROVISIONED BillingMode = "PROVISIONED"
)

