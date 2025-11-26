package previewawsmediaconnectmixins


// The VPC interface that is used for the media stream associated with the source or output.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   interfaceProperty := &InterfaceProperty{
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-interface.html
//
type CfnFlowPropsMixin_InterfaceProperty struct {
	// The name of the VPC interface.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flow-interface.html#cfn-mediaconnect-flow-interface-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}

