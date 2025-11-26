package previewawsmediaconnectmixins


// The settings for attaching a VPC interface to an resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   vpcInterfaceAttachmentProperty := &VpcInterfaceAttachmentProperty{
//   	VpcInterfaceName: jsii.String("vpcInterfaceName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-bridge-vpcinterfaceattachment.html
//
type CfnBridgePropsMixin_VpcInterfaceAttachmentProperty struct {
	// The name of the VPC interface to use for this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-bridge-vpcinterfaceattachment.html#cfn-mediaconnect-bridge-vpcinterfaceattachment-vpcinterfacename
	//
	VpcInterfaceName *string `field:"optional" json:"vpcInterfaceName" yaml:"vpcInterfaceName"`
}

