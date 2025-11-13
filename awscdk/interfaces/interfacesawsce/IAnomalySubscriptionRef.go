package interfacesawsce

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsce/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AnomalySubscription.
// Experimental.
type IAnomalySubscriptionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a AnomalySubscription resource.
	// Experimental.
	AnomalySubscriptionRef() *AnomalySubscriptionReference
}

// The jsii proxy for IAnomalySubscriptionRef
type jsiiProxy_IAnomalySubscriptionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
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

func (j *jsiiProxy_IAnomalySubscriptionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAnomalySubscriptionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

