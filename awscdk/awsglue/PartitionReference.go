package awsglue


// A reference to a Partition resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   partitionReference := &PartitionReference{
//   	PartitionId: jsii.String("partitionId"),
//   }
//
type PartitionReference struct {
	// The Id of the Partition resource.
	PartitionId *string `field:"required" json:"partitionId" yaml:"partitionId"`
}

