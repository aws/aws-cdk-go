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
//
//   	// the properties below are optional
//   	SortKey: &Attribute{
//   		Name: jsii.String("name"),
//   		Type: awscdk.*Aws_dynamodb.AttributeType_BINARY,
//   	},
//   }
//
type SchemaOptions struct {
	// Partition key attribute definition.
	PartitionKey *Attribute `field:"required" json:"partitionKey" yaml:"partitionKey"`
	// Sort key attribute definition.
	// Default: no sort key.
	//
	SortKey *Attribute `field:"optional" json:"sortKey" yaml:"sortKey"`
}

