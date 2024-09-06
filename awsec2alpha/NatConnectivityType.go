package awsec2alpha


// Indicates whether the NAT gateway supports public or private connectivity.
//
// The default is public connectivity.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-natgateway.html#cfn-ec2-natgateway-connectivitytype
//
// Example:
//   myVpc := vpc_v2.NewVpcV2(this, jsii.String("Vpc"))
//   routeTable := vpc_v2.NewRouteTable(this, jsii.String("RouteTable"), &RouteTableProps{
//   	Vpc: myVpc,
//   })
//   subnet := vpc_v2.NewSubnetV2(this, jsii.String("Subnet"), &SubnetV2Props{
//   	Vpc: myVpc,
//   	AvailabilityZone: jsii.String("eu-west-2a"),
//   	Ipv4CidrBlock: awsec2alpha.NewIpCidr(jsii.String("10.0.0.0/24")),
//   	SubnetType: ec2.SubnetType_PRIVATE_ISOLATED,
//   })
//
//   natgw := vpc_v2.NewNatGateway(this, jsii.String("NatGW"), &NatGatewayProps{
//   	Subnet: subnet,
//   	Vpc: myVpc,
//   	ConnectivityType: awsec2alpha.NatConnectivityType_PRIVATE,
//   	PrivateIpAddress: jsii.String("10.0.0.42"),
//   })
//   vpc_v2.NewRoute(this, jsii.String("NatGwRoute"), &RouteProps{
//   	RouteTable: RouteTable,
//   	Destination: jsii.String("0.0.0.0/0"),
//   	Target: map[string]iRouteTarget{
//   		"gateway": natgw,
//   	},
//   })
//
// Experimental.
type NatConnectivityType string

const (
	// Sets Connectivity type to PUBLIC.
	// Experimental.
	NatConnectivityType_PUBLIC NatConnectivityType = "PUBLIC"
	// Sets Connectivity type to PRIVATE.
	// Experimental.
	NatConnectivityType_PRIVATE NatConnectivityType = "PRIVATE"
)

