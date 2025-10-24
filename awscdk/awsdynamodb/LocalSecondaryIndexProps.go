package awsdynamodb


// Properties for a local secondary index.
//
// Example:
//   table := dynamodb.NewTableV2(this, jsii.String("Table"), &TablePropsV2{
//   	PartitionKey: &Attribute{
//   		Name: jsii.String("pk"),
//   		Type: dynamodb.AttributeType_STRING,
//   	},
//   	SortKey: &Attribute{
//   		Name: jsii.String("sk"),
//   		Type: dynamodb.AttributeType_NUMBER,
//   	},
//   	LocalSecondaryIndexes: []LocalSecondaryIndexProps{
//   		&LocalSecondaryIndexProps{
//   			IndexName: jsii.String("lsi1"),
//   			SortKey: &Attribute{
//   				Name: jsii.String("sk"),
//   				Type: dynamodb.AttributeType_NUMBER,
//   			},
//   		},
//   	},
//   })
//
//   table.AddLocalSecondaryIndex(&LocalSecondaryIndexProps{
//   	IndexName: jsii.String("lsi2"),
//   	SortKey: &Attribute{
//   		Name: jsii.String("sk"),
//   		Type: dynamodb.AttributeType_NUMBER,
//   	},
//   })
//
type LocalSecondaryIndexProps struct {
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
	// The attribute of a sort key for the local secondary index.
	SortKey *Attribute `field:"required" json:"sortKey" yaml:"sortKey"`
}

