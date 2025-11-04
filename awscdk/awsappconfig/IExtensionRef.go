package awsappconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappconfig/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Extension.
// Experimental.
type IExtensionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Extension resource.
	// Experimental.
	ExtensionRef() *ExtensionReference
}

// The jsii proxy for IExtensionRef
type jsiiProxy_IExtensionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IExtensionRef) ExtensionRef() *ExtensionReference {
	var returns *ExtensionReference
	_jsii_.Get(
		j,
		"extensionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExtensionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IExtensionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

