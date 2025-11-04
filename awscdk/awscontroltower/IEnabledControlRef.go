package awscontroltower

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscontroltower/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EnabledControl.
// Experimental.
type IEnabledControlRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a EnabledControl resource.
	// Experimental.
	EnabledControlRef() *EnabledControlReference
}

// The jsii proxy for IEnabledControlRef
type jsiiProxy_IEnabledControlRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IEnabledControlRef) EnabledControlRef() *EnabledControlReference {
	var returns *EnabledControlReference
	_jsii_.Get(
		j,
		"enabledControlRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnabledControlRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnabledControlRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

