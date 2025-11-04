package awsiot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResourceSpecificLogging.
// Experimental.
type IResourceSpecificLoggingRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ResourceSpecificLogging resource.
	// Experimental.
	ResourceSpecificLoggingRef() *ResourceSpecificLoggingReference
}

// The jsii proxy for IResourceSpecificLoggingRef
type jsiiProxy_IResourceSpecificLoggingRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IResourceSpecificLoggingRef) ResourceSpecificLoggingRef() *ResourceSpecificLoggingReference {
	var returns *ResourceSpecificLoggingReference
	_jsii_.Get(
		j,
		"resourceSpecificLoggingRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResourceSpecificLoggingRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResourceSpecificLoggingRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

