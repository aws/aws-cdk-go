package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FlowLog.
// Experimental.
type IFlowLogRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a FlowLog resource.
	// Experimental.
	FlowLogRef() *FlowLogReference
}

// The jsii proxy for IFlowLogRef
type jsiiProxy_IFlowLogRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IFlowLogRef) FlowLogRef() *FlowLogReference {
	var returns *FlowLogReference
	_jsii_.Get(
		j,
		"flowLogRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlowLogRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlowLogRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

