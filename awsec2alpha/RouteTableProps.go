package awsec2alpha


// Properties to define a route table.
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
type RouteTableProps struct {
	// The ID of the VPC.
	// Experimental.
	Vpc IVpcV2 `field:"required" json:"vpc" yaml:"vpc"`
	// The resource name of the route table.
	// Default: - provisioned without a route table name.
	//
	// Experimental.
	RouteTableName *string `field:"optional" json:"routeTableName" yaml:"routeTableName"`
}

