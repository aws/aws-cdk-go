package awswisdom

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awswisdom/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a QuickResponse.
// Experimental.
type IQuickResponseRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a QuickResponse resource.
	// Experimental.
	QuickResponseRef() *QuickResponseReference
}

// The jsii proxy for IQuickResponseRef
type jsiiProxy_IQuickResponseRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IQuickResponseRef) QuickResponseRef() *QuickResponseReference {
	var returns *QuickResponseReference
	_jsii_.Get(
		j,
		"quickResponseRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IQuickResponseRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IQuickResponseRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

