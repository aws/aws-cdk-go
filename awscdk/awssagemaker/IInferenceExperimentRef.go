package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InferenceExperiment.
// Experimental.
type IInferenceExperimentRef interface {
	constructs.IConstruct
}

// The jsii proxy for IInferenceExperimentRef
type jsiiProxy_IInferenceExperimentRef struct {
	internal.Type__constructsIConstruct
}

