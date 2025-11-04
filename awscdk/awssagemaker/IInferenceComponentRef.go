package awssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InferenceComponent.
// Experimental.
type IInferenceComponentRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a InferenceComponent resource.
	// Experimental.
	InferenceComponentRef() *InferenceComponentReference
}

// The jsii proxy for IInferenceComponentRef
type jsiiProxy_IInferenceComponentRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IInferenceComponentRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInferenceComponentRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

