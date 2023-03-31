package awsdynamodb


// Represents the properties of a local secondary index.
//
// A local secondary index can only be created when its parent table is created.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   localSecondaryIndexProperty := &localSecondaryIndexProperty{
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
//   }
//
type CfnGlobalTable_LocalSecondaryIndexProperty struct {
	// The name of the local secondary index.
	//
	// The name must be unique among all other indexes on this table.
	IndexName *string `field:"required" json:"indexName" yaml:"indexName"`
	// The complete key schema for the local secondary index, consisting of one or more pairs of attribute names and key types:  - `HASH` - partition key - `RANGE` - sort key  > The partition key of an item is also known as its *hash attribute* .
	//
	// The term "hash attribute" derives from DynamoDB's usage of an internal hash function to evenly distribute data items across partitions, based on their partition key values.
	// >
	// > The sort key of an item is also known as its *range attribute* . The term "range attribute" derives from the way DynamoDB stores items with the same partition key physically close together, in sorted order by the sort key value.
	KeySchema interface{} `field:"required" json:"keySchema" yaml:"keySchema"`
	// Represents attributes that are copied (projected) from the table into the local secondary index.
	//
	// These are in addition to the primary key attributes and index key attributes, which are automatically projected.
	Projection interface{} `field:"required" json:"projection" yaml:"projection"`
}

