package interfacesawscodeguruprofiler

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscodeguruprofiler/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ProfilingGroup.
// Experimental.
type IProfilingGroupRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ProfilingGroup resource.
	// Experimental.
	ProfilingGroupRef() *ProfilingGroupReference
}

// The jsii proxy for IProfilingGroupRef
type jsiiProxy_IProfilingGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
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

func (j *jsiiProxy_IProfilingGroupRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
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

