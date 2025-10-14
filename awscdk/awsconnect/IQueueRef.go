package awsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Queue.
// Experimental.
type IQueueRef interface {
	constructs.IConstruct
	// A reference to a Queue resource.
	// Experimental.
	QueueRef() *QueueReference
}

// The jsii proxy for IQueueRef
type jsiiProxy_IQueueRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IQueueRef) QueueRef() *QueueReference {
	var returns *QueueReference
	_jsii_.Get(
		j,
		"queueRef",
		&returns,
	)
	return returns
}

