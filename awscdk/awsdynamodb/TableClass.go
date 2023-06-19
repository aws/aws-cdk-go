package awsdynamodb


// DynamoDB's table class.
//
// Example:
//   table := dynamodb.NewTable(this, jsii.String("Table"), &TableProps{
//   	PartitionKey: &Attribute{
//   		Name: jsii.String("id"),
//   		Type: dynamodb.AttributeType_STRING,
//   	},
//   	TableClass: dynamodb.TableClass_STANDARD_INFREQUENT_ACCESS,
//   })
//
// See: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.TableClasses.html
//
// Experimental.
type TableClass string

const (
	// Default table class for DynamoDB.
	// Experimental.
	TableClass_STANDARD TableClass = "STANDARD"
	// Table class for DynamoDB that reduces storage costs compared to existing DynamoDB Standard tables.
	// Experimental.
	TableClass_STANDARD_INFREQUENT_ACCESS TableClass = "STANDARD_INFREQUENT_ACCESS"
)

