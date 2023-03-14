package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Implementations for ip address management.
type IIpAddresses interface {
	// Called by the VPC to retrieve Subnet options from the Ipam.
	//
	// Don't call this directly, the VPC will call it automatically.
	AllocateSubnetsCidr(input *AllocateCidrRequest) *SubnetIpamOptions
	// Called by the VPC to retrieve VPC options from the Ipam.
	//
	// Don't call this directly, the VPC will call it automatically.
	AllocateVpcCidr() *VpcIpamOptions
}

// The jsii proxy for IIpAddresses
type jsiiProxy_IIpAddresses struct {
	_ byte // padding
}

func (i *jsiiProxy_IIpAddresses) AllocateSubnetsCidr(input *AllocateCidrRequest) *SubnetIpamOptions {
	if err := i.validateAllocateSubnetsCidrParameters(input); err != nil {
		panic(err)
	}
	var returns *SubnetIpamOptions

	_jsii_.Invoke(
		i,
		"allocateSubnetsCidr",
		[]interface{}{input},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IIpAddresses) AllocateVpcCidr() *VpcIpamOptions {
	var returns *VpcIpamOptions

	_jsii_.Invoke(
		i,
		"allocateVpcCidr",
		nil, // no parameters
		&returns,
	)

	return returns
}

