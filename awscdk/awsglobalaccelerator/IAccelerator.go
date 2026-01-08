package awsglobalaccelerator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglobalaccelerator/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsglobalaccelerator"
	"github.com/aws/constructs-go/constructs/v10"
)

// The interface of the Accelerator.
type IAccelerator interface {
	interfacesawsglobalaccelerator.IAcceleratorRef
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
	internal.Type__interfacesawsglobalacceleratorIAcceleratorRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IAccelerator) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
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

func (j *jsiiProxy_IAccelerator) AcceleratorRef() *interfacesawsglobalaccelerator.AcceleratorReference {
	var returns *interfacesawsglobalaccelerator.AcceleratorReference
	_jsii_.Get(
		j,
		"acceleratorRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccelerator) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccelerator) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAccelerator) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

