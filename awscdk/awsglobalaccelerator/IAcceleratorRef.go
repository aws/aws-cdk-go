package awsglobalaccelerator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsglobalaccelerator/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Accelerator.
// Experimental.
type IAcceleratorRef interface {
	constructs.IConstruct
	// A reference to a Accelerator resource.
	// Experimental.
	AcceleratorRef() *AcceleratorReference
}

// The jsii proxy for IAcceleratorRef
type jsiiProxy_IAcceleratorRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAcceleratorRef) AcceleratorRef() *AcceleratorReference {
	var returns *AcceleratorReference
	_jsii_.Get(
		j,
		"acceleratorRef",
		&returns,
	)
	return returns
}

