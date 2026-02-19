package interfacesawslogs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SubscriptionFilter.
// Experimental.
type ISubscriptionFilterRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a SubscriptionFilter resource.
	// Experimental.
	SubscriptionFilterRef() *SubscriptionFilterReference
}

// The jsii proxy for ISubscriptionFilterRef
type jsiiProxy_ISubscriptionFilterRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ISubscriptionFilterRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
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

func (j *jsiiProxy_ISubscriptionFilterRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISubscriptionFilterRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

