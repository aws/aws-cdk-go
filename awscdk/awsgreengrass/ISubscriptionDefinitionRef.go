package awsgreengrass

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsgreengrass/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SubscriptionDefinition.
// Experimental.
type ISubscriptionDefinitionRef interface {
	constructs.IConstruct
	// A reference to a SubscriptionDefinition resource.
	// Experimental.
	SubscriptionDefinitionRef() *SubscriptionDefinitionReference
}

// The jsii proxy for ISubscriptionDefinitionRef
type jsiiProxy_ISubscriptionDefinitionRef struct {
	internal.Type__constructsIConstruct
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

