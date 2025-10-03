package awssecuritylake

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecuritylake/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Subscriber.
// Experimental.
type ISubscriberRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISubscriberRef
type jsiiProxy_ISubscriberRef struct {
	internal.Type__constructsIConstruct
}

