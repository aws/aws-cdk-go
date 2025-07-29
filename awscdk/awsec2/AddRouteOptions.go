package awsec2


// Options for adding a new route to a subnet.
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
type AddRouteOptions struct {
	// The ID of the router.
	//
	// Can be an instance ID, gateway ID, etc, depending on the router type.
	RouterId *string `field:"required" json:"routerId" yaml:"routerId"`
	// What type of router to route this traffic to.
	RouterType RouterType `field:"required" json:"routerType" yaml:"routerType"`
	// IPv4 range this route applies to.
	// Default: '0.0.0.0/0'
	//
	DestinationCidrBlock *string `field:"optional" json:"destinationCidrBlock" yaml:"destinationCidrBlock"`
	// IPv6 range this route applies to.
	// Default: - Uses IPv6.
	//
	DestinationIpv6CidrBlock *string `field:"optional" json:"destinationIpv6CidrBlock" yaml:"destinationIpv6CidrBlock"`
	// Whether this route will enable internet connectivity.
	//
	// If true, this route will be added before any AWS resources that depend
	// on internet connectivity in the VPC will be created.
	// Default: false.
	//
	EnablesInternetConnectivity *bool `field:"optional" json:"enablesInternetConnectivity" yaml:"enablesInternetConnectivity"`
}

