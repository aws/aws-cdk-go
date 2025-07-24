package awsec2alpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Implements ip address allocation according to the IPAdress type.
// Experimental.
type IIpAddresses interface {
	// Method to define the implementation logic of IP address allocation.
	// Experimental.
	AllocateVpcCidr() *VpcCidrOptions
}

// The jsii proxy for IIpAddresses
type jsiiProxy_IIpAddresses struct {
	_ byte // padding
}

func (i *jsiiProxy_IIpAddresses) AllocateVpcCidr() *VpcCidrOptions {
	var returns *VpcCidrOptions

	_jsii_.Invoke(
		i,
		"allocateVpcCidr",
		nil, // no parameters
		&returns,
	)

	return returns
}

