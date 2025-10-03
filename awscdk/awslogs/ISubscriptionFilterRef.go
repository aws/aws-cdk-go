package awslogs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SubscriptionFilter.
// Experimental.
type ISubscriptionFilterRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISubscriptionFilterRef
type jsiiProxy_ISubscriptionFilterRef struct {
	internal.Type__constructsIConstruct
}

