package awsce

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsce/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AnomalySubscription.
// Experimental.
type IAnomalySubscriptionRef interface {
	constructs.IConstruct
	// A reference to a AnomalySubscription resource.
	// Experimental.
	AnomalySubscriptionRef() *AnomalySubscriptionReference
}

// The jsii proxy for IAnomalySubscriptionRef
type jsiiProxy_IAnomalySubscriptionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAnomalySubscriptionRef) AnomalySubscriptionRef() *AnomalySubscriptionReference {
	var returns *AnomalySubscriptionReference
	_jsii_.Get(
		j,
		"anomalySubscriptionRef",
		&returns,
	)
	return returns
}

