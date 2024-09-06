package awsec2alpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsec2alpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// IPv4 or IPv6 CIDR range for the subnet.
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

