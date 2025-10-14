package awssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InferenceComponent.
// Experimental.
type IInferenceComponentRef interface {
	constructs.IConstruct
	// A reference to a InferenceComponent resource.
	// Experimental.
	InferenceComponentRef() *InferenceComponentReference
}

// The jsii proxy for IInferenceComponentRef
type jsiiProxy_IInferenceComponentRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IInferenceComponentRef) InferenceComponentRef() *InferenceComponentReference {
	var returns *InferenceComponentReference
	_jsii_.Get(
		j,
		"inferenceComponentRef",
		&returns,
	)
	return returns
}

