package awsqbusiness

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsqbusiness/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Plugin.
// Experimental.
type IPluginRef interface {
	constructs.IConstruct
	// A reference to a Plugin resource.
	// Experimental.
	PluginRef() *PluginReference
}

// The jsii proxy for IPluginRef
type jsiiProxy_IPluginRef struct {
	internal.Type__constructsIConstruct
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

