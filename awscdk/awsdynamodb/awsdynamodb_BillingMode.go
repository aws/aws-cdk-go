package awsdynamodb


// DynamoDB's Read/Write capacity modes.
//
// Example:
//   table := dynamodb.NewTable(this, jsii.String("Table"), &tableProps{
//   	partitionKey: &attribute{
//   		name: jsii.String("id"),
//   		type: dynamodb.attributeType_STRING,
//   	},
//   	billingMode: dynamodb.billingMode_PAY_PER_REQUEST,
//   })
//
type BillingMode string

const (
	// Pay only for what you use.
	//
	// You don't configure Read/Write capacity units.
	BillingMode_PAY_PER_REQUEST BillingMode = "PAY_PER_REQUEST"
	// Explicitly specified Read/Write capacity units.
	BillingMode_PROVISIONED BillingMode = "PROVISIONED"
)

