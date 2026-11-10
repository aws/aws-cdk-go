package interfacesawsosis

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsosis/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PipelineEndpoint.
// Experimental.
type IPipelineEndpointRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a PipelineEndpoint resource.
	// Experimental.
	PipelineEndpointRef() *PipelineEndpointReference
}

// The jsii proxy for IPipelineEndpointRef
type jsiiProxy_IPipelineEndpointRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IPipelineEndpointRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IPipelineEndpointRef) PipelineEndpointRef() *PipelineEndpointReference {
	var returns *PipelineEndpointReference
	_jsii_.Get(
		j,
		"pipelineEndpointRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPipelineEndpointRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPipelineEndpointRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

