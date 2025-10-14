package awskafkaconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskafkaconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CustomPlugin.
// Experimental.
type ICustomPluginRef interface {
	constructs.IConstruct
	// A reference to a CustomPlugin resource.
	// Experimental.
	CustomPluginRef() *CustomPluginReference
}

// The jsii proxy for ICustomPluginRef
type jsiiProxy_ICustomPluginRef struct {
	internal.Type__constructsIConstruct
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

