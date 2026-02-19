package previewawsdynamodbmixins


// Allows you to specify a global secondary index for the global table.
//
// The index will be defined on all replicas.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   globalSecondaryIndexProperty := &GlobalSecondaryIndexProperty{
//   	IndexName: jsii.String("indexName"),
//   	KeySchema: []interface{}{
//   		&KeySchemaProperty{
//   			AttributeName: jsii.String("attributeName"),
//   			KeyType: jsii.String("keyType"),
//   		},
//   	},
//   	Projection: &ProjectionProperty{
//   		NonKeyAttributes: []*string{
//   			jsii.String("nonKeyAttributes"),
//   		},
//   		ProjectionType: jsii.String("projectionType"),
//   	},
//   	ReadOnDemandThroughputSettings: &ReadOnDemandThroughputSettingsProperty{
//   		MaxReadRequestUnits: jsii.Number(123),
//   	},
//   	ReadProvisionedThroughputSettings: &GlobalReadProvisionedThroughputSettingsProperty{
//   		ReadCapacityUnits: jsii.Number(123),
//   	},
//   	WarmThroughput: &WarmThroughputProperty{
//   		ReadUnitsPerSecond: jsii.Number(123),
//   		WriteUnitsPerSecond: jsii.Number(123),
//   	},
//   	WriteOnDemandThroughputSettings: &WriteOnDemandThroughputSettingsProperty{
//   		MaxWriteRequestUnits: jsii.Number(123),
//   	},
//   	WriteProvisionedThroughputSettings: &WriteProvisionedThroughputSettingsProperty{
//   		WriteCapacityAutoScalingSettings: &CapacityAutoScalingSettingsProperty{
//   			MaxCapacity: jsii.Number(123),
//   			MinCapacity: jsii.Number(123),
//   			SeedCapacity: jsii.Number(123),
//   			TargetTrackingScalingPolicyConfiguration: &TargetTrackingScalingPolicyConfigurationProperty{
//   				DisableScaleIn: jsii.Boolean(false),
//   				ScaleInCooldown: jsii.Number(123),
//   				ScaleOutCooldown: jsii.Number(123),
//   				TargetValue: jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-globalsecondaryindex.html
//
type CfnGlobalTablePropsMixin_GlobalSecondaryIndexProperty struct {
	// The name of the global secondary index.
	//
	// The name must be unique among all other indexes on this table.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-globalsecondaryindex.html#cfn-dynamodb-globaltable-globalsecondaryindex-indexname
	//
	IndexName *string `field:"optional" json:"indexName" yaml:"indexName"`
	// The complete key schema for a global secondary index, which consists of one or more pairs of attribute names and key types:  - `HASH` - partition key - `RANGE` - sort key  > The partition key of an item is also known as its *hash attribute* .
	//
	// The term "hash attribute" derives from DynamoDB's usage of an internal hash function to evenly distribute data items across partitions, based on their partition key values.
	// >
	// > The sort key of an item is also known as its *range attribute* . The term "range attribute" derives from the way DynamoDB stores items with the same partition key physically close together, in sorted order by the sort key value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-globalsecondaryindex.html#cfn-dynamodb-globaltable-globalsecondaryindex-keyschema
	//
	KeySchema interface{} `field:"optional" json:"keySchema" yaml:"keySchema"`
	// Represents attributes that are copied (projected) from the table into the global secondary index.
	//
	// These are in addition to the primary key attributes and index key attributes, which are automatically projected.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-globalsecondaryindex.html#cfn-dynamodb-globaltable-globalsecondaryindex-projection
	//
	Projection interface{} `field:"optional" json:"projection" yaml:"projection"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-globalsecondaryindex.html#cfn-dynamodb-globaltable-globalsecondaryindex-readondemandthroughputsettings
	//
	ReadOnDemandThroughputSettings interface{} `field:"optional" json:"readOnDemandThroughputSettings" yaml:"readOnDemandThroughputSettings"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-globalsecondaryindex.html#cfn-dynamodb-globaltable-globalsecondaryindex-readprovisionedthroughputsettings
	//
	ReadProvisionedThroughputSettings interface{} `field:"optional" json:"readProvisionedThroughputSettings" yaml:"readProvisionedThroughputSettings"`
	// Represents the warm throughput value (in read units per second and write units per second) for the specified secondary index.
	//
	// If you use this parameter, you must specify `ReadUnitsPerSecond` , `WriteUnitsPerSecond` , or both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-globalsecondaryindex.html#cfn-dynamodb-globaltable-globalsecondaryindex-warmthroughput
	//
	WarmThroughput interface{} `field:"optional" json:"warmThroughput" yaml:"warmThroughput"`
	// Sets the write request settings for a global table or a global secondary index.
	//
	// You can only specify this setting if your resource uses the `PAY_PER_REQUEST` `BillingMode` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-globalsecondaryindex.html#cfn-dynamodb-globaltable-globalsecondaryindex-writeondemandthroughputsettings
	//
	WriteOnDemandThroughputSettings interface{} `field:"optional" json:"writeOnDemandThroughputSettings" yaml:"writeOnDemandThroughputSettings"`
	// Defines write capacity settings for the global secondary index.
	//
	// You must specify a value for this property if the table's `BillingMode` is `PROVISIONED` . All replicas will have the same write capacity settings for this global secondary index.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dynamodb-globaltable-globalsecondaryindex.html#cfn-dynamodb-globaltable-globalsecondaryindex-writeprovisionedthroughputsettings
	//
	WriteProvisionedThroughputSettings interface{} `field:"optional" json:"writeProvisionedThroughputSettings" yaml:"writeProvisionedThroughputSettings"`
}

