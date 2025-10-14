package awsvpclattice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsvpclattice/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AccessLogSubscription.
// Experimental.
type IAccessLogSubscriptionRef interface {
	constructs.IConstruct
	// A reference to a AccessLogSubscription resource.
	// Experimental.
	AccessLogSubscriptionRef() *AccessLogSubscriptionReference
}

// The jsii proxy for IAccessLogSubscriptionRef
type jsiiProxy_IAccessLogSubscriptionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAccessLogSubscriptionRef) AccessLogSubscriptionRef() *AccessLogSubscriptionReference {
	var returns *AccessLogSubscriptionReference
	_jsii_.Get(
		j,
		"accessLogSubscriptionRef",
		&returns,
	)
	return returns
}

