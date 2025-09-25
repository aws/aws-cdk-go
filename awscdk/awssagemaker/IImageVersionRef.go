package awssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ImageVersion.
// Experimental.
type IImageVersionRef interface {
	constructs.IConstruct
	// A reference to a ImageVersion resource.
	// Experimental.
	ImageVersionRef() *ImageVersionReference
}

// The jsii proxy for IImageVersionRef
type jsiiProxy_IImageVersionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IImageVersionRef) ImageVersionRef() *ImageVersionReference {
	var returns *ImageVersionReference
	_jsii_.Get(
		j,
		"imageVersionRef",
		&returns,
	)
	return returns
}

