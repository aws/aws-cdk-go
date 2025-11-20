package awsdynamodb


// A description of a key schema of an LSI, GSI or Table.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   keySchema := &KeySchema{
//   	PartitionKeys: []Attribute{
//   		&Attribute{
//   			Name: jsii.String("name"),
//   			Type: awscdk.Aws_dynamodb.AttributeType_BINARY,
//   		},
//   	},
//   	SortKeys: []Attribute{
//   		&Attribute{
//   			Name: jsii.String("name"),
//   			Type: awscdk.*Aws_dynamodb.AttributeType_BINARY,
//   		},
//   	},
//   }
//
type KeySchema struct {
	// Partition key definition.
	//
	// This array has at least one, but potentially multiple entries.  Together,
	// they form the partition key.
	PartitionKeys *[]*Attribute `field:"required" json:"partitionKeys" yaml:"partitionKeys"`
	// Sort key definition.
	//
	// This array has zero or more entries. Together, they form the sort key.
	SortKeys *[]*Attribute `field:"required" json:"sortKeys" yaml:"sortKeys"`
}

