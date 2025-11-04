package awscloudformation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CustomResource.
// Experimental.
type ICustomResourceRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a CustomResource resource.
	// Experimental.
	CustomResourceRef() *CustomResourceReference
}

// The jsii proxy for ICustomResourceRef
type jsiiProxy_ICustomResourceRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ICustomResourceRef) CustomResourceRef() *CustomResourceReference {
	var returns *CustomResourceReference
	_jsii_.Get(
		j,
		"customResourceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICustomResourceRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICustomResourceRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

