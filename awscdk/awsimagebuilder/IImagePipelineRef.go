package awsimagebuilder

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsimagebuilder/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ImagePipeline.
// Experimental.
type IImagePipelineRef interface {
	constructs.IConstruct
}

// The jsii proxy for IImagePipelineRef
type jsiiProxy_IImagePipelineRef struct {
	internal.Type__constructsIConstruct
}

