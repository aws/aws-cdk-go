package awsmediaconnect


// The VPC interface that you want to use for the media stream associated with the output.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   interfaceProperty := &InterfaceProperty{
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-interface.html
//
type CfnFlow_InterfaceProperty struct {
	// The name of the VPC interface that you want to use for the media stream associated with the output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-interface.html#cfn-mediaconnect-flow-interface-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
}

