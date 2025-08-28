package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// An abstract Provider of Ipv6Addresses.
//
// Note this is specific to the IPv6 CIDR.
type Ipv6Addresses interface {
}

// The jsii proxy struct for Ipv6Addresses
type jsiiProxy_Ipv6Addresses struct {
	_ byte // padding
}

// Used for IPv6 address management with Amazon provided CIDRs.
//
// Note this is specific to the IPv6 CIDR.
func Ipv6Addresses_AmazonProvided() IIpv6Addresses {
	_init_.Initialize()

	var returns IIpv6Addresses

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.Ipv6Addresses",
		"amazonProvided",
		nil, // no parameters
		&returns,
	)

	return returns
}

