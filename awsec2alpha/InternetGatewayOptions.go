package awsec2alpha


// Options to define InternetGateway for VPC.
//
// Example:
//   stack := awscdk.Newstack()
//   myVpc := awsec2alpha.NewVpcV2(this, jsii.String("Vpc"))
//
//   subnet := awsec2alpha.NewSubnetV2(this, jsii.String("Subnet"), &SubnetV2Props{
//   	Vpc: myVpc,
//   	AvailabilityZone: jsii.String("eu-west-2a"),
//   	Ipv4CidrBlock: awsec2alpha.NewIpCidr(jsii.String("10.0.0.0/24")),
//   	SubnetType: awscdk.SubnetType_PUBLIC,
//   })
//
//   myVpc.AddInternetGateway(&InternetGatewayOptions{
//   	Ipv4Destination: jsii.String("192.168.0.0/16"),
//   })
//
// Experimental.
type InternetGatewayOptions struct {
	// Destination Ipv6 address for EGW route.
	// Default: - '0.0.0.0' all Ipv4 traffic
	//
	// Experimental.
	Ipv4Destination *string `field:"optional" json:"ipv4Destination" yaml:"ipv4Destination"`
	// Destination Ipv6 address for EGW route.
	// Default: - '::/0' all Ipv6 traffic.
	//
	// Experimental.
	Ipv6Destination *string `field:"optional" json:"ipv6Destination" yaml:"ipv6Destination"`
}

