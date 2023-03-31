package awsdynamodb


// Properties for a global secondary index.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   globalSecondaryIndexProps := &globalSecondaryIndexProps{
//   	indexName: jsii.String("indexName"),
//   	partitionKey: &attribute{
//   		name: jsii.String("name"),
//   		type: awscdk.Aws_dynamodb.attributeType_BINARY,
//   	},
//
//   	// the properties below are optional
//   	nonKeyAttributes: []*string{
//   		jsii.String("nonKeyAttributes"),
//   	},
//   	projectionType: awscdk.*Aws_dynamodb.projectionType_KEYS_ONLY,
//   	readCapacity: jsii.Number(123),
//   	sortKey: &attribute{
//   		name: jsii.String("name"),
//   		type: awscdk.*Aws_dynamodb.*attributeType_BINARY,
//   	},
//   	writeCapacity: jsii.Number(123),
//   }
//
// Experimental.
type GlobalSecondaryIndexProps struct {
	// The name of the secondary index.
	// Experimental.
	IndexName *string `field:"required" json:"indexName" yaml:"indexName"`
	// The non-key attributes that are projected into the secondary index.
	// Experimental.
	NonKeyAttributes *[]*string `field:"optional" json:"nonKeyAttributes" yaml:"nonKeyAttributes"`
	// The set of attributes that are projected into the secondary index.
	// Experimental.
	ProjectionType ProjectionType `field:"optional" json:"projectionType" yaml:"projectionType"`
	// Partition key attribute definition.
	// Experimental.
	PartitionKey *Attribute `field:"required" json:"partitionKey" yaml:"partitionKey"`
	// Sort key attribute definition.
	// Experimental.
	SortKey *Attribute `field:"optional" json:"sortKey" yaml:"sortKey"`
	// The read capacity for the global secondary index.
	//
	// Can only be provided if table billingMode is Provisioned or undefined.
	// Experimental.
	ReadCapacity *float64 `field:"optional" json:"readCapacity" yaml:"readCapacity"`
	// The write capacity for the global secondary index.
	//
	// Can only be provided if table billingMode is Provisioned or undefined.
	// Experimental.
	WriteCapacity *float64 `field:"optional" json:"writeCapacity" yaml:"writeCapacity"`
}

