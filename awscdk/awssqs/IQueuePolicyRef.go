package awssqs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a QueuePolicy.
// Experimental.
type IQueuePolicyRef interface {
	constructs.IConstruct
	// A reference to a QueuePolicy resource.
	// Experimental.
	QueuePolicyRef() *QueuePolicyReference
}

// The jsii proxy for IQueuePolicyRef
type jsiiProxy_IQueuePolicyRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IQueuePolicyRef) QueuePolicyRef() *QueuePolicyReference {
	var returns *QueuePolicyReference
	_jsii_.Get(
		j,
		"queuePolicyRef",
		&returns,
	)
	return returns
}

