package awswisdom

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awswisdom/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AIPrompt.
// Experimental.
type IAIPromptRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAIPromptRef
type jsiiProxy_IAIPromptRef struct {
	internal.Type__constructsIConstruct
}

