package interfacesawsmediaconnect


// A reference to a FlowVpcInterface resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   flowVpcInterfaceReference := &FlowVpcInterfaceReference{
//   	FlowArn: jsii.String("flowArn"),
//   	FlowVpcInterfaceName: jsii.String("flowVpcInterfaceName"),
//   }
//
type FlowVpcInterfaceReference struct {
	// The FlowArn of the FlowVpcInterface resource.
	FlowArn *string `field:"required" json:"flowArn" yaml:"flowArn"`
	// The Name of the FlowVpcInterface resource.
	FlowVpcInterfaceName *string `field:"required" json:"flowVpcInterfaceName" yaml:"flowVpcInterfaceName"`
}

