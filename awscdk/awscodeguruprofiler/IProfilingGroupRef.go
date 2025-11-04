package awscodeguruprofiler

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodeguruprofiler/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ProfilingGroup.
// Experimental.
type IProfilingGroupRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ProfilingGroup resource.
	// Experimental.
	ProfilingGroupRef() *ProfilingGroupReference
}

// The jsii proxy for IProfilingGroupRef
type jsiiProxy_IProfilingGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IProfilingGroupRef) ProfilingGroupRef() *ProfilingGroupReference {
	var returns *ProfilingGroupReference
	_jsii_.Get(
		j,
		"profilingGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProfilingGroupRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IProfilingGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

