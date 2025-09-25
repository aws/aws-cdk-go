package awsbatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a JobQueue.
// Experimental.
type IJobQueueRef interface {
	constructs.IConstruct
	// A reference to a JobQueue resource.
	// Experimental.
	JobQueueRef() *JobQueueReference
}

// The jsii proxy for IJobQueueRef
type jsiiProxy_IJobQueueRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IJobQueueRef) JobQueueRef() *JobQueueReference {
	var returns *JobQueueReference
	_jsii_.Get(
		j,
		"jobQueueRef",
		&returns,
	)
	return returns
}

