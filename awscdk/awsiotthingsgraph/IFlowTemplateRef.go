package awsiotthingsgraph

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotthingsgraph/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FlowTemplate.
// Experimental.
type IFlowTemplateRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a FlowTemplate resource.
	// Experimental.
	FlowTemplateRef() *FlowTemplateReference
}

// The jsii proxy for IFlowTemplateRef
type jsiiProxy_IFlowTemplateRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IFlowTemplateRef) FlowTemplateRef() *FlowTemplateReference {
	var returns *FlowTemplateReference
	_jsii_.Get(
		j,
		"flowTemplateRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlowTemplateRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFlowTemplateRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

