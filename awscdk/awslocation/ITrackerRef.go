package awslocation

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslocation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Tracker.
// Experimental.
type ITrackerRef interface {
	constructs.IConstruct
}

// The jsii proxy for ITrackerRef
type jsiiProxy_ITrackerRef struct {
	internal.Type__constructsIConstruct
}

