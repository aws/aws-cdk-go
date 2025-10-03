package awssqs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a QueueInlinePolicy.
// Experimental.
type IQueueInlinePolicyRef interface {
	constructs.IConstruct
}

// The jsii proxy for IQueueInlinePolicyRef
type jsiiProxy_IQueueInlinePolicyRef struct {
	internal.Type__constructsIConstruct
}

