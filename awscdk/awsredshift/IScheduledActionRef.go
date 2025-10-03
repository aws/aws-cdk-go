package awsredshift

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsredshift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ScheduledAction.
// Experimental.
type IScheduledActionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IScheduledActionRef
type jsiiProxy_IScheduledActionRef struct {
	internal.Type__constructsIConstruct
}

