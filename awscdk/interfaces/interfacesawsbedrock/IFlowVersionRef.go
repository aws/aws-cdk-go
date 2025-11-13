package interfacesawsbedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbedrock/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FlowVersion.
// Experimental.
type IFlowVersionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a FlowVersion resource.
	// Experimental.
	FlowVersionRef() *FlowVersionReference
}

// The jsii proxy for IFlowVersionRef
type jsiiProxy_IFlowVersionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IFlowVersionRef) FlowVersionRef() *FlowVersionReference {
	var returns *FlowVersionReference
	_jsii_.Get(
		j,
		"flowVersionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlowVersionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlowVersionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

