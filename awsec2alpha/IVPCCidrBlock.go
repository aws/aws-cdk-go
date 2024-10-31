package awsec2alpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface to create L2 for VPC Cidr Block.
// Experimental.
type IVPCCidrBlock interface {
	// Amazon Provided Ipv6.
	// Experimental.
	AmazonProvidedIpv6CidrBlock() *bool
	// The secondary IPv4 CIDR Block.
	// Default: - no CIDR block provided.
	//
	// Experimental.
	CidrBlock() *string
	// IPAM pool for IPv4 address type.
	// Experimental.
	Ipv4IpamPoolId() *string
	// IPAM pool for IPv6 address type.
	// Experimental.
	Ipv6IpamPoolId() *string
}

// The jsii proxy for IVPCCidrBlock
type jsiiProxy_IVPCCidrBlock struct {
	_ byte // padding
}

func (j *jsiiProxy_IVPCCidrBlock) AmazonProvidedIpv6CidrBlock() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"amazonProvidedIpv6CidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVPCCidrBlock) CidrBlock() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cidrBlock",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVPCCidrBlock) Ipv4IpamPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv4IpamPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVPCCidrBlock) Ipv6IpamPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipv6IpamPoolId",
		&returns,
	)
	return returns
}

