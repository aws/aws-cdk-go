package awsdynamodb


// Properties used to configure a global secondary index.
//
// Example:
//   table := dynamodb.NewTableV2(this, jsii.String("Table"), &TablePropsV2{
//   	PartitionKey: &Attribute{
//   		Name: jsii.String("pk"),
//   		Type: dynamodb.AttributeType_STRING,
//   	},
//   	GlobalSecondaryIndexes: []GlobalSecondaryIndexPropsV2{
//   		&GlobalSecondaryIndexPropsV2{
//   			IndexName: jsii.String("gsi1"),
//   			PartitionKey: &Attribute{
//   				Name: jsii.String("pk"),
//   				Type: dynamodb.AttributeType_STRING,
//   			},
//   		},
//   	},
//   })
//
//   table.AddGlobalSecondaryIndex(&GlobalSecondaryIndexPropsV2{
//   	IndexName: jsii.String("gsi2"),
//   	PartitionKey: &Attribute{
//   		Name: jsii.String("pk"),
//   		Type: dynamodb.AttributeType_STRING,
//   	},
//   })
//
type GlobalSecondaryIndexPropsV2 struct {
	// The name of the secondary index.
	IndexName *string `field:"required" json:"indexName" yaml:"indexName"`
	// The non-key attributes that are projected into the secondary index.
	// Default: - No additional attributes.
	//
	NonKeyAttributes *[]*string `field:"optional" json:"nonKeyAttributes" yaml:"nonKeyAttributes"`
	// The set of attributes that are projected into the secondary index.
	// Default: ALL.
	//
	ProjectionType ProjectionType `field:"optional" json:"projectionType" yaml:"projectionType"`
	// Partition key attribute definition.
	PartitionKey *Attribute `field:"required" json:"partitionKey" yaml:"partitionKey"`
	// The maximum read request units.
	//
	// Note: This can only be configured if the primary table billing is PAY_PER_REQUEST.
	// Default: - inherited from the primary table.
	//
	MaxReadRequestUnits *float64 `field:"optional" json:"maxReadRequestUnits" yaml:"maxReadRequestUnits"`
	// The maximum write request units.
	//
	// Note: This can only be configured if the primary table billing is PAY_PER_REQUEST.
	// Default: - inherited from the primary table.
	//
	MaxWriteRequestUnits *float64 `field:"optional" json:"maxWriteRequestUnits" yaml:"maxWriteRequestUnits"`
	// The read capacity.
	//
	// Note: This can only be configured if the primary table billing is provisioned.
	// Default: - inherited from the primary table.
	//
	ReadCapacity Capacity `field:"optional" json:"readCapacity" yaml:"readCapacity"`
	// Sort key attribute definition.
	// Default: - no sort key.
	//
	SortKey *Attribute `field:"optional" json:"sortKey" yaml:"sortKey"`
	// The warm throughput configuration for the global secondary index.
	// Default: - no warm throughput is configured.
	//
	WarmThroughput *WarmThroughput `field:"optional" json:"warmThroughput" yaml:"warmThroughput"`
	// The write capacity.
	//
	// Note: This can only be configured if the primary table billing is provisioned.
	// Default: - inherited from the primary table.
	//
	WriteCapacity Capacity `field:"optional" json:"writeCapacity" yaml:"writeCapacity"`
}

