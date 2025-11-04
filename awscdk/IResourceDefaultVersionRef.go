package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResourceDefaultVersion.
// Experimental.
type IResourceDefaultVersionRef interface {
	constructs.IConstruct
	IEnvironmentAware
	// A reference to a ResourceDefaultVersion resource.
	// Experimental.
	ResourceDefaultVersionRef() *ResourceDefaultVersionReference
}

// The jsii proxy for IResourceDefaultVersionRef
type jsiiProxy_IResourceDefaultVersionRef struct {
	internal.Type__constructsIConstruct
	jsiiProxy_IEnvironmentAware
}

func (j *jsiiProxy_IResourceDefaultVersionRef) ResourceDefaultVersionRef() *ResourceDefaultVersionReference {
	var returns *ResourceDefaultVersionReference
	_jsii_.Get(
		j,
		"resourceDefaultVersionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResourceDefaultVersionRef) Env() *ResourceEnvironment {
	var returns *ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResourceDefaultVersionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

