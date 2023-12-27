package awsglobalaccelerator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglobalaccelerator/internal"
)

// The interface of the Accelerator.
type IAccelerator interface {
	awscdk.IResource
	// The ARN of the accelerator.
	AcceleratorArn() *string
	// The Domain Name System (DNS) name that Global Accelerator creates that points to your accelerator's static IP addresses.
	DnsName() *string
	// The DNS name that Global Accelerator creates that points to a dual-stack accelerator's four static IP addresses: two IPv4 addresses and two IPv6 addresses.
	DualStackDnsName() *string
	// The array of IPv4 addresses in the IP address set.
	//
	// An IP address set can have a maximum of two IP addresses.
	Ipv4Addresses() *[]*string
	// The array of IPv6 addresses in the IP address set.
	//
	// An IP address set can have a maximum of two IP addresses.
	Ipv6Addresses() *[]*string
}

// The jsii proxy for IAccelerator
type jsiiProxy_IAccelerator struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IAccelerator) AcceleratorArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acceleratorArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccelerator) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccelerator) DualStackDnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dualStackDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccelerator) Ipv4Addresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv4Addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccelerator) Ipv6Addresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipv6Addresses",
		&returns,
	)
	return returns
}

