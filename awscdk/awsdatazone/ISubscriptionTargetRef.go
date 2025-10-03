package awsdatazone

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatazone/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SubscriptionTarget.
// Experimental.
type ISubscriptionTargetRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISubscriptionTargetRef
type jsiiProxy_ISubscriptionTargetRef struct {
	internal.Type__constructsIConstruct
}

