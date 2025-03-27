package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Implementations for IPv6 address management.
//
// Note this is specific to the IPv6 CIDR.
type IIpv6Addresses interface {
	// Allocates Subnets IPv6 CIDRs. Called by VPC when creating subnets with IPv6 enabled.
	//
	// Note this is specific to the IPv6 CIDR.
	AllocateSubnetsIpv6Cidr(input *AllocateIpv6CidrRequest) *SubnetIpamOptions
	// Called by VPC to allocate IPv6 CIDR.
	//
	// Note this is specific to the IPv6 CIDR.
	AllocateVpcIpv6Cidr(input *AllocateVpcIpv6CidrRequest) CfnVPCCidrBlock
	// Split IPv6 CIDR block up for subnets.
	//
	// Note this is specific to the IPv6 CIDR.
	CreateIpv6CidrBlocks(input *CreateIpv6CidrBlocksRequest) *[]*string
	// Whether the IPv6 CIDR is Amazon provided or not.
	//
	// Note this is specific to the IPv6 CIDR.
	AmazonProvided() *bool
	SetAmazonProvided(a *bool)
}

// The jsii proxy for IIpv6Addresses
type jsiiProxy_IIpv6Addresses struct {
	_ byte // padding
}

func (i *jsiiProxy_IIpv6Addresses) AllocateSubnetsIpv6Cidr(input *AllocateIpv6CidrRequest) *SubnetIpamOptions {
	if err := i.validateAllocateSubnetsIpv6CidrParameters(input); err != nil {
		panic(err)
	}
	var returns *SubnetIpamOptions

	_jsii_.Invoke(
		i,
		"allocateSubnetsIpv6Cidr",
		[]interface{}{input},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IIpv6Addresses) AllocateVpcIpv6Cidr(input *AllocateVpcIpv6CidrRequest) CfnVPCCidrBlock {
	if err := i.validateAllocateVpcIpv6CidrParameters(input); err != nil {
		panic(err)
	}
	var returns CfnVPCCidrBlock

	_jsii_.Invoke(
		i,
		"allocateVpcIpv6Cidr",
		[]interface{}{input},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IIpv6Addresses) CreateIpv6CidrBlocks(input *CreateIpv6CidrBlocksRequest) *[]*string {
	if err := i.validateCreateIpv6CidrBlocksParameters(input); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"createIpv6CidrBlocks",
		[]interface{}{input},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IIpv6Addresses) AmazonProvided() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"amazonProvided",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIpv6Addresses)SetAmazonProvided(val *bool) {
	if err := j.validateSetAmazonProvidedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"amazonProvided",
		val,
	)
}

