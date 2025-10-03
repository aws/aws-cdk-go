package awsredshift

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsredshift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EventSubscription.
// Experimental.
type IEventSubscriptionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IEventSubscriptionRef
type jsiiProxy_IEventSubscriptionRef struct {
	internal.Type__constructsIConstruct
}

