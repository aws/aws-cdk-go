package awsec2alpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Properties to define a VPN gateway.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import ec2_alpha "github.com/aws/aws-cdk-go/awsec2alpha"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var vpcV2 vpcV2
//
//   vPNGatewayProps := &VPNGatewayProps{
//   	Type: awscdk.Aws_ec2.VpnConnectionType_IPSEC_1,
//   	Vpc: vpcV2,
//
//   	// the properties below are optional
//   	AmazonSideAsn: jsii.Number(123),
//   	VpnGatewayName: jsii.String("vpnGatewayName"),
//   }
//
// Experimental.
type VPNGatewayProps struct {
	// The type of VPN connection the virtual private gateway supports.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpngateway.html#cfn-ec2-vpngateway-type
	//
	// Experimental.
	Type awsec2.VpnConnectionType `field:"required" json:"type" yaml:"type"`
	// The ID of the VPC for which to create the VPN gateway.
	// Experimental.
	Vpc IVpcV2 `field:"required" json:"vpc" yaml:"vpc"`
	// The private Autonomous System Number (ASN) for the Amazon side of a BGP session.
	// Default: none.
	//
	// Experimental.
	AmazonSideAsn *float64 `field:"optional" json:"amazonSideAsn" yaml:"amazonSideAsn"`
	// The resource name of the VPN gateway.
	// Default: none.
	//
	// Experimental.
	VpnGatewayName *string `field:"optional" json:"vpnGatewayName" yaml:"vpnGatewayName"`
}

