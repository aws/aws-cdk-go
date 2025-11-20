package awsdynamodb


// Represents the table schema attributes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   schemaOptions := &SchemaOptions{
//   	PartitionKey: &Attribute{
//   		Name: jsii.String("name"),
//   		Type: awscdk.Aws_dynamodb.AttributeType_BINARY,
//   	},
//   	SortKey: &Attribute{
//   		Name: jsii.String("name"),
//   		Type: awscdk.*Aws_dynamodb.AttributeType_BINARY,
//   	},
//   }
//
type SchemaOptions struct {
	// Partition key attribute definition.
	//
	// If a single field forms the partition key, you can use this field.  Use the
	// `partitionKeys` field if the partition key is a compound key (consists of
	// multiple fields).
	// Default: - exactly one of `partitionKey` and `partitionKeys` must be specified.
	//
	PartitionKey *Attribute `field:"optional" json:"partitionKey" yaml:"partitionKey"`
	// Sort key attribute definition.
	//
	// If a single field forms the sort key, you can use this field.  Use the
	// `sortKeys` field if the sort key is a compound key (consists of multiple
	// fields).
	// Default: - no sort key.
	//
	SortKey *Attribute `field:"optional" json:"sortKey" yaml:"sortKey"`
}

