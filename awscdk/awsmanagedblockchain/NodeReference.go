package awsmanagedblockchain


// A reference to a Node resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nodeReference := &NodeReference{
//   	NodeArn: jsii.String("nodeArn"),
//   	NodeId: jsii.String("nodeId"),
//   }
//
type NodeReference struct {
	// The ARN of the Node resource.
	NodeArn *string `field:"required" json:"nodeArn" yaml:"nodeArn"`
	// The NodeId of the Node resource.
	NodeId *string `field:"required" json:"nodeId" yaml:"nodeId"`
}

