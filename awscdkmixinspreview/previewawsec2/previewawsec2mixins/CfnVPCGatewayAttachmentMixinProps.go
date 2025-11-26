package previewawsec2mixins


// Properties for CfnVPCGatewayAttachmentPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVPCGatewayAttachmentMixinProps := &CfnVPCGatewayAttachmentMixinProps{
//   	InternetGatewayId: jsii.String("internetGatewayId"),
//   	VpcId: jsii.String("vpcId"),
//   	VpnGatewayId: jsii.String("vpnGatewayId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcgatewayattachment.html
//
type CfnVPCGatewayAttachmentMixinProps struct {
	// The ID of the internet gateway.
	//
	// You must specify either `InternetGatewayId` or `VpnGatewayId` , but not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcgatewayattachment.html#cfn-ec2-vpcgatewayattachment-internetgatewayid
	//
	InternetGatewayId *string `field:"optional" json:"internetGatewayId" yaml:"internetGatewayId"`
	// The ID of the VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcgatewayattachment.html#cfn-ec2-vpcgatewayattachment-vpcid
	//
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
	// The ID of the virtual private gateway.
	//
	// You must specify either `InternetGatewayId` or `VpnGatewayId` , but not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcgatewayattachment.html#cfn-ec2-vpcgatewayattachment-vpngatewayid
	//
	VpnGatewayId *string `field:"optional" json:"vpnGatewayId" yaml:"vpnGatewayId"`
}

