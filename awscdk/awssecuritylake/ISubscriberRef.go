package awssecuritylake

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssecuritylake/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Subscriber.
// Experimental.
type ISubscriberRef interface {
	constructs.IConstruct
	// A reference to a Subscriber resource.
	// Experimental.
	SubscriberRef() *SubscriberReference
}

// The jsii proxy for ISubscriberRef
type jsiiProxy_ISubscriberRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISubscriberRef) SubscriberRef() *SubscriberReference {
	var returns *SubscriberReference
	_jsii_.Get(
		j,
		"subscriberRef",
		&returns,
	)
	return returns
}

