package awsqbusiness

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsqbusiness/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Plugin.
// Experimental.
type IPluginRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Plugin resource.
	// Experimental.
	PluginRef() *PluginReference
}

// The jsii proxy for IPluginRef
type jsiiProxy_IPluginRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IPluginRef) PluginRef() *PluginReference {
	var returns *PluginReference
	_jsii_.Get(
		j,
		"pluginRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPluginRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPluginRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

