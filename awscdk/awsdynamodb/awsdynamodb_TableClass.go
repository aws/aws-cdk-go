package awsdynamodb


// DynamoDB's table class.
//
// Example:
//   table := dynamodb.NewTable(this, jsii.String("Table"), &tableProps{
//   	partitionKey: &attribute{
//   		name: jsii.String("id"),
//   		type: dynamodb.attributeType_STRING,
//   	},
//   	tableClass: dynamodb.tableClass_STANDARD_INFREQUENT_ACCESS,
//   })
//
// See: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.TableClasses.html
//
type TableClass string

const (
	// Default table class for DynamoDB.
	TableClass_STANDARD TableClass = "STANDARD"
	// Table class for DynamoDB that reduces storage costs compared to existing DynamoDB Standard tables.
	TableClass_STANDARD_INFREQUENT_ACCESS TableClass = "STANDARD_INFREQUENT_ACCESS"
)

