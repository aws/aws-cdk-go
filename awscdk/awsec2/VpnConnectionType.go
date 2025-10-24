package awsec2


// The VPN connection type.
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
type VpnConnectionType string

const (
	// The IPsec 1 VPN connection type.
	VpnConnectionType_IPSEC_1 VpnConnectionType = "IPSEC_1"
	// Dummy member TODO: remove once https://github.com/aws/jsii/issues/231 is fixed.
	VpnConnectionType_DUMMY VpnConnectionType = "DUMMY"
)

