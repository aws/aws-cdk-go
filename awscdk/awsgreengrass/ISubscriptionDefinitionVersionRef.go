package awsgreengrass

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsgreengrass/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SubscriptionDefinitionVersion.
// Experimental.
type ISubscriptionDefinitionVersionRef interface {
	constructs.IConstruct
	// A reference to a SubscriptionDefinitionVersion resource.
	// Experimental.
	SubscriptionDefinitionVersionRef() *SubscriptionDefinitionVersionReference
}

// The jsii proxy for ISubscriptionDefinitionVersionRef
type jsiiProxy_ISubscriptionDefinitionVersionRef struct {
	internal.Type__constructsIConstruct
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

