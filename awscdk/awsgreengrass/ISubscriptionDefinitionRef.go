package awsgreengrass

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgreengrass/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SubscriptionDefinition.
// Experimental.
type ISubscriptionDefinitionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a SubscriptionDefinition resource.
	// Experimental.
	SubscriptionDefinitionRef() *SubscriptionDefinitionReference
}

// The jsii proxy for ISubscriptionDefinitionRef
type jsiiProxy_ISubscriptionDefinitionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ISubscriptionDefinitionRef) SubscriptionDefinitionRef() *SubscriptionDefinitionReference {
	var returns *SubscriptionDefinitionReference
	_jsii_.Get(
		j,
		"subscriptionDefinitionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISubscriptionDefinitionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISubscriptionDefinitionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

