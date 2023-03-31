package awsglobalaccelerator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsglobalaccelerator/internal"
)

// The interface of the Accelerator.
// Experimental.
type IAccelerator interface {
	awscdk.IResource
	// The ARN of the accelerator.
	// Experimental.
	AcceleratorArn() *string
	// The Domain Name System (DNS) name that Global Accelerator creates that points to your accelerator's static IP addresses.
	// Experimental.
	DnsName() *string
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

