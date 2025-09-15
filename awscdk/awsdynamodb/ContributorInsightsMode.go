package awsdynamodb


// DynamoDB's Contributor Insights Mode.
//
// Example:
//   table := dynamodb.NewTableV2(this, jsii.String("Table"), &TablePropsV2{
//   	PartitionKey: &Attribute{
//   		Name: jsii.String("pk"),
//   		Type: dynamodb.AttributeType_STRING,
//   	},
//   	ContributorInsightsSpecification: &ContributorInsightsSpecification{
//   		Enabled: jsii.Boolean(true),
//   		Mode: dynamodb.ContributorInsightsMode_ACCESSED_AND_THROTTLED_KEYS,
//   	},
//   })
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dynamodb-table-contributorinsightsspecification.html
//
type ContributorInsightsMode string

const (
	// Emits metrics for all read and write requests, whether successful or throttled.
	ContributorInsightsMode_ACCESSED_AND_THROTTLED_KEYS ContributorInsightsMode = "ACCESSED_AND_THROTTLED_KEYS"
	// Emits metrics for read and write requests that were throttled.
	ContributorInsightsMode_THROTTLED_KEYS ContributorInsightsMode = "THROTTLED_KEYS"
)

