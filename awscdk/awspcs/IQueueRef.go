package awspcs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awspcs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Queue.
// Experimental.
type IQueueRef interface {
	constructs.IConstruct
}

// The jsii proxy for IQueueRef
type jsiiProxy_IQueueRef struct {
	internal.Type__constructsIConstruct
}

