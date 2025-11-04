package awsimagebuilder

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsimagebuilder/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ImagePipeline.
// Experimental.
type IImagePipelineRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ImagePipeline resource.
	// Experimental.
	ImagePipelineRef() *ImagePipelineReference
}

// The jsii proxy for IImagePipelineRef
type jsiiProxy_IImagePipelineRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IImagePipelineRef) ImagePipelineRef() *ImagePipelineReference {
	var returns *ImagePipelineReference
	_jsii_.Get(
		j,
		"imagePipelineRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IImagePipelineRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IImagePipelineRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

