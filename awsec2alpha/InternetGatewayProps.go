package awsec2alpha


// Properties to define an internet gateway.
//
// Example:
//   stack := awscdk.Newstack()
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
//   igw := awsec2alpha.NewInternetGateway(this, jsii.String("IGW"), &InternetGatewayProps{
//   	Vpc: myVpc,
//   })
//   awsec2alpha.NewRoute(this, jsii.String("IgwRoute"), &RouteProps{
//   	RouteTable: RouteTable,
//   	Destination: jsii.String("0.0.0.0/0"),
//   	Target: map[string]iRouteTarget{
//   		"gateway": igw,
//   	},
//   })
//
// Experimental.
type InternetGatewayProps struct {
	// The ID of the VPC for which to create the internet gateway.
	// Experimental.
	Vpc IVpcV2 `field:"required" json:"vpc" yaml:"vpc"`
	// The resource name of the internet gateway.
	// Default: - provisioned without a resource name.
	//
	// Experimental.
	InternetGatewayName *string `field:"optional" json:"internetGatewayName" yaml:"internetGatewayName"`
}

