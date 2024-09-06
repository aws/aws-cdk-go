package awsec2alpha


// Properties to define a route table.
//
// Example:
//   stack := awscdk.Newstack()
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
//   igw := vpc_v2.NewInternetGateway(this, jsii.String("IGW"), &InternetGatewayProps{
//   	Vpc: myVpc,
//   })
//   vpc_v2.NewRoute(this, jsii.String("IgwRoute"), &RouteProps{
//   	RouteTable: RouteTable,
//   	Destination: jsii.String("0.0.0.0/0"),
//   	Target: map[string]iRouteTarget{
//   		"gateway": igw,
//   	},
//   })
//
// Experimental.
type RouteTableProps struct {
	// The ID of the VPC.
	// Experimental.
	Vpc IVpcV2 `field:"required" json:"vpc" yaml:"vpc"`
	// The resource name of the route table.
	// Default: none.
	//
	// Experimental.
	RouteTableName *string `field:"optional" json:"routeTableName" yaml:"routeTableName"`
}

