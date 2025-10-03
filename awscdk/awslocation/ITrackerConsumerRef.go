package awslocation

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslocation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TrackerConsumer.
// Experimental.
type ITrackerConsumerRef interface {
	constructs.IConstruct
}

// The jsii proxy for ITrackerConsumerRef
type jsiiProxy_ITrackerConsumerRef struct {
	internal.Type__constructsIConstruct
}

