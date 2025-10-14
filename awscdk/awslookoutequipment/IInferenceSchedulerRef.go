package awslookoutequipment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslookoutequipment/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InferenceScheduler.
// Experimental.
type IInferenceSchedulerRef interface {
	constructs.IConstruct
	// A reference to a InferenceScheduler resource.
	// Experimental.
	InferenceSchedulerRef() *InferenceSchedulerReference
}

// The jsii proxy for IInferenceSchedulerRef
type jsiiProxy_IInferenceSchedulerRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IInferenceSchedulerRef) InferenceSchedulerRef() *InferenceSchedulerReference {
	var returns *InferenceSchedulerReference
	_jsii_.Get(
		j,
		"inferenceSchedulerRef",
		&returns,
	)
	return returns
}

