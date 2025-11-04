package awsgreengrass

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgreengrass/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SubscriptionDefinitionVersion.
// Experimental.
type ISubscriptionDefinitionVersionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a SubscriptionDefinitionVersion resource.
	// Experimental.
	SubscriptionDefinitionVersionRef() *SubscriptionDefinitionVersionReference
}

// The jsii proxy for ISubscriptionDefinitionVersionRef
type jsiiProxy_ISubscriptionDefinitionVersionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ISubscriptionDefinitionVersionRef) SubscriptionDefinitionVersionRef() *SubscriptionDefinitionVersionReference {
	var returns *SubscriptionDefinitionVersionReference
	_jsii_.Get(
		j,
		"subscriptionDefinitionVersionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISubscriptionDefinitionVersionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISubscriptionDefinitionVersionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

