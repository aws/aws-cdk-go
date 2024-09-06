package awsec2alpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsec2alpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// IpAddress options to define VPC V2.
//
// Example:
//   stack := awscdk.Newstack()
//   myVpc := vpc_v2.NewVpcV2(this, jsii.String("Vpc"), &VpcV2Props{
//   	SecondaryAddressBlocks: []iIpAddresses{
//   		vpc_v2.IpAddresses_AmazonProvidedIpv6(&SecondaryAddressProps{
//   			CidrBlockName: jsii.String("AmazonProvidedIp"),
//   		}),
//   	},
//   })
//
//   vpc_v2.NewSubnetV2(this, jsii.String("subnetA"), &SubnetV2Props{
//   	Vpc: myVpc,
//   	AvailabilityZone: jsii.String("us-east-1a"),
//   	Ipv4CidrBlock: vpc_v2.NewIpCidr(jsii.String("10.0.0.0/24")),
//   	Ipv6CidrBlock: vpc_v2.NewIpCidr(jsii.String("2a05:d02c:25:4000::/60")),
//   	SubnetType: ec2.SubnetType_PRIVATE_ISOLATED,
//   })
//
// Experimental.
type IpAddresses interface {
}

// The jsii proxy struct for IpAddresses
type jsiiProxy_IpAddresses struct {
	_ byte // padding
}

// Experimental.
func NewIpAddresses() IpAddresses {
	_init_.Initialize()

	j := jsiiProxy_IpAddresses{}

	_jsii_.Create(
		"@aws-cdk/aws-ec2-alpha.IpAddresses",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewIpAddresses_Override(i IpAddresses) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-ec2-alpha.IpAddresses",
		nil, // no parameters
		i,
	)
}

// Amazon Provided Ipv6 range.
// Experimental.
func IpAddresses_AmazonProvidedIpv6(props *SecondaryAddressProps) IIpAddresses {
	_init_.Initialize()

	if err := validateIpAddresses_AmazonProvidedIpv6Parameters(props); err != nil {
		panic(err)
	}
	var returns IIpAddresses

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-ec2-alpha.IpAddresses",
		"amazonProvidedIpv6",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// An IPv4 CIDR Range.
// Experimental.
func IpAddresses_Ipv4(ipv4Cidr *string, props *SecondaryAddressProps) IIpAddresses {
	_init_.Initialize()

	if err := validateIpAddresses_Ipv4Parameters(ipv4Cidr, props); err != nil {
		panic(err)
	}
	var returns IIpAddresses

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-ec2-alpha.IpAddresses",
		"ipv4",
		[]interface{}{ipv4Cidr, props},
		&returns,
	)

	return returns
}

// An Ipv4 Ipam Pool.
// Experimental.
func IpAddresses_Ipv4Ipam(ipv4IpamOptions *IpamOptions) IIpAddresses {
	_init_.Initialize()

	if err := validateIpAddresses_Ipv4IpamParameters(ipv4IpamOptions); err != nil {
		panic(err)
	}
	var returns IIpAddresses

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-ec2-alpha.IpAddresses",
		"ipv4Ipam",
		[]interface{}{ipv4IpamOptions},
		&returns,
	)

	return returns
}

// An Ipv6 Ipam Pool.
// Experimental.
func IpAddresses_Ipv6Ipam(ipv6IpamOptions *IpamOptions) IIpAddresses {
	_init_.Initialize()

	if err := validateIpAddresses_Ipv6IpamParameters(ipv6IpamOptions); err != nil {
		panic(err)
	}
	var returns IIpAddresses

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-ec2-alpha.IpAddresses",
		"ipv6Ipam",
		[]interface{}{ipv6IpamOptions},
		&returns,
	)

	return returns
}

