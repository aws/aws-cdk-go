package awslogs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SubscriptionFilter.
// Experimental.
type ISubscriptionFilterRef interface {
	constructs.IConstruct
	// A reference to a SubscriptionFilter resource.
	// Experimental.
	SubscriptionFilterRef() *SubscriptionFilterReference
}

// The jsii proxy for ISubscriptionFilterRef
type jsiiProxy_ISubscriptionFilterRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISubscriptionFilterRef) SubscriptionFilterRef() *SubscriptionFilterReference {
	var returns *SubscriptionFilterReference
	_jsii_.Get(
		j,
		"subscriptionFilterRef",
		&returns,
	)
	return returns
}

