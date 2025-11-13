package interfacesawsappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsappflow/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Flow.
// Experimental.
type IFlowRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Flow resource.
	// Experimental.
	FlowRef() *FlowReference
}

// The jsii proxy for IFlowRef
type jsiiProxy_IFlowRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IFlowRef) FlowRef() *FlowReference {
	var returns *FlowReference
	_jsii_.Get(
		j,
		"flowRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlowRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlowRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

