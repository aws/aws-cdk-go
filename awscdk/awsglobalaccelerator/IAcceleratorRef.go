package awsglobalaccelerator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglobalaccelerator/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Accelerator.
// Experimental.
type IAcceleratorRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Accelerator resource.
	// Experimental.
	AcceleratorRef() *AcceleratorReference
}

// The jsii proxy for IAcceleratorRef
type jsiiProxy_IAcceleratorRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IAcceleratorRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAcceleratorRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

