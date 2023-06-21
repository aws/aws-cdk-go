package awsec2


// Type of router used in route.
//
// Example:
//   vpc := ec2.NewVpc(this, jsii.String("VPC"), &VpcProps{
//   	SubnetConfiguration: []subnetConfiguration{
//   		&subnetConfiguration{
//   			SubnetType: ec2.SubnetType_PUBLIC,
//   			Name: jsii.String("Public"),
//   		},
//   		&subnetConfiguration{
//   			SubnetType: ec2.SubnetType_PRIVATE_ISOLATED,
//   			Name: jsii.String("Isolated"),
//   		},
//   	},
//   })
//
//   (vpc.IsolatedSubnets[0].(subnet)).AddRoute(jsii.String("StaticRoute"), &AddRouteOptions{
//   	RouterId: vpc.InternetGatewayId,
//   	RouterType: ec2.RouterType_GATEWAY,
//   	DestinationCidrBlock: jsii.String("8.8.8.8/32"),
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

