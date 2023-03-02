package awsdynamodb


// Allows you to specify a global secondary index for the global table.
//
// The index will be defined on all replicas.
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
//   	writeProvisionedThroughputSettings: &writeProvisionedThroughputSettingsProperty{
//   		writeCapacityAutoScalingSettings: &capacityAutoScalingSettingsProperty{
//   			maxCapacity: jsii.Number(123),
//   			minCapacity: jsii.Number(123),
//   			targetTrackingScalingPolicyConfiguration: &targetTrackingScalingPolicyConfigurationProperty{
//   				targetValue: jsii.Number(123),
//
//   				// the properties below are optional
//   				disableScaleIn: jsii.Boolean(false),
//   				scaleInCooldown: jsii.Number(123),
//   				scaleOutCooldown: jsii.Number(123),
//   			},
//
//   			// the properties below are optional
//   			seedCapacity: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnGlobalTable_GlobalSecondaryIndexProperty struct {
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
	// Defines write capacity settings for the global secondary index.
	//
	// You must specify a value for this property if the table's `BillingMode` is `PROVISIONED` . All replicas will have the same write capacity settings for this global secondary index.
	WriteProvisionedThroughputSettings interface{} `field:"optional" json:"writeProvisionedThroughputSettings" yaml:"writeProvisionedThroughputSettings"`
}

