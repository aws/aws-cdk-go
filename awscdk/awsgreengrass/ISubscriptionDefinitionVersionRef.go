package awsgreengrass

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgreengrass/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SubscriptionDefinitionVersion.
// Experimental.
type ISubscriptionDefinitionVersionRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISubscriptionDefinitionVersionRef
type jsiiProxy_ISubscriptionDefinitionVersionRef struct {
	internal.Type__constructsIConstruct
}

