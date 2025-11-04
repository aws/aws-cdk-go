package awspinpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awspinpoint/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApplicationSettings.
// Experimental.
type IApplicationSettingsRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ApplicationSettings resource.
	// Experimental.
	ApplicationSettingsRef() *ApplicationSettingsReference
}

// The jsii proxy for IApplicationSettingsRef
type jsiiProxy_IApplicationSettingsRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IApplicationSettingsRef) ApplicationSettingsRef() *ApplicationSettingsReference {
	var returns *ApplicationSettingsReference
	_jsii_.Get(
		j,
		"applicationSettingsRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationSettingsRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApplicationSettingsRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

