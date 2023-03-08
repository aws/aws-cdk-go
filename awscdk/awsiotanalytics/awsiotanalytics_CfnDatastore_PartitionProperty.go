package awsiotanalytics


// A single dimension to partition a data store.
//
// The dimension must be an `AttributePartition` or a `TimestampPartition` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   partitionProperty := &partitionProperty{
//   	attributeName: jsii.String("attributeName"),
//   }
//
type CfnDatastore_PartitionProperty struct {
	// The name of the attribute that defines a partition dimension.
	AttributeName *string `field:"required" json:"attributeName" yaml:"attributeName"`
}

