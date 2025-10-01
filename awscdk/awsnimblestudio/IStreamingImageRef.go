package awsnimblestudio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsnimblestudio/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StreamingImage.
// Experimental.
type IStreamingImageRef interface {
	constructs.IConstruct
	// A reference to a StreamingImage resource.
	// Experimental.
	StreamingImageRef() *StreamingImageReference
}

// The jsii proxy for IStreamingImageRef
type jsiiProxy_IStreamingImageRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IStreamingImageRef) StreamingImageRef() *StreamingImageReference {
	var returns *StreamingImageReference
	_jsii_.Get(
		j,
		"streamingImageRef",
		&returns,
	)
	return returns
}

