package awskafkaconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskafkaconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CustomPlugin.
// Experimental.
type ICustomPluginRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a CustomPlugin resource.
	// Experimental.
	CustomPluginRef() *CustomPluginReference
}

// The jsii proxy for ICustomPluginRef
type jsiiProxy_ICustomPluginRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ICustomPluginRef) CustomPluginRef() *CustomPluginReference {
	var returns *CustomPluginReference
	_jsii_.Get(
		j,
		"customPluginRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICustomPluginRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICustomPluginRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

