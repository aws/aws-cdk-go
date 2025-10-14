package awssqs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a QueueInlinePolicy.
// Experimental.
type IQueueInlinePolicyRef interface {
	constructs.IConstruct
	// A reference to a QueueInlinePolicy resource.
	// Experimental.
	QueueInlinePolicyRef() *QueueInlinePolicyReference
}

// The jsii proxy for IQueueInlinePolicyRef
type jsiiProxy_IQueueInlinePolicyRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IQueueInlinePolicyRef) QueueInlinePolicyRef() *QueueInlinePolicyReference {
	var returns *QueueInlinePolicyReference
	_jsii_.Get(
		j,
		"queueInlinePolicyRef",
		&returns,
	)
	return returns
}

