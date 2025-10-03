package awssecurityhub

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecurityhub/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ProductSubscription.
// Experimental.
type IProductSubscriptionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IProductSubscriptionRef
type jsiiProxy_IProductSubscriptionRef struct {
	internal.Type__constructsIConstruct
}

