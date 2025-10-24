package awsec2alpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Options to define VPNGatewayV2 for VPC.
//
// Example:
//   stack := awscdk.Newstack()
//   myVpc := awsec2alpha.NewVpcV2(this, jsii.String("Vpc"))
//   vpnGateway := myVpc.EnableVpnGatewayV2(&VPNGatewayV2Options{
//   	VpnRoutePropagation: []SubnetSelection{
//   		&SubnetSelection{
//   			SubnetType: awscdk.SubnetType_PUBLIC,
//   		},
//   	},
//   	Type: awscdk.VpnConnectionType_IPSEC_1,
//   })
//
//   routeTable := awsec2alpha.NewRouteTable(stack, jsii.String("routeTable"), &RouteTableProps{
//   	Vpc: myVpc,
//   })
//
//   awsec2alpha.NewRoute(stack, jsii.String("route"), &RouteProps{
//   	Destination: jsii.String("172.31.0.0/24"),
//   	Target: map[string]iRouteTarget{
//   		"gateway": vpnGateway,
//   	},
//   	RouteTable: routeTable,
//   })
//
// Experimental.
type VPNGatewayV2Options struct {
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
}

