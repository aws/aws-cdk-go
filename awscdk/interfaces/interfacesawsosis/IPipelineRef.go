package interfacesawsosis

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsosis/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Pipeline.
// Experimental.
type IPipelineRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Pipeline resource.
	// Experimental.
	PipelineRef() *PipelineReference
}

// The jsii proxy for IPipelineRef
type jsiiProxy_IPipelineRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IPipelineRef) PipelineRef() *PipelineReference {
	var returns *PipelineReference
	_jsii_.Get(
		j,
		"pipelineRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPipelineRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPipelineRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

