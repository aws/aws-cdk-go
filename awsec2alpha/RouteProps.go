package awsec2alpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Properties to define a route.
//
// Example:
//   myVpc := awsec2alpha.NewVpcV2(this, jsii.String("Vpc"))
//   routeTable := awsec2alpha.NewRouteTable(this, jsii.String("RouteTable"), &RouteTableProps{
//   	Vpc: myVpc,
//   })
//   subnet := awsec2alpha.NewSubnetV2(this, jsii.String("Subnet"), &SubnetV2Props{
//   	Vpc: myVpc,
//   	AvailabilityZone: jsii.String("eu-west-2a"),
//   	Ipv4CidrBlock: awsec2alpha.NewIpCidr(jsii.String("10.0.0.0/24")),
//   	SubnetType: awscdk.SubnetType_PRIVATE_ISOLATED,
//   })
//
//   natgw := awsec2alpha.NewNatGateway(this, jsii.String("NatGW"), &NatGatewayProps{
//   	Subnet: subnet,
//   	Vpc: myVpc,
//   	ConnectivityType: awsec2alpha.NatConnectivityType_PRIVATE,
//   	PrivateIpAddress: jsii.String("10.0.0.42"),
//   })
//   awsec2alpha.NewRoute(this, jsii.String("NatGwRoute"), &RouteProps{
//   	RouteTable: RouteTable,
//   	Destination: jsii.String("0.0.0.0/0"),
//   	Target: map[string]iRouteTarget{
//   		"gateway": natgw,
//   	},
//   })
//
// Experimental.
type RouteProps struct {
	// The IPv4 or IPv6 CIDR block used for the destination match.
	//
	// Routing decisions are based on the most specific match.
	// Experimental.
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// The ID of the route table for the route.
	// Experimental.
	RouteTable awsec2.IRouteTable `field:"required" json:"routeTable" yaml:"routeTable"`
	// The gateway or endpoint targeted by the route.
	// Experimental.
	Target RouteTargetType `field:"required" json:"target" yaml:"target"`
	// The resource name of the route.
	// Default: - provisioned without a route name.
	//
	// Experimental.
	RouteName *string `field:"optional" json:"routeName" yaml:"routeName"`
}

