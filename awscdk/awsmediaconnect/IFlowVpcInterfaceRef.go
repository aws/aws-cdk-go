package awsmediaconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmediaconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FlowVpcInterface.
// Experimental.
type IFlowVpcInterfaceRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a FlowVpcInterface resource.
	// Experimental.
	FlowVpcInterfaceRef() *FlowVpcInterfaceReference
}

// The jsii proxy for IFlowVpcInterfaceRef
type jsiiProxy_IFlowVpcInterfaceRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IFlowVpcInterfaceRef) FlowVpcInterfaceRef() *FlowVpcInterfaceReference {
	var returns *FlowVpcInterfaceReference
	_jsii_.Get(
		j,
		"flowVpcInterfaceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlowVpcInterfaceRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlowVpcInterfaceRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

