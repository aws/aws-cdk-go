package awsec2


// Properties for defining a `CfnVPCGatewayAttachment`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVPCGatewayAttachmentProps := &cfnVPCGatewayAttachmentProps{
//   	vpcId: jsii.String("vpcId"),
//
//   	// the properties below are optional
//   	internetGatewayId: jsii.String("internetGatewayId"),
//   	vpnGatewayId: jsii.String("vpnGatewayId"),
//   }
//
type CfnVPCGatewayAttachmentProps struct {
	// The ID of the VPC.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// The ID of the internet gateway.
	//
	// You must specify either `InternetGatewayId` or `VpnGatewayId` , but not both.
	InternetGatewayId *string `field:"optional" json:"internetGatewayId" yaml:"internetGatewayId"`
	// The ID of the virtual private gateway.
	//
	// You must specify either `InternetGatewayId` or `VpnGatewayId` , but not both.
	VpnGatewayId *string `field:"optional" json:"vpnGatewayId" yaml:"vpnGatewayId"`
}

