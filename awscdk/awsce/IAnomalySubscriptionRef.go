package awsce

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsce/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AnomalySubscription.
// Experimental.
type IAnomalySubscriptionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAnomalySubscriptionRef
type jsiiProxy_IAnomalySubscriptionRef struct {
	internal.Type__constructsIConstruct
}

