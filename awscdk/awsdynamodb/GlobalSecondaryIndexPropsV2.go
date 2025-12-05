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
//   // Add a GSI with multi-attribute keys
//   table.AddGlobalSecondaryIndex(&GlobalSecondaryIndexPropsV2{
//   	IndexName: jsii.String("multi-attribute-gsi2"),
//   	PartitionKeys: []Attribute{
//   		&Attribute{
//   			Name: jsii.String("multi-attribute_pk1"),
//   			Type: dynamodb.AttributeType_STRING,
//   		},
//   		&Attribute{
//   			Name: jsii.String("multi-attribute_pk2"),
//   			Type: dynamodb.AttributeType_NUMBER,
//   		},
//   	},
//   	SortKey: &Attribute{
//   		Name: jsii.String("sk"),
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
	// Partition key attribute definition.
	//
	// If a single field forms the partition key, you can use this field.  Use the
	// `partitionKeys` field if the partition key is a multi-attribute key (consists of
	// multiple fields).
	// Default: - exactly one of `partitionKey` and `partitionKeys` must be specified.
	//
	PartitionKey *Attribute `field:"optional" json:"partitionKey" yaml:"partitionKey"`
	// Multi-attribute partition key.
	//
	// If a single field forms the partition key, you can use either
	// `partitionKey` or `partitionKeys` to specify the partition key. Exactly
	// one of these must be specified.
	//
	// You must use `partitionKeys` field if the partition key is a multi-attribute key
	// (consists of multiple fields).
	//
	// NOTE: although the name of this field makes it sound like it creates
	// multiple keys, it does not. It defines a single key that consists of
	// of multiple fields.
	//
	// The order of fields is not important.
	// Default: - exactly one of `partitionKey` and `partitionKeys` must be specified.
	//
	PartitionKeys *[]*Attribute `field:"optional" json:"partitionKeys" yaml:"partitionKeys"`
	// The read capacity.
	//
	// Note: This can only be configured if the primary table billing is provisioned.
	// Default: - inherited from the primary table.
	//
	ReadCapacity Capacity `field:"optional" json:"readCapacity" yaml:"readCapacity"`
	// Sort key attribute definition.
	//
	// If a single field forms the sort key, you can use this field.  Use the
	// `sortKeys` field if the sort key is a multi-attribute key (consists of multiple
	// fields).
	// Default: - no sort key.
	//
	SortKey *Attribute `field:"optional" json:"sortKey" yaml:"sortKey"`
	// Multi-attribute sort key.
	//
	// If a single field forms the sort key, you can use either
	// `sortKey` or `sortKeys` to specify the sort key. At most one of these
	// may be specified.
	//
	// You must use `sortKeys` field if the sort key is a multi-attribute key
	// (consists of multiple fields).
	//
	// NOTE: although the name of this field makes it sound like it creates
	// multiple keys, it does not. It defines a single key that consists of
	// of multiple fields at the same time.
	//
	// NOTE: The order of fields is important!
	// Default: - no sort key.
	//
	SortKeys *[]*Attribute `field:"optional" json:"sortKeys" yaml:"sortKeys"`
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

