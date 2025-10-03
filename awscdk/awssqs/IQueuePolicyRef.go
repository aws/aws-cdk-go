package awssqs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a QueuePolicy.
// Experimental.
type IQueuePolicyRef interface {
	constructs.IConstruct
}

// The jsii proxy for IQueuePolicyRef
type jsiiProxy_IQueuePolicyRef struct {
	internal.Type__constructsIConstruct
}

