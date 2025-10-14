package awsdatazone

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatazone/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SubscriptionTarget.
// Experimental.
type ISubscriptionTargetRef interface {
	constructs.IConstruct
	// A reference to a SubscriptionTarget resource.
	// Experimental.
	SubscriptionTargetRef() *SubscriptionTargetReference
}

// The jsii proxy for ISubscriptionTargetRef
type jsiiProxy_ISubscriptionTargetRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISubscriptionTargetRef) SubscriptionTargetRef() *SubscriptionTargetReference {
	var returns *SubscriptionTargetReference
	_jsii_.Get(
		j,
		"subscriptionTargetRef",
		&returns,
	)
	return returns
}

