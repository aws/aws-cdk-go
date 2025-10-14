package awssns

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssns/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Subscription.
// Experimental.
type ISubscriptionRef interface {
	constructs.IConstruct
	// A reference to a Subscription resource.
	// Experimental.
	SubscriptionRef() *SubscriptionReference
}

// The jsii proxy for ISubscriptionRef
type jsiiProxy_ISubscriptionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISubscriptionRef) SubscriptionRef() *SubscriptionReference {
	var returns *SubscriptionReference
	_jsii_.Get(
		j,
		"subscriptionRef",
		&returns,
	)
	return returns
}

