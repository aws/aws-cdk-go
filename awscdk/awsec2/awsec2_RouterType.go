package awsec2


// Type of router used in route.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   vpc := ec2.NewVpc(this, jsii.String("VPC"), &vpcProps{
//   	subnetConfiguration: []subnetConfiguration{
//   		&subnetConfiguration{
//   			subnetType: ec2.subnetType_PUBLIC,
//   			name: jsii.String("Public"),
//   		},
//   		&subnetConfiguration{
//   			subnetType: ec2.*subnetType_PRIVATE_ISOLATED,
//   			name: jsii.String("Isolated"),
//   		},
//   	},
//   })
//
//   (vpc.isolatedSubnets[0].(subnet)).addRoute(jsii.String("StaticRoute"), &addRouteOptions{
//   	routerId: vpc.internetGatewayId,
//   	routerType: ec2.routerType_GATEWAY,
//   	destinationCidrBlock: jsii.String("8.8.8.8/32"),
//   })
//
type RouterType string

const (
	// Carrier gateway.
	RouterType_CARRIER_GATEWAY RouterType = "CARRIER_GATEWAY"
	// Egress-only Internet Gateway.
	RouterType_EGRESS_ONLY_INTERNET_GATEWAY RouterType = "EGRESS_ONLY_INTERNET_GATEWAY"
	// Internet Gateway.
	RouterType_GATEWAY RouterType = "GATEWAY"
	// Instance.
	RouterType_INSTANCE RouterType = "INSTANCE"
	// Local Gateway.
	RouterType_LOCAL_GATEWAY RouterType = "LOCAL_GATEWAY"
	// NAT Gateway.
	RouterType_NAT_GATEWAY RouterType = "NAT_GATEWAY"
	// Network Interface.
	RouterType_NETWORK_INTERFACE RouterType = "NETWORK_INTERFACE"
	// Transit Gateway.
	RouterType_TRANSIT_GATEWAY RouterType = "TRANSIT_GATEWAY"
	// VPC peering connection.
	RouterType_VPC_PEERING_CONNECTION RouterType = "VPC_PEERING_CONNECTION"
	// VPC Endpoint for gateway load balancers.
	RouterType_VPC_ENDPOINT RouterType = "VPC_ENDPOINT"
)

