package awspcs


// A reference to a ComputeNodeGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   computeNodeGroupReference := &ComputeNodeGroupReference{
//   	ComputeNodeGroupArn: jsii.String("computeNodeGroupArn"),
//   }
//
type ComputeNodeGroupReference struct {
	// The Arn of the ComputeNodeGroup resource.
	ComputeNodeGroupArn *string `field:"required" json:"computeNodeGroupArn" yaml:"computeNodeGroupArn"`
}

