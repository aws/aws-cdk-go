package awsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ContactFlowVersion.
// Experimental.
type IContactFlowVersionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ContactFlowVersion resource.
	// Experimental.
	ContactFlowVersionRef() *ContactFlowVersionReference
}

// The jsii proxy for IContactFlowVersionRef
type jsiiProxy_IContactFlowVersionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IContactFlowVersionRef) ContactFlowVersionRef() *ContactFlowVersionReference {
	var returns *ContactFlowVersionReference
	_jsii_.Get(
		j,
		"contactFlowVersionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IContactFlowVersionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IContactFlowVersionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

