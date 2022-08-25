package awsec2


// Options for adding a new route to a subnet.
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
type AddRouteOptions struct {
	// The ID of the router.
	//
	// Can be an instance ID, gateway ID, etc, depending on the router type.
	RouterId *string `field:"required" json:"routerId" yaml:"routerId"`
	// What type of router to route this traffic to.
	RouterType RouterType `field:"required" json:"routerType" yaml:"routerType"`
	// IPv4 range this route applies to.
	DestinationCidrBlock *string `field:"optional" json:"destinationCidrBlock" yaml:"destinationCidrBlock"`
	// IPv6 range this route applies to.
	DestinationIpv6CidrBlock *string `field:"optional" json:"destinationIpv6CidrBlock" yaml:"destinationIpv6CidrBlock"`
	// Whether this route will enable internet connectivity.
	//
	// If true, this route will be added before any AWS resources that depend
	// on internet connectivity in the VPC will be created.
	EnablesInternetConnectivity *bool `field:"optional" json:"enablesInternetConnectivity" yaml:"enablesInternetConnectivity"`
}

