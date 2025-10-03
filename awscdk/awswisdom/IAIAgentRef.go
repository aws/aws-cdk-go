package awswisdom

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awswisdom/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AIAgent.
// Experimental.
type IAIAgentRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAIAgentRef
type jsiiProxy_IAIAgentRef struct {
	internal.Type__constructsIConstruct
}

