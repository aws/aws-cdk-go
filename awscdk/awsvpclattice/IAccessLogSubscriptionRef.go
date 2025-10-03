package awsvpclattice

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsvpclattice/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AccessLogSubscription.
// Experimental.
type IAccessLogSubscriptionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAccessLogSubscriptionRef
type jsiiProxy_IAccessLogSubscriptionRef struct {
	internal.Type__constructsIConstruct
}

