package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Either an IPv4 or an IPv6 CIDR.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aclCidr := awscdk.Aws_ec2.AclCidr_AnyIpv4()
//
type AclCidr interface {
	ToCidrConfig() *AclCidrConfig
}

// The jsii proxy struct for AclCidr
type jsiiProxy_AclCidr struct {
	_ byte // padding
}

func NewAclCidr_Override(a AclCidr) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.AclCidr",
		nil, // no parameters
		a,
	)
}

// The CIDR containing all IPv4 addresses (i.e., 0.0.0.0/0).
func AclCidr_AnyIpv4() AclCidr {
	_init_.Initialize()

	var returns AclCidr

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.AclCidr",
		"anyIpv4",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The CIDR containing all IPv6 addresses (i.e., ::/0).
func AclCidr_AnyIpv6() AclCidr {
	_init_.Initialize()

	var returns AclCidr

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.AclCidr",
		"anyIpv6",
		nil, // no parameters
		&returns,
	)

	return returns
}

// An IP network range in CIDR notation (for example, 172.16.0.0/24).
func AclCidr_Ipv4(ipv4Cidr *string) AclCidr {
	_init_.Initialize()

	if err := validateAclCidr_Ipv4Parameters(ipv4Cidr); err != nil {
		panic(err)
	}
	var returns AclCidr

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.AclCidr",
		"ipv4",
		[]interface{}{ipv4Cidr},
		&returns,
	)

	return returns
}

// An IPv6 network range in CIDR notation (for example, 2001:db8::/48).
func AclCidr_Ipv6(ipv6Cidr *string) AclCidr {
	_init_.Initialize()

	if err := validateAclCidr_Ipv6Parameters(ipv6Cidr); err != nil {
		panic(err)
	}
	var returns AclCidr

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.AclCidr",
		"ipv6",
		[]interface{}{ipv6Cidr},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AclCidr) ToCidrConfig() *AclCidrConfig {
	var returns *AclCidrConfig

	_jsii_.Invoke(
		a,
		"toCidrConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

