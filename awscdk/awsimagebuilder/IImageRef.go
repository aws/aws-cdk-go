package awsimagebuilder

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsimagebuilder/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Image.
// Experimental.
type IImageRef interface {
	constructs.IConstruct
	// A reference to a Image resource.
	// Experimental.
	ImageRef() *ImageReference
}

// The jsii proxy for IImageRef
type jsiiProxy_IImageRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IImageRef) ImageRef() *ImageReference {
	var returns *ImageReference
	_jsii_.Get(
		j,
		"imageRef",
		&returns,
	)
	return returns
}

