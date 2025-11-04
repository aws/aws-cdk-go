package awssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InferenceExperiment.
// Experimental.
type IInferenceExperimentRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a InferenceExperiment resource.
	// Experimental.
	InferenceExperimentRef() *InferenceExperimentReference
}

// The jsii proxy for IInferenceExperimentRef
type jsiiProxy_IInferenceExperimentRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IInferenceExperimentRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInferenceExperimentRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

