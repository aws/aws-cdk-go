package awsappstream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappstream/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ImageBuilder.
// Experimental.
type IImageBuilderRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ImageBuilder resource.
	// Experimental.
	ImageBuilderRef() *ImageBuilderReference
}

// The jsii proxy for IImageBuilderRef
type jsiiProxy_IImageBuilderRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IImageBuilderRef) ImageBuilderRef() *ImageBuilderReference {
	var returns *ImageBuilderReference
	_jsii_.Get(
		j,
		"imageBuilderRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IImageBuilderRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IImageBuilderRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

