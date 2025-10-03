package awssns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Subscription.
// Experimental.
type ISubscriptionRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISubscriptionRef
type jsiiProxy_ISubscriptionRef struct {
	internal.Type__constructsIConstruct
}

