package awsdynamodb


// Properties for a global secondary index.
//
// Example:
//   table := dynamodb.NewTable(this, jsii.String("Table"), &TableProps{
//   	PartitionKey: &Attribute{
//   		Name: jsii.String("pk"),
//   		Type: dynamodb.AttributeType_STRING,
//   	},
//   	ContributorInsightsSpecification: &ContributorInsightsSpecification{
//   		 // for a table
//   		Enabled: jsii.Boolean(true),
//   		Mode: dynamodb.ContributorInsightsMode_THROTTLED_KEYS,
//   	},
//   })
//
//   table.AddGlobalSecondaryIndex(&GlobalSecondaryIndexProps{
//   	ContributorInsightsSpecification: &ContributorInsightsSpecification{
//   		 // for a specific global secondary index
//   		Enabled: jsii.Boolean(true),
//   	},
//   	IndexName: jsii.String("gsi"),
//   	PartitionKey: &Attribute{
//   		Name: jsii.String("pk"),
//   		Type: dynamodb.AttributeType_STRING,
//   	},
//   })
//
type GlobalSecondaryIndexProps struct {
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
	//
	// If a single field forms the partition key, you can use this field.  Use the
	// `partitionKeys` field if the partition key is a multi-attribute key (consists of
	// multiple fields).
	// Default: - exactly one of `partitionKey` and `partitionKeys` must be specified.
	//
	PartitionKey *Attribute `field:"optional" json:"partitionKey" yaml:"partitionKey"`
	// Sort key attribute definition.
	//
	// If a single field forms the sort key, you can use this field.  Use the
	// `sortKeys` field if the sort key is a multi-attribute key (consists of multiple
	// fields).
	// Default: - no sort key.
	//
	SortKey *Attribute `field:"optional" json:"sortKey" yaml:"sortKey"`
	// Whether CloudWatch contributor insights is enabled for the specified global secondary index.
	// Default: false.
	//
	// Deprecated: use `contributorInsightsSpecification` instead.
	ContributorInsightsEnabled *bool `field:"optional" json:"contributorInsightsEnabled" yaml:"contributorInsightsEnabled"`
	// Whether CloudWatch contributor insights is enabled and what mode is selected.
	// Default: - contributor insights is not enabled.
	//
	ContributorInsightsSpecification *ContributorInsightsSpecification `field:"optional" json:"contributorInsightsSpecification" yaml:"contributorInsightsSpecification"`
	// The maximum read request units for the global secondary index.
	//
	// Can only be provided if table billingMode is PAY_PER_REQUEST.
	// Default: - on-demand throughput is disabled.
	//
	MaxReadRequestUnits *float64 `field:"optional" json:"maxReadRequestUnits" yaml:"maxReadRequestUnits"`
	// The maximum write request units for the global secondary index.
	//
	// Can only be provided if table billingMode is PAY_PER_REQUEST.
	// Default: - on-demand throughput is disabled.
	//
	MaxWriteRequestUnits *float64 `field:"optional" json:"maxWriteRequestUnits" yaml:"maxWriteRequestUnits"`
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
	// The read capacity for the global secondary index.
	//
	// Can only be provided if table billingMode is Provisioned or undefined.
	// Default: 5.
	//
	ReadCapacity *float64 `field:"optional" json:"readCapacity" yaml:"readCapacity"`
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
	// The write capacity for the global secondary index.
	//
	// Can only be provided if table billingMode is Provisioned or undefined.
	// Default: 5.
	//
	WriteCapacity *float64 `field:"optional" json:"writeCapacity" yaml:"writeCapacity"`
}

