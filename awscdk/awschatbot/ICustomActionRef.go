package awschatbot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awschatbot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CustomAction.
// Experimental.
type ICustomActionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a CustomAction resource.
	// Experimental.
	CustomActionRef() *CustomActionReference
}

// The jsii proxy for ICustomActionRef
type jsiiProxy_ICustomActionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ICustomActionRef) CustomActionRef() *CustomActionReference {
	var returns *CustomActionReference
	_jsii_.Get(
		j,
		"customActionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICustomActionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICustomActionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

