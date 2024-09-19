package awsec2alpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsec2alpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// IPv4 or IPv6 CIDR range for the subnet.
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
//   	SubnetType: awscdk.SubnetType_PUBLIC,
//   })
//
//   myVpc.AddInternetGateway()
//   myVpc.AddNatGateway(&NatGatewayOptions{
//   	Subnet: subnet,
//   	ConnectivityType: awsec2alpha.NatConnectivityType_PUBLIC,
//   })
//
// Experimental.
type IpCidr interface {
	// IPv6 CIDR range for the subnet Allowed only if IPv6 is enabled on VPc.
	// Experimental.
	Cidr() *string
}

// The jsii proxy struct for IpCidr
type jsiiProxy_IpCidr struct {
	_ byte // padding
}

func (j *jsiiProxy_IpCidr) Cidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cidr",
		&returns,
	)
	return returns
}


// Experimental.
func NewIpCidr(props *string) IpCidr {
	_init_.Initialize()

	if err := validateNewIpCidrParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_IpCidr{}

	_jsii_.Create(
		"@aws-cdk/aws-ec2-alpha.IpCidr",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewIpCidr_Override(i IpCidr, props *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-ec2-alpha.IpCidr",
		[]interface{}{props},
		i,
	)
}

