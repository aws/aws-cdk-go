package previewawsdynamodbmixins


// Represents the properties of a global secondary index.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   globalSecondaryIndexProperty := &GlobalSecondaryIndexProperty{
//   	ContributorInsightsSpecification: &ContributorInsightsSpecificationProperty{
//   		Enabled: jsii.Boolean(false),
//   		Mode: jsii.String("mode"),
//   	},
//   	IndexName: jsii.String("indexName"),
//   	KeySchema: []interface{}{
//   		&KeySchemaProperty{
//   			AttributeName: jsii.String("attributeName"),
//   			KeyType: jsii.String("keyType"),
//   		},
//   	},
//   	OnDemandThroughput: &OnDemandThroughputProperty{
//   		MaxReadRequestUnits: jsii.Number(123),
//   		MaxWriteRequestUnits: jsii.Number(123),
//   	},
//   	Projection: &ProjectionProperty{
//   		NonKeyAttributes: []*string{
//   			jsii.String("nonKeyAttributes"),
//   		},
//   		ProjectionType: jsii.String("projectionType"),
//   	},
//   	ProvisionedThroughput: &ProvisionedThroughputProperty{
//   		ReadCapacityUnits: jsii.Number(123),
//   		WriteCapacityUnits: jsii.Number(123),
//   	},
//   	WarmThroughput: &WarmThroughputProperty{
//   		ReadUnitsPerSecond: jsii.Number(123),
//   		WriteUnitsPerSecond: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-globalsecondaryindex.html
//
type CfnTablePropsMixin_GlobalSecondaryIndexProperty struct {
	// The settings used to specify whether to enable CloudWatch Contributor Insights for the global table and define which events to monitor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-globalsecondaryindex.html#cfn-dynamodb-table-globalsecondaryindex-contributorinsightsspecification
	//
	ContributorInsightsSpecification interface{} `field:"optional" json:"contributorInsightsSpecification" yaml:"contributorInsightsSpecification"`
	// The name of the global secondary index.
	//
	// The name must be unique among all other indexes on this table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-globalsecondaryindex.html#cfn-dynamodb-table-globalsecondaryindex-indexname
	//
	IndexName *string `field:"optional" json:"indexName" yaml:"indexName"`
	// The complete key schema for a global secondary index, which consists of one or more pairs of attribute names and key types:  - `HASH` - partition key - `RANGE` - sort key  > The partition key of an item is also known as its *hash attribute* .
	//
	// The term "hash attribute" derives from DynamoDB's usage of an internal hash function to evenly distribute data items across partitions, based on their partition key values.
	// >
	// > The sort key of an item is also known as its *range attribute* . The term "range attribute" derives from the way DynamoDB stores items with the same partition key physically close together, in sorted order by the sort key value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-globalsecondaryindex.html#cfn-dynamodb-table-globalsecondaryindex-keyschema
	//
	KeySchema interface{} `field:"optional" json:"keySchema" yaml:"keySchema"`
	// The maximum number of read and write units for the specified global secondary index.
	//
	// If you use this parameter, you must specify `MaxReadRequestUnits` , `MaxWriteRequestUnits` , or both. You must use either `OnDemandThroughput` or `ProvisionedThroughput` based on your table's capacity mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-globalsecondaryindex.html#cfn-dynamodb-table-globalsecondaryindex-ondemandthroughput
	//
	OnDemandThroughput interface{} `field:"optional" json:"onDemandThroughput" yaml:"onDemandThroughput"`
	// Represents attributes that are copied (projected) from the table into the global secondary index.
	//
	// These are in addition to the primary key attributes and index key attributes, which are automatically projected.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-globalsecondaryindex.html#cfn-dynamodb-table-globalsecondaryindex-projection
	//
	Projection interface{} `field:"optional" json:"projection" yaml:"projection"`
	// Represents the provisioned throughput settings for the specified global secondary index.
	//
	// You must use either `OnDemandThroughput` or `ProvisionedThroughput` based on your table's capacity mode.
	//
	// For current minimum and maximum provisioned throughput values, see [Service, Account, and Table Quotas](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html) in the *Amazon DynamoDB Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-globalsecondaryindex.html#cfn-dynamodb-table-globalsecondaryindex-provisionedthroughput
	//
	ProvisionedThroughput interface{} `field:"optional" json:"provisionedThroughput" yaml:"provisionedThroughput"`
	// Represents the warm throughput value (in read units per second and write units per second) for the specified secondary index.
	//
	// If you use this parameter, you must specify `ReadUnitsPerSecond` , `WriteUnitsPerSecond` , or both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-table-globalsecondaryindex.html#cfn-dynamodb-table-globalsecondaryindex-warmthroughput
	//
	WarmThroughput interface{} `field:"optional" json:"warmThroughput" yaml:"warmThroughput"`
}

