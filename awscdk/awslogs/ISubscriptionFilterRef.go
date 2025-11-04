package awslogs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SubscriptionFilter.
// Experimental.
type ISubscriptionFilterRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a SubscriptionFilter resource.
	// Experimental.
	SubscriptionFilterRef() *SubscriptionFilterReference
}

// The jsii proxy for ISubscriptionFilterRef
type jsiiProxy_ISubscriptionFilterRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_ISubscriptionFilterRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
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

