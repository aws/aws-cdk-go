package interfacesawsmediaconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmediaconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FlowSource.
// Experimental.
type IFlowSourceRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a FlowSource resource.
	// Experimental.
	FlowSourceRef() *FlowSourceReference
}

// The jsii proxy for IFlowSourceRef
type jsiiProxy_IFlowSourceRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IFlowSourceRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IFlowSourceRef) FlowSourceRef() *FlowSourceReference {
	var returns *FlowSourceReference
	_jsii_.Get(
		j,
		"flowSourceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlowSourceRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlowSourceRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

