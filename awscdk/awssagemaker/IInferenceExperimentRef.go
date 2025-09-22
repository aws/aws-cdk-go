package awssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InferenceExperiment.
// Experimental.
type IInferenceExperimentRef interface {
	constructs.IConstruct
	// A reference to a InferenceExperiment resource.
	// Experimental.
	InferenceExperimentRef() *InferenceExperimentReference
}

// The jsii proxy for IInferenceExperimentRef
type jsiiProxy_IInferenceExperimentRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IInferenceExperimentRef) InferenceExperimentRef() *InferenceExperimentReference {
	var returns *InferenceExperimentReference
	_jsii_.Get(
		j,
		"inferenceExperimentRef",
		&returns,
	)
	return returns
}

