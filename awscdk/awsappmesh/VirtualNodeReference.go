package awsappmesh


// A reference to a VirtualNode resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualNodeReference := &VirtualNodeReference{
//   	VirtualNodeArn: jsii.String("virtualNodeArn"),
//   	VirtualNodeId: jsii.String("virtualNodeId"),
//   }
//
type VirtualNodeReference struct {
	// The ARN of the VirtualNode resource.
	VirtualNodeArn *string `field:"required" json:"virtualNodeArn" yaml:"virtualNodeArn"`
	// The Id of the VirtualNode resource.
	VirtualNodeId *string `field:"required" json:"virtualNodeId" yaml:"virtualNodeId"`
}

