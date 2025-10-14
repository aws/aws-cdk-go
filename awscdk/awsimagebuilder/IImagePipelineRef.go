package awsimagebuilder

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsimagebuilder/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ImagePipeline.
// Experimental.
type IImagePipelineRef interface {
	constructs.IConstruct
	// A reference to a ImagePipeline resource.
	// Experimental.
	ImagePipelineRef() *ImagePipelineReference
}

// The jsii proxy for IImagePipelineRef
type jsiiProxy_IImagePipelineRef struct {
	internal.Type__constructsIConstruct
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

