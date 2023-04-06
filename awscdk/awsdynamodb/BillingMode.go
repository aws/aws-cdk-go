package awsdynamodb


// DynamoDB's Read/Write capacity modes.
//
// Example:
//   table := dynamodb.NewTable(this, jsii.String("Table"), &TableProps{
//   	PartitionKey: &Attribute{
//   		Name: jsii.String("id"),
//   		Type: dynamodb.AttributeType_STRING,
//   	},
//   	BillingMode: dynamodb.BillingMode_PAY_PER_REQUEST,
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

