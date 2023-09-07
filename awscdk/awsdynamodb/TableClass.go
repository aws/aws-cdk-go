package awsdynamodb


// DynamoDB's table class.
//
// Example:
//   table := dynamodb.NewTableV2(this, jsii.String("Table"), &TablePropsV2{
//   	PartitionKey: &Attribute{
//   		Name: jsii.String("pk"),
//   		Type: dynamodb.AttributeType_STRING,
//   	},
//   	ContributorInsights: jsii.Boolean(true),
//   	TableClass: dynamodb.TableClass_STANDARD_INFREQUENT_ACCESS,
//   	PointInTimeRecovery: jsii.Boolean(true),
//   })
//
// See: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.TableClasses.html
//
type TableClass string

const (
	// Default table class for DynamoDB.
	TableClass_STANDARD TableClass = "STANDARD"
	// Table class for DynamoDB that reduces storage costs compared to existing DynamoDB standard tables.
	TableClass_STANDARD_INFREQUENT_ACCESS TableClass = "STANDARD_INFREQUENT_ACCESS"
)

