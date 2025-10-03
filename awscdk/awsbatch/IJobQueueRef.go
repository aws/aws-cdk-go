package awsbatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a JobQueue.
// Experimental.
type IJobQueueRef interface {
	constructs.IConstruct
}

// The jsii proxy for IJobQueueRef
type jsiiProxy_IJobQueueRef struct {
	internal.Type__constructsIConstruct
}

