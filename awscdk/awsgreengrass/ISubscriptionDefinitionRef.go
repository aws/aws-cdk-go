package awsgreengrass

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgreengrass/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SubscriptionDefinition.
// Experimental.
type ISubscriptionDefinitionRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISubscriptionDefinitionRef
type jsiiProxy_ISubscriptionDefinitionRef struct {
	internal.Type__constructsIConstruct
}

