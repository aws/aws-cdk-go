package awsmediaconnect


// The VPC interface that you want to send your output to.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcInterfaceAttachmentProperty := &VpcInterfaceAttachmentProperty{
//   	VpcInterfaceName: jsii.String("vpcInterfaceName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowsource-vpcinterfaceattachment.html
//
type CfnFlowSource_VpcInterfaceAttachmentProperty struct {
	// The name of the VPC interface that you want to send your output to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-flowsource-vpcinterfaceattachment.html#cfn-mediaconnect-flowsource-vpcinterfaceattachment-vpcinterfacename
	//
	VpcInterfaceName *string `field:"optional" json:"vpcInterfaceName" yaml:"vpcInterfaceName"`
}

