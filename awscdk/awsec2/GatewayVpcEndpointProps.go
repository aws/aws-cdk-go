package awsec2


// Construction properties for a GatewayVpcEndpoint.
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
//   	SubnetType: awscdk.SubnetType_PRIVATE,
//   })
//
//   dynamoEndpoint := ec2.NewGatewayVpcEndpoint(this, jsii.String("DynamoEndpoint"), &GatewayVpcEndpointProps{
//   	Service: ec2.GatewayVpcEndpointAwsService_DYNAMODB(),
//   	Vpc: myVpc,
//   	Subnets: []subnetSelection{
//   		subnet,
//   	},
//   })
//   awsec2alpha.NewRoute(this, jsii.String("DynamoDBRoute"), &RouteProps{
//   	RouteTable: RouteTable,
//   	Destination: jsii.String("0.0.0.0/0"),
//   	Target: map[string]iVpcEndpoint{
//   		"endpoint": dynamoEndpoint,
//   	},
//   })
//
type GatewayVpcEndpointProps struct {
	// The service to use for this gateway VPC endpoint.
	Service IGatewayVpcEndpointService `field:"required" json:"service" yaml:"service"`
	// Where to add endpoint routing.
	//
	// By default, this endpoint will be routable from all subnets in the VPC.
	// Specify a list of subnet selection objects here to be more specific.
	//
	// Example:
	//   declare const vpc: ec2.Vpc;
	//
	//   vpc.addGatewayEndpoint('DynamoDbEndpoint', {
	//     service: ec2.GatewayVpcEndpointAwsService.DYNAMODB,
	//     // Add only to ISOLATED subnets
	//     subnets: [
	//       { subnetType: ec2.SubnetType.PRIVATE_ISOLATED }
	//     ]
	//   });
	//
	// Default: - All subnets in the VPC.
	//
	Subnets *[]*SubnetSelection `field:"optional" json:"subnets" yaml:"subnets"`
	// The VPC network in which the gateway endpoint will be used.
	Vpc IVpc `field:"required" json:"vpc" yaml:"vpc"`
}

