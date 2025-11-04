package awsfms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsfms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResourceSet.
// Experimental.
type IResourceSetRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ResourceSet resource.
	// Experimental.
	ResourceSetRef() *ResourceSetReference
}

// The jsii proxy for IResourceSetRef
type jsiiProxy_IResourceSetRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IResourceSetRef) ResourceSetRef() *ResourceSetReference {
	var returns *ResourceSetReference
	_jsii_.Get(
		j,
		"resourceSetRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResourceSetRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResourceSetRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

