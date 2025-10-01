package awsappstream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappstream/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ImageBuilder.
// Experimental.
type IImageBuilderRef interface {
	constructs.IConstruct
	// A reference to a ImageBuilder resource.
	// Experimental.
	ImageBuilderRef() *ImageBuilderReference
}

// The jsii proxy for IImageBuilderRef
type jsiiProxy_IImageBuilderRef struct {
	internal.Type__constructsIConstruct
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

