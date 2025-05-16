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
//   var subnet subnet
//   var subnetFilter subnetFilter
//   var vpcV2 vpcV2
//
//   vPNGatewayV2Props := &VPNGatewayV2Props{
//   	Type: awscdk.Aws_ec2.VpnConnectionType_IPSEC_1,
//   	Vpc: vpcV2,
//
//   	// the properties below are optional
//   	AmazonSideAsn: jsii.Number(123),
//   	VpnGatewayName: jsii.String("vpnGatewayName"),
//   	VpnRoutePropagation: []subnetSelection{
//   		&subnetSelection{
//   			AvailabilityZones: []*string{
//   				jsii.String("availabilityZones"),
//   			},
//   			OnePerAz: jsii.Boolean(false),
//   			SubnetFilters: []*subnetFilter{
//   				subnetFilter,
//   			},
//   			SubnetGroupName: jsii.String("subnetGroupName"),
//   			Subnets: []iSubnet{
//   				subnet,
//   			},
//   			SubnetType: awscdk.*Aws_ec2.SubnetType_PRIVATE_ISOLATED,
//   		},
//   	},
//   }
//
// Experimental.
type VPNGatewayV2Props struct {
	// The type of VPN connection the virtual private gateway supports.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpngateway.html#cfn-ec2-vpngateway-type
	//
	// Experimental.
	Type awsec2.VpnConnectionType `field:"required" json:"type" yaml:"type"`
	// The private Autonomous System Number (ASN) for the Amazon side of a BGP session.
	// Default: - no ASN set for BGP session.
	//
	// Experimental.
	AmazonSideAsn *float64 `field:"optional" json:"amazonSideAsn" yaml:"amazonSideAsn"`
	// The resource name of the VPN gateway.
	// Default: - resource provisioned without any name.
	//
	// Experimental.
	VpnGatewayName *string `field:"optional" json:"vpnGatewayName" yaml:"vpnGatewayName"`
	// Subnets where the route propagation should be added.
	// Default: - no propogation for routes.
	//
	// Experimental.
	VpnRoutePropagation *[]*awsec2.SubnetSelection `field:"optional" json:"vpnRoutePropagation" yaml:"vpnRoutePropagation"`
	// The ID of the VPC for which to create the VPN gateway.
	// Experimental.
	Vpc IVpcV2 `field:"required" json:"vpc" yaml:"vpc"`
}

