package awsram

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsram/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResourceShare.
// Experimental.
type IResourceShareRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ResourceShare resource.
	// Experimental.
	ResourceShareRef() *ResourceShareReference
}

// The jsii proxy for IResourceShareRef
type jsiiProxy_IResourceShareRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IResourceShareRef) ResourceShareRef() *ResourceShareReference {
	var returns *ResourceShareReference
	_jsii_.Get(
		j,
		"resourceShareRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResourceShareRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResourceShareRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

