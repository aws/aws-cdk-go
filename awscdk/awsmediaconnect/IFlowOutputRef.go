package awsmediaconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediaconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FlowOutput.
// Experimental.
type IFlowOutputRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a FlowOutput resource.
	// Experimental.
	FlowOutputRef() *FlowOutputReference
}

// The jsii proxy for IFlowOutputRef
type jsiiProxy_IFlowOutputRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IFlowOutputRef) FlowOutputRef() *FlowOutputReference {
	var returns *FlowOutputReference
	_jsii_.Get(
		j,
		"flowOutputRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlowOutputRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlowOutputRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

