package awslookoutequipment

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslookoutequipment/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InferenceScheduler.
// Experimental.
type IInferenceSchedulerRef interface {
	constructs.IConstruct
}

// The jsii proxy for IInferenceSchedulerRef
type jsiiProxy_IInferenceSchedulerRef struct {
	internal.Type__constructsIConstruct
}

