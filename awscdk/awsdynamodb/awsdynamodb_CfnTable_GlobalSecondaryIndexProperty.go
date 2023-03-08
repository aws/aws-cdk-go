package awsdynamodb


// Represents the properties of a global secondary index.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   globalSecondaryIndexProperty := &globalSecondaryIndexProperty{
//   	indexName: jsii.String("indexName"),
//   	keySchema: []interface{}{
//   		&keySchemaProperty{
//   			attributeName: jsii.String("attributeName"),
//   			keyType: jsii.String("keyType"),
//   		},
//   	},
//   	projection: &projectionProperty{
//   		nonKeyAttributes: []*string{
//   			jsii.String("nonKeyAttributes"),
//   		},
//   		projectionType: jsii.String("projectionType"),
//   	},
//
//   	// the properties below are optional
//   	contributorInsightsSpecification: &contributorInsightsSpecificationProperty{
//   		enabled: jsii.Boolean(false),
//   	},
//   	provisionedThroughput: &provisionedThroughputProperty{
//   		readCapacityUnits: jsii.Number(123),
//   		writeCapacityUnits: jsii.Number(123),
//   	},
//   }
//
type CfnTable_GlobalSecondaryIndexProperty struct {
	// The name of the global secondary index.
	//
	// The name must be unique among all other indexes on this table.
	IndexName *string `field:"required" json:"indexName" yaml:"indexName"`
	// The complete key schema for a global secondary index, which consists of one or more pairs of attribute names and key types:  - `HASH` - partition key - `RANGE` - sort key  > The partition key of an item is also known as its *hash attribute* .
	//
	// The term "hash attribute" derives from DynamoDB's usage of an internal hash function to evenly distribute data items across partitions, based on their partition key values.
	// >
	// > The sort key of an item is also known as its *range attribute* . The term "range attribute" derives from the way DynamoDB stores items with the same partition key physically close together, in sorted order by the sort key value.
	KeySchema interface{} `field:"required" json:"keySchema" yaml:"keySchema"`
	// Represents attributes that are copied (projected) from the table into the global secondary index.
	//
	// These are in addition to the primary key attributes and index key attributes, which are automatically projected.
	Projection interface{} `field:"required" json:"projection" yaml:"projection"`
	// The settings used to enable or disable CloudWatch Contributor Insights for the specified global secondary index.
	ContributorInsightsSpecification interface{} `field:"optional" json:"contributorInsightsSpecification" yaml:"contributorInsightsSpecification"`
	// Represents the provisioned throughput settings for the specified global secondary index.
	//
	// For current minimum and maximum provisioned throughput values, see [Service, Account, and Table Quotas](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html) in the *Amazon DynamoDB Developer Guide* .
	ProvisionedThroughput interface{} `field:"optional" json:"provisionedThroughput" yaml:"provisionedThroughput"`
}

